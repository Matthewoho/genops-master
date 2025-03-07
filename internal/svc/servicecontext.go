package svc

import (
	"fmt"
	"genops-master/internal/config"
	"genops-master/internal/middleware"
	"genops-master/internal/models"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx" // 导入 go-zero 的 sqlx 包，用于数据库连接操作
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config          config.Config
	AuthMiddleware  rest.Middleware
	RbacMiddleware  rest.Middleware
	RedisClient     *redis.Client
	MySQLClient     sqlx.SqlConn
	MySQLUsersModel models.USERSModel
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

	// 初始化 MySQLClient
	mysqlClient := sqlx.NewMysql(c.Mysql.Addr)

	// 测试 MySQL 连接
	if err := testMySQLConnection(mysqlClient); err != nil {
		fmt.Printf("无法连接到 MySQL: %v\n", err)
	}

	return &ServiceContext{
		Config:          c,
		AuthMiddleware:  middleware.NewAuthMiddleware(rdb).Handle,
		RbacMiddleware:  middleware.NewRbacMiddleware().Handle,
		RedisClient:     rdb,
		MySQLClient:     mysqlClient,
		MySQLUsersModel: models.NewUSERSModel(mysqlClient),
	}
}

// 测试 MySQL 连接
func testMySQLConnection(conn sqlx.SqlConn) error {
	// 尝试执行一个简单的查询来测试数据库连接
	_, err := conn.Exec("SELECT 1")
	if err != nil {
		// 如果发生错误，返回错误对象
		return fmt.Errorf("MySQL connection test failed: %w", err)
	}
	// 如果没有错误，返回nil表示连接正常
	return nil
}
