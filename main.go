package main

import (
	"flag"
	"fmt"
	"net/http"
	"spotify-server/core"
	"spotify-server/router"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	// 从命令行配置中获取配置文件路径
	// 默认从根目录config/config.yaml中读取
	filename := flag.String("config", "config.yaml", "config filename")
	flag.Parse()

	// 初始化viper
	core.InitViper(filename)
}

func main() {
	gin.SetMode(getMode())
	fmt.Printf("mode: %v, port: %d\n", getMode(), getPort())

	engine := router.New()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", getPort()),
		Handler:        engine,
		ReadTimeout:    1 * time.Minute,
		WriteTimeout:   1 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}

	core.RunServer(server)
}

// 根据配置获取mode，默认为release
func getMode() string {
	debug := core.Viper.GetBool("server.debug")
	mode := gin.DebugMode

	if !debug {
		mode = gin.ReleaseMode
	}

	return mode
}

// getPort 根据配置文件设置端口，默认8080
func getPort() int {
	port := core.Viper.GetInt("server.port")

	if port == 0 {
		port = 8080
	}

	return port
}
