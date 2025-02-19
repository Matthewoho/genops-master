package svc

import (
	"fmt"
	"genops-master/internal/config"
	"genops-master/internal/middleware"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	RbacMiddleware rest.Middleware
	RedisClient    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr, // 假设你有 Redis 的配置在 config.Config 中
		Password: c.Redis.Password,
	})

	// 测试连接
	_, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		fmt.Printf("Could not connect to Redis: %v", err)
	}

	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
		RbacMiddleware: middleware.NewRbacMiddleware().Handle,
		RedisClient:    rdb,
	}
}
