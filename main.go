package main

import (
	"flag"
	"fmt"
	"net/http"
	"spotify-server/core"
	"spotify-server/router"
	"time"
)

func init() {
	// 从命令行配置中获取配置文件路径
	// 默认从根目录config.yaml中读取
	filename := flag.String("config", "config.yaml", "config filename")
	flag.Parse()

	// 初始化viper
	core.InitViper(filename)
}

func main() {
	addr := fmt.Sprintf(":%d", core.Viper.GetInt("server.port"))

	engine := router.New()

	server := &http.Server{
		Addr:           addr,
		Handler:        engine,
		ReadTimeout:    1 * time.Minute,
		WriteTimeout:   1 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}

	core.RunServer(server)
}
