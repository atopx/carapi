package model

const GinSchemaCommonFile = `package schema

import (
	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var validate *validator.Validate
var trans ut.Translator

func init() {
	// 初始化翻译器
	trans, _ = ut.New(zh_Hans_CN.New()).GetTranslator("zh")
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("label")
	})
	if err := zh.RegisterDefaultTranslations(validate, trans); err != nil {
		zap.L().Error("register validator failed", zap.Error(err))
	}
}

// Validator 参数验证器
func Validator(data interface{}, c *gin.Context, b binding.Binding) (err error) {
	if err = c.ShouldBindWith(data, b); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return err
	}
	if err = validate.Struct(data); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.(validator.ValidationErrors)[0].Translate(trans)})
		return err
	}
	return nil
}

// SuccessResponse 完成响应结构体
type SuccessResponse struct {
	Status bool        ` + "`" + `json:"status"` + "`" + `
	Data   interface{} ` + "`" + `json:"data"` + "`" + `
}

// ErrorResponse 错误响应结构体
type ErrorResponse struct {
	Status  bool   ` + "`" + `json:"status"` + "`" + `
	Message string ` + "`" + `json:"message"` + "`" + `
}

// QueryListSchema 列表查询结构体
type QueryListSchema struct {
	TotalCount  int64       ` + "`" + `json:"total_count"` + "`" + `  // 请求资源总计数
	FilterCount int64       ` + "`" + `json:"filter_count"` + "`" + ` // 请求资源过滤计数
	Records     interface{} ` + "`" + `json:"records"` + "`" + `      // 资源列表数据
}
`

const FiberSchemaCommonFile = `package schema

import (
	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"reflect"
)

var validate *validator.Validate
var trans ut.Translator

type BindSchema uint8

const (
	Json BindSchema = iota
	Query
)

func init() {
	// 初始化翻译器
	trans, _ = ut.New(zh_Hans_CN.New()).GetTranslator("zh")
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("label")
	})
	if err := zh.RegisterDefaultTranslations(validate, trans); err != nil {
		zap.L().Error("register validator failed", zap.Error(err))
	}
}

// Validator 参数验证器
func Validator(data interface{}, c *fiber.Ctx, t BindSchema) (bindErr error, err error) {
	switch t {
	case Json:
		err = c.BodyParser(data)
	case Query:
		err = c.QueryParser(data)
	}
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			ErrorResponse{Message: err.Error()}), err
	}
	if err = validate.Struct(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: err.(validator.ValidationErrors)[0].Translate(trans)}), err
	}
	return nil, nil
}

// SuccessResponse 完成响应结构体
type SuccessResponse struct {
	Status bool        ` + "`" + `json:"status"` + "`" + `
	Data   interface{} ` + "`" + `json:"data"` + "`" + `
}

// ErrorResponse 错误响应结构体
type ErrorResponse struct {
	Status  bool   ` + "`" + `json:"status"` + "`" + `
	Message string ` + "`" + `json:"message"` + "`" + `
}

// QueryListSchema 列表查询结构体
type QueryListSchema struct {
	TotalCount  int64       ` + "`" + `json:"total_count"` + "`" + `  // 请求资源总计数
	FilterCount int64       ` + "`" + `json:"filter_count"` + "`" + ` // 请求资源过滤计数
	Records     interface{} ` + "`" + `json:"records"` + "`" + `      // 资源列表数据
}
`

const GinSchemaTaskFile = `package schema

// TaskCreate 创建任务请求参数
type TaskCreate struct {
	Title string ` + "`" + `json:"title" validate:"required,min=4,max=100" label:"任务标题"` + "`" + `
}

// TaskList 任务列表请求参数
type TaskList struct {
	Page  int ` + "`" + `form:"page" validate:"required,min=1" label:"请求页码"` + "`" + `
	Size  int ` + "`" + `form:"size" validate:"required,min=1,max=50" label:"请求数量"` + "`" + `
	Title string
}

// TaskIDSchema 任务ID公用请求参数
type TaskIDSchema struct {
	ID uint ` + "`" + `json:"id" form:"id" validate:"required"` + "`" + `
}
`

const FiberSchemaTaskFile = `package schema

// TaskCreate 创建任务请求参数
type TaskCreate struct {
	Title string ` + "`" + `json:"title" validate:"required,min=4,max=100" label:"任务标题"` + "`" + `
}

// TaskList 任务列表请求参数
type TaskList struct {
	Page  int ` + "`" + `json:"page" validate:"required,min=1" label:"请求页码"` + "`" + `
	Size  int ` + "`" + `json:"size" validate:"required,min=1,max=50" label:"请求数量"` + "`" + `
	Title string
}

// TaskIDSchema 任务ID公用请求参数
type TaskIDSchema struct {
	ID uint ` + "`" + `json:"id" validate:"required"` + "`" + `
}
`
