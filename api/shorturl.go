package main

import (
	"flag"
	"fmt"

	"shorturl/api/internal/config"
	"shorturl/api/internal/handler"
	"shorturl/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/shorturl-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	//调用 handler.RegisterHandlers 函数，将 server 和 ctx 传入，注册请求处理函数到服务器上。
	//这些处理函数会在服务器接收到请求时被调用，根据请求的 URL、方法等信息进行相应的处理。
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
