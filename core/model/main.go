package model

const MainFile = `package main

import (
	"app/public"
	"app/server"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// @title ginhelper
// @version 1.0
// @description ginhelper

// @contact.name yanmengfei
// @contact.email 3940422@qq.com

// @host 127.0.0.1:9404
func main() {
	if err := server.LoadConfig("config.yaml"); err != nil {
		panic(err)
	}
	var engine = server.InitEngine()
	httpServer := &http.Server{
		Addr:           ":" + public.Config.Server.RunPort,              // 监听地址
		MaxHeaderBytes: 1 << 20,                                         // 1048576
		Handler:        engine,                                          // 服务引擎
		ReadTimeout:    public.Config.Server.ReadTimeout * time.Second,  // 请求超市
		WriteTimeout:   public.Config.Server.WriteTimeout * time.Second, // 响应超时
	}
	if err := httpServer.ListenAndServe(); err != nil {
		public.Logger.Error("服务启动失败", zap.Any("err", err))
	}
}
`
