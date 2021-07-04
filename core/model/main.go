package model

const GinMainFile = `package main

import (
	"app/public"
	"app/setup"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// @title carapi Example
// @version 1.0
// @description This is a sample swagger for carapi
// @termsOfService http://swagger.io/terms/
// @contact.name yanmengfei
// @contact.email 3940422@qq.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9000
// @BasePath /
func main() {
	if err := setup.Config("config.toml"); err != nil {
		panic(err)
	}
	var engine = setup.InitEngine()
	httpServer := &http.Server{
		Addr:           ":" + public.Cfg.Server.Port,                 // 监听地址
		MaxHeaderBytes: 1 << 20,                                      // 1048576
		Handler:        engine,                                       // 服务引擎
		ReadTimeout:    public.Cfg.Server.ReadTimeout * time.Second,  // 请求超市
		WriteTimeout:   public.Cfg.Server.WriteTimeout * time.Second, // 响应超时
	}
	if err := httpServer.ListenAndServe(); err != nil {
		zap.L().Error("服务启动失败", zap.Any("err", err))
	}
}
`

const FiberMainFile = `package main

import (
	_ "app/docs"
	"app/public"
	"app/setup"
	"fmt"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"log"
)

// @title carapi Example
// @version 1.0
// @description This is a sample swagger for carapi
// @termsOfService http://swagger.io/terms/
// @contact.name yanmengfei
// @contact.email 3940422@qq.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9000
// @BasePath /
func main() {
	if err := setup.Config("config.toml"); err != nil {
		log.Fatal(err)
	}
	engine := setup.Engine()

	engine.Get("/swagger/*", swagger.Handler)
	log.Fatalln(engine.Listen(fmt.Sprintf(":%d", public.Cfg.Server.Port)))
}
`
