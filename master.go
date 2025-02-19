package main

import (
	"flag"
	"fmt"
	"genops-master/internal/biz"
	"genops-master/internal/config"
	"genops-master/internal/handler"
	"genops-master/internal/svc"
	"net/http"

	"errors"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/master.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 统一的错误处理
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		var e *biz.Error
		switch {
		case errors.As(err, &e):
			return http.StatusBadRequest, biz.Fail(e)

		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
