package model

const SchemaCommonFile = `package schemas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// registerValidatorRule 注册参数验证错误消息, Key = e.StructNamespace(), value.key = e.Field()+e.Tag()
var registerValidatorRule = map[string]map[string]string{}

// serializeValidatorError 参数tag验证失败转换
func serializeValidatorError(e validator.FieldError) (message string) {
	switch e.Field() {
	default:
		message = registerValidatorRule[strings.Split(e.StructNamespace(), ".")[0]][e.Field()+e.Tag()]
	case "Page", "Size":
		switch e.Tag() {
		case "min":
			message = e.Field() + "最小值为" + e.Param()
		case "max":
			message = e.Field() + "最大值为" + e.Param()
		}
	}
	return message
}

// serializeTypeError 参数类型错误转换
func serializeTypeError(e *json.UnmarshalTypeError) string {
	return fmt.Sprintf("参数%s类型错误, 预期%s, 接收到%s", e.Field, e.Type, e.Value)
}

// BindSchema 绑定单类型的请求参数
func BindSchema(c *gin.Context, obj interface{}, bind binding.Binding) (err error) {
	if err = c.ShouldBindWith(obj, bind); err != nil {
		var response ErrorResponse
		switch err := err.(type) {
		case *json.UnmarshalTypeError:
			response.Message = serializeTypeError(err)
		case validator.ValidationErrors:
			response.Message = serializeValidatorError(err[0])
		default:
			response.Message = "无效的请求参数"
		}
		zap.L().Error(response.Message, zap.Error(err))
		c.JSON(http.StatusBadRequest, response)
	}
	return err
}
`

const SchemaResponseFile = `package schemas

type SuccessResponse struct {
	Status bool        ` + "`" + `json:"status"` + "`" + ` // 请求状态
	Data   interface{} ` + "`" + `json:"data"` + "`" + `   // 响应数据体
}

type ErrorResponse struct {
	Status  bool   ` + "`" + `json:"status"` + "`" + `  // 状态
	Message string ` + "`" + `json:"message"` + "`" + ` // 错误消息
}

type QueryListResponse struct {
	TotalCount  int64       ` + "`" + `json:"total_count"` + "`" + `  // 请求资源总计数
	FilterCount int64       ` + "`" + `json:"filter_count"` + "`" + ` // 请求资源过滤计数
	Records     interface{} ` + "`" + `json:"records"` + "`" + `      // 资源记录
}
`
