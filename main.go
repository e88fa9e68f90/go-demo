package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kiwi633/go-demo/package1"
	"github.com/kiwi633/go-demo/remote"
	zap_log "github.com/kiwi633/go-demo/zap-log"
)

func main() {
	// 初始化logger
	zap_log.InitLogger()
	defer zap_log.Log.Sync()

	// 创建gin引擎
	r := gin.New()

	// 使用中间件
	r.Use(zap_log.TraceMiddleware())
	r.Use(gin.Recovery())

	// 注册路由
	r.GET("/ping", package1.UserHandler)
	r.GET("/user", package1.PersonList)
	r.GET("/getpolicylist", remote.GetPolicyList)
	// 启动服务
	r.Run(":8888")
}
