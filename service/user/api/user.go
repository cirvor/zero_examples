package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"

	"zero_examples/common/response/errorx"
	"zero_examples/service/user/api/internal/config"
	"zero_examples/service/user/api/internal/handler"
	"zero_examples/service/user/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "service/user/api/etc/config.conf", "the config file")

func main() {
	flag.Parse()

	baseConfigFile := "common/etc/base-config.yaml"

	var c config.Config

	// load all configs
	conf.MustLoad(baseConfigFile, &c)

	// load private configs
	if *configFile != "" {
		props, err := conf.LoadProperties(*configFile)
		if err != nil {
			panic(err)
		}
		c.Name = props.GetString("Name")
		c.Port = props.GetInt("Port")
	}

	// load logx configs
	logx.MustSetup(c.LogConf)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			logx.Error(e.Error())
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, e.Error()
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
