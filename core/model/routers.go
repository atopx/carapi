package model

const RouterTestFile = `package routers

import (
	"app/api"

	"github.com/gin-gonic/gin"
)

func InitPingRouter(engine *gin.Engine) gin.IRoutes {
	var group = engine.Group("/test")
	{
		group.GET("/ping", api.PingApi)
	}
	return group

}
`
