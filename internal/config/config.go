package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret  string
		RefreshSecret string
	}
	Mysql struct {
		Addr string
	}
	Redis struct {
		Addr     string
		Password string
	}
	Consul struct {
		Addr string
	}
}
