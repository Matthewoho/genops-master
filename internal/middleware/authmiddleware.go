package middleware

import (
	"fmt"
	"genops-master/internal/config"
	"genops-master/internal/tools"
	"net/http"
	"strings"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/conf"
)

type AuthMiddleware struct {
	Config      *config.Config
	RedisClient *redis.Client
}

var (
	configOnce sync.Once
	globalCfg  config.Config
)

// NewAuthMiddleware 初始化中间件，并全局加载配置（只加载一次）
func NewAuthMiddleware(redisClient *redis.Client) *AuthMiddleware {
	configOnce.Do(func() {
		conf.MustLoad("etc/master.yaml", &globalCfg)
	})

	return &AuthMiddleware{
		Config:      &globalCfg,
		RedisClient: redisClient,
	}
}

// Handle 认证中间件
func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 获取 Authorization 头
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// 确保 Token 以 "Bearer " 开头
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == authHeader { // 没有 "Bearer " 前缀
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		redisKey := "access." + tokenStr

		// 检查 Token 是否在 Redis 黑名单中
		check, err := m.RedisClient.Get(m.RedisClient.Context(), redisKey).Result()
		if err != nil {
			if err == redis.Nil {
				http.Error(w, "Token not found", http.StatusUnauthorized)
				// 尝试刷新 Token
				return
			}
			http.Error(w, "Error checking token in Redis", http.StatusInternalServerError)
			return
		}
		if check == "revoked" {
			http.Error(w, "Token revoked or expired", http.StatusUnauthorized)
			return
		}

		// 解析 JWT Token
		claims, err := tools.ParseToken(redisKey, m.Config.Auth.AccessSecret)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid access token: %s", err), http.StatusUnauthorized)
			return
		}

		// 确保 Token 里有用户名
		username := claims.Username
		if username == "" {
			http.Error(w, "Invalid token: missing username", http.StatusUnauthorized)
			return
		}

		// 将用户名存入请求头（用于后续处理）
		r.Header.Set("X-Username", username)

		// 继续执行下一个处理器
		next(w, r)
	}
}

func TryRefreshToken(AccessSecuretKeyString string, refreshTokenSecuretKeyString string, refreshToken string) error {
	tools.RefreshAccessToken(AccessSecuretKeyString, refreshTokenSecuretKeyString, refreshToken)
	return nil
}
