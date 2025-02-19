package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"genops-master/internal/types"
	"time"

	"github.com/dchest/captcha"
	"github.com/go-redis/redis/v8"
)

const captchaExpireDuration = 5 * time.Minute // 验证码过期时间常量

type Captcha struct {
	types.Captcha
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
