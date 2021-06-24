package model

const ApiTestFile = `// 服务测试接口
package api

import (
	"app/schemas"

	"github.com/gin-gonic/gin"
)

// @Summary 服务连通性测试
// @Description 服务连通性测试接口
// @Tags 测试
// @Success 200 {object} schemas.SuccessResponse
// @Failure 400 {object} schemas.ErrorResponse
// @Router /test/ping [get]
func PingApi(c *gin.Context) {
	c.JSON(200, schemas.SuccessResponse{Status: true, Data: "pong"})
}
`
