// Package app kratos.go
package app

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	conf "github.com/ickeep/yimoko-go/config"
)

// NewApp http.Server 支持多个 可分别支持后台管理 API 和 BFF API
func NewApp(logger log.Logger, conf *conf.Server, gs *grpc.Server, hss ...*http.Server) *kratos.App {
	servers := []transport.Server{}
	for _, s := range hss {
		if s != nil {
			servers = append(servers, s)
		}
	}
	if gs != nil {
		servers = append(servers, gs)
	}
	if len(servers) == 0 {
		panic("请检查配置文件，至少保证启动 GRPC 或 HTTP 中的一个服务")
	}
	return kratos.New(
		kratos.ID(conf.Id),
		kratos.Name(conf.Name),
		kratos.Version(conf.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(servers...),
	)
}

// UpdateDefaultConfig 根据启动参数更新配置文件
func UpdateDefaultConfig(conf *conf.Config, id string, name string, version string) *conf.Config {
	if conf.Server != nil {
		if conf.Server.Id == "" {
			conf.Server.Id = id
		}
		if conf.Server.Name == "" {
			conf.Server.Name = name
		}
		if conf.Server.Version == "" {
			conf.Server.Version = version
		}
	}
	if conf.Trace != nil {
		if conf.Trace.Service == "" {
			conf.Trace.Service = name
		}
	}

	return conf
}

// LoadFileConf 加载文件配置
func LoadFileConf(flagconf string) (config.Config, *conf.Config) {
	c := config.New(config.WithSource(file.NewSource(flagconf)))
	if err := c.Load(); err != nil {
		panic(err)
	}
	var bc conf.Config
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	return c, &bc
}
