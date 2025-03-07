Genops Master
Genops Master 是一个具备认证、权限管理、用户信息处理等功能的 Go 语言项目，集成了 MySQL、Redis 等服务，支持用户认证、角色管理、权限验证等操作。
功能特性
用户认证：基于 JWT 的认证机制，支持用户登录和 Token 验证。
权限管理：使用 RBAC（基于角色的访问控制）进行权限管理。
数据存储：集成 MySQL 和 Redis，用于持久化存储用户数据和缓存 Token 信息。
环境要求
Go 1.16 或更高版本
MySQL 5.7 或更高版本
Redis 6.0 或更高版本
Consul 1.8 或更高版本
安装步骤
1. 克隆代码库
bash
git clone https://github.com/your-repo/genops-master.git
cd genops-master
2. 配置环境
在 etc/master.yaml 文件中配置相关服务信息，示例配置如下：
yaml
Name: master                      # 集群名称
Host: 0.0.0.0                     # 监听地址
Port: 8888                        # 监听端口
Auth:
  AccessSecret: "damndamn"            # 访问密钥，可为空
  RefreshSecret: "damndamnit"            # 刷新密钥，可为空
Mysql:
  Addr: root:000000@tcp(192.168.110.46:3306)/test?charset=utf8mb4&parseTime=True&loc=Local
 
Redis:
  Addr: 192.168.110.46:6379       # Redis地址
  Password: ""                    # Redis密码
Consul:
  Addr: 192.168.110.46:8500        # Consul地址
  Token: ""                        # Consul Token
3. 下载依赖
bash
go mod tidy
4. 构建项目
bash
go build -o /var/api/genops-master/master-api genops-master
5. 复制配置文件
bash
cp -r ./etc /var/api/genops-master/etc
6. 部署项目
bash
scp -r /var/api/genops-master root@192.168.110.46:/var/api/genops-master
使用说明
认证中间件
认证中间件用于验证用户的 Token，确保请求的合法性。在 internal/middleware/authmiddleware.go 中实现，使用示例如下：
go
// 创建认证中间件实例
authMiddleware := middleware.NewAuthMiddleware(&config, redisClient)

// 使用认证中间件
router.Use(authMiddleware.Handle)
RBAC 中间件
RBAC 中间件用于实现基于角色的访问控制，在 internal/middleware/rbacmiddleware.go 中实现。目前该中间件还处于待实现状态，需要根据具体需求进行完善。
用户服务
用户服务提供了获取用户信息、验证密码等功能，在 internal/svc/userservice.go 中实现，使用示例如下：
go
// 创建用户服务实例
userService := svc.NewUserService(userModel)

// 获取用户信息
user, err := userService.GetUserInfoByUsername(ctx, "username")
if err != nil {
    // 处理错误
}

// 验证密码
valid, err := userService.VerifyPassword(ctx, "username", "password")
if err != nil {
    // 处理错误
}
贡献指南
如果你想为该项目做出贡献，请遵循以下步骤：
Fork 该项目
创建你的特性分支 (git checkout -b feature/your-feature)
提交你的更改 (git commit -am 'Add some feature')
将你的更改推送到远程分支 (git push origin feature/your-feature)
打开一个 Pull Request
许可证
本项目采用 MIT 许可证。
联系信息
如果你有任何问题或建议，请随时联系我们：
邮箱：matthewohmygosh@gmail.com
项目地址：https://github.com/your-repo/genops-master