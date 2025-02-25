package main // 声明主包

import ( // 导入所需的包
	"flag"                           // 用于解析命令行参数
	"fmt"                            // 提供格式化 I/O 功能
	"genops-master/internal/biz"     // 引入业务逻辑相关包
	"genops-master/internal/config"  // 引入配置相关包
	"genops-master/internal/handler" // 引入请求处理器相关包
	"genops-master/internal/svc"     // 引入服务上下文相关包
	"net/http"                       // 提供 HTTP 客户端和服务器功能

	"errors" // 标准错误处理包

	"github.com/zeromicro/go-zero/core/conf"  // go-zero 框架的配置加载包
	"github.com/zeromicro/go-zero/rest"       // go-zero 的 REST 服务器包
	"github.com/zeromicro/go-zero/rest/httpx" // go-zero 的 HTTP 扩展包，用于处理错误等
)

var configFile = flag.String("f", "etc/master.yaml", "the config file") // 定义命令行标志 -f，指定配置文件路径，默认值为 etc/master.yaml

func main() { // main 函数，程序入口点
	flag.Parse() // 解析命令行参数

	var c config.Config            // 声明变量 c，类型为 config.Config，用于存储配置
	conf.MustLoad(*configFile, &c) // 从指定的配置文件加载配置到变量 c，加载失败则终止程序

	server := rest.MustNewServer(c.RestConf) // 根据配置中的 RestConf 创建新的 REST 服务器，创建失败时程序将终止
	defer server.Stop()                      // 程序退出前，确保停止服务器以释放资源

	ctx := svc.NewServiceContext(c)       // 创建服务上下文 ctx，将配置传递进去
	handler.RegisterHandlers(server, ctx) // 注册 HTTP 请求处理器，将服务器和服务上下文关联

	// 统一的错误处理
	httpx.SetErrorHandler(func(err error) (int, interface{}) { // 设置全局错误处理函数
		var e *biz.Error // 声明一个 *biz.Error 类型的变量，用于错误类型断言
		switch {
		case errors.As(err, &e): // 如果错误 err 能够转换为 *biz.Error 类型，则进入该分支
			return http.StatusBadRequest, biz.Fail(e) // 返回 HTTP 400 状态码和 biz.Fail(e) 生成的错误响应
		default:
			return http.StatusInternalServerError, nil // 否则返回 HTTP 500 状态码和 nil 响应
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port) // 输出服务器启动信息，包括主机地址和端口号
	server.Start()                                              // 启动 REST 服务器，开始监听并处理 HTTP 请求
}
