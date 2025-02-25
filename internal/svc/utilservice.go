package svc

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"genops-master/internal/types"
	"math/rand"
	"time"

	"github.com/dchest/captcha"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Captcha struct {
	types.Captcha
}

type Tokens struct {
	types.TokenRspData
}

// Claims 定义 JWT 中携带的数据，你可以根据需要增加字段
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

const captchaExpireDuration = 5 * time.Minute         // 验证码过期时间常量
const accessTokenExpireDuration = 24 * time.Hour      // accessToken过期时间常量
const refreshTokenExpireDuration = 7 * 24 * time.Hour // refreshToken过期时间常量

// 生成一个随机盐
func GenerateSalt() (string, error) {
	saltBytes := make([]byte, 16) // 16字节的盐
	_, err := rand.Read(saltBytes)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(saltBytes), nil
}

// 生成哈希密码（使用盐）
func HashPasswordWithSalt(password, salt string) (string, error) {
	saltedPassword := password + salt // 拼接密码和盐
	hash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// GenerateCaptcha 生成图形验证码
func GenerateCaptcha(redisClient *redis.Client) (*Captcha, error) {
	// 生成验证码ID
	captchaID := captcha.New()

	// 生成6位数字验证码
	digits := captcha.RandomDigits(6)

	// 设置验证码过期时间
	expireAt := time.Now().Add(captchaExpireDuration)

	// 生成验证码图片
	img := captcha.NewImage(captchaID, digits, 240, 80)
	var buf bytes.Buffer
	if _, err := img.WriteTo(&buf); err != nil {
		return nil, fmt.Errorf("could not generate captcha image: %w", err)
	}

	// 将图片转换为 Base64 编码
	captchaImage := base64.StdEncoding.EncodeToString(buf.Bytes())

	// 将数字验证码拼接成字符串
	captchaValue := fmt.Sprintf("%d%d%d%d%d%d", digits[0], digits[1], digits[2], digits[3], digits[4], digits[5])

	// 存储验证码到 Redis
	if err := StoreCaptchaInRedis(redisClient, captchaID, captchaValue, expireAt); err != nil {
		return nil, fmt.Errorf("could not store captcha in redis: %w", err)
	}

	// 返回生成的验证码
	return &Captcha{
		Captcha: types.Captcha{
			CaptchaID:     captchaID,
			CaptchaValue:  captchaValue,
			CaptchaImage:  captchaImage,
			CaptchaExpire: expireAt.Unix(),
		},
	}, nil
}

// Verify 验证图形验证码
func (c *Captcha) Verify(redisClient *redis.Client, captchaID, captchaValue string) (bool, error) {
	// 从 Redis 中获取存储的验证码
	storedValue, err := redisClient.Get(redisClient.Context(), captchaID).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil // 验证码不存在
		}
		return false, fmt.Errorf("error getting captcha from Redis: %w", err)
	}

	// 比较用户输入的验证码值与存储的验证码值
	return storedValue == captchaValue, nil
}

// storeCaptchaInRedis 存储验证码到 Redis
func StoreCaptchaInRedis(redisClient *redis.Client, captchaID, captchaValue string, expireAt time.Time) error {
	// 计算验证码过期时长
	duration := time.Until(expireAt)
	// 将验证码ID、值和过期时间存储到 Redis
	if err := redisClient.Set(redisClient.Context(), captchaID, captchaValue, duration).Err(); err != nil {
		return fmt.Errorf("could not store captcha in redis: %w", err)
	}
	return nil
}

// 创建token
func GenerateTokens(ctx *ServiceContext, username string) (*Tokens, error) {
	accessClaims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenExpireDuration)), // 访问令牌有效期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	refreshClaims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenExpireDuration)), // 刷新令牌有效期（7天）
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// 生成 Access Token
	secretKey := []byte(ctx.Config.Auth.AccessSecret)
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	// 生成 Refresh Token
	refreshSecretKey := []byte(ctx.Config.Auth.RefreshSecret)
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(refreshSecretKey)
	if err != nil {
		return nil, err
	}

	// 计算 Access Token 过期时间（秒）
	expiresIn := int(accessTokenExpireDuration.Seconds())

	// 返回 Tokens 结构体
	tokens := &Tokens{
		TokenRspData: types.TokenRspData{
			AccessToken:  accessToken,
			TokenType:    "Bearer",
			ExpiresIn:    expiresIn,
			RefreshToken: refreshToken,
		},
	}

	return tokens, nil
}

// 解析token
func ParseToken(tokenString string, secretKey string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// 刷新token
func RefreshAccessToken(ctx *ServiceContext, refreshToken string) (*Tokens, error) {
	// 解析 Refresh Token
	claims, err := ParseToken(refreshToken, ctx.Config.Auth.RefreshSecret)
	if err != nil {
		return nil, fmt.Errorf("refresh token 无效或已过期: %w", err)
	}

	// 重新生成 Access Token
	return GenerateTokens(ctx, claims.Username)
}

// 撤销token（存储黑名单）
func RevokeToken(redisClient *redis.Client, token string, expireDuration time.Duration) error {
	err := redisClient.Set(redisClient.Context(), token, "revoked", expireDuration).Err()
	if err != nil {
		return fmt.Errorf("failed to revoke token: %w", err)
	}
	return nil
}
