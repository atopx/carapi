package model

const MainCode = `package main

import (
	"app/engine"
	"app/public"
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
	if err := engine.LoadConfig("config.yaml"); err != nil {
		panic(err)
	}
	var engine = engine.InitEngine()
	server := &http.Server{
		Addr:           ":" + public.Config.Server.RunPort,              // 监听地址
		MaxHeaderBytes: 1 << 20,                                         // 1048576
		Handler:        engine,                                          // 服务引擎
		ReadTimeout:    public.Config.Server.ReadTimeout * time.Second,  // 请求超市
		WriteTimeout:   public.Config.Server.WriteTimeout * time.Second, // 响应超时
	}
	if err := server.ListenAndServe(); err != nil {
		public.Logger.Error("服务启动失败", zap.Any("err", err))
	}
}
`
