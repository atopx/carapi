package model

const GinApiTestFile = `// 服务测试接口
package api

import (
	"app/schema"

	"github.com/gin-gonic/gin"
)

// @Summary 服务连通性测试
// @Description 服务连通性测试接口
// @Tags 测试
// @Success 200 {object} schema.SuccessResponse
// @Failure 400 {object} schema.ErrorResponse
// @Router /ping [get]
func PingApi(c *gin.Context) {
	c.JSON(200, schema.SuccessResponse{Status: true, Data: "pong"})
}
`

const FiberApiTestFile = `package api

import (
	"app/schema"
	"github.com/gofiber/fiber/v2"
)

// @Summary 服务连通性测试
// @Description 服务连通性测试接口
// @Tags 测试
// @Success 200 {object} schema.SuccessResponse
// @Failure 400 {object} schema.ErrorResponse
// @Router /ping [get]
func Ping(c *fiber.Ctx) error {
	return c.JSON(&schema.SuccessResponse{Status: true, Data: "pong"})
}
`

const GinApiTaskFile = `package api

import (
	"app/schema"
	"app/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

// @Summary 列表
// @Description 查询任务列表
// @Tags 任务管理
// @Accept application/json
// @Produce application/json
// @param page query int true "请求页" minimum(1) default(1)
// @param size query int true "请求数" minimum(1) maximum(50) default(10)
// @param title query string false "标题模糊搜索"
// @Success 200 {object} schema.SuccessResponse
// @Failure 400 {object} schema.ErrorResponse
// @Router /task/list [get]
func TaskList(c *gin.Context) {
	var s schema.TaskList
	if err := schema.Validator(&s, c, binding.Query); err == nil {
		c.JSON(http.StatusOK, schema.SuccessResponse{Status: true, Data: service.TaskList(s)})
	}
}

// @Summary 新建
// @Description 新建任务
// @Tags 任务管理
// @Accept application/json
// @Produce application/json
// @param param body schema.TaskCreate true "请求参数,字段说明点击model"
// @Success 200 {object} schema.SuccessResponse
// @Failure 400 {object} schema.ErrorResponse
// @Router /task/create [post]
func TaskCreate(c *gin.Context) {
	var s schema.TaskCreate
	if err := schema.Validator(&s, c, binding.JSON); err == nil {
		fmt.Println(s)
		if data, err := service.TaskCreate(s); err == nil {
			c.JSON(http.StatusOK, schema.SuccessResponse{Status: true, Data: data})
		} else {
			c.JSON(http.StatusCreated, schema.ErrorResponse{Message: err.Error()})
		}
	}
}

// @Summary 开始
// @Description 标记任务开始
// @Tags 任务管理
// @Accept application/json
// @Produce application/json
// @param id query int true "任务id" minimum(1) default(1)
// @Success 200 {object} schema.SuccessResponse
// @Failure 400 {object} schema.ErrorResponse
// @Router /task/start [patch]
func TaskStart(c *gin.Context) {
	var s schema.TaskIDSchema
	if err := schema.Validator(&s, c, binding.Query); err == nil {
		if data, err := service.TaskStart(s); err == nil {
			c.JSON(http.StatusOK, schema.SuccessResponse{Status: true, Data: data})
		} else {
			c.JSON(http.StatusBadRequest, schema.ErrorResponse{Message: err.Error()})
		}
	}
}

// @Summary 完成
// @Description 标记任务完成
// @Tags 任务管理
// @Accept application/json
// @Produce application/json
// @param id query int true "任务id" minimum(1) default(1)
// @Success 200 {object} schema.SuccessResponse
// @Failure 400 {object} schema.ErrorResponse
// @Router /task/done [patch]
func TaskDone(c *gin.Context) {
	var s schema.TaskIDSchema
	if err := schema.Validator(&s, c, binding.Query); err == nil {
		if data, err := service.TaskDone(s); err == nil {
			c.JSON(http.StatusOK, schema.SuccessResponse{Status: true, Data: data})
		} else {
			c.JSON(http.StatusBadRequest, schema.ErrorResponse{Message: err.Error()})
		}
	}
}

// @Summary 删除
// @Description 删除任务
// @Tags 任务管理
// @Accept application/json
// @Produce application/json
// @param param body schema.TaskIDSchema true "请求参数,字段说明点击model"
// @Success 200 {object} schema.SuccessResponse
// @Failure 400 {object} schema.ErrorResponse
// @Router /task/delete [delete]
func TaskDelete(c *gin.Context) {
	var s schema.TaskIDSchema
	if err := schema.Validator(&s, c, binding.JSON); err == nil {
		if err = service.TaskDelete(s); err == nil {
			c.JSON(http.StatusOK, schema.SuccessResponse{Status: true, Data: nil})
		} else {
			c.JSON(http.StatusBadRequest, schema.ErrorResponse{Message: err.Error()})
		}
	}
}
`

const FiberApiTaskFile = `package api

import (
	"app/schema"
	"app/service"

	"github.com/gofiber/fiber/v2"
)

// @Summary 列表
// @Description 查询任务列表
// @Tags 任务管理
// @Accept application/json
// @Produce application/json
// @param page query int true "请求页" minimum(1) default(1)
// @param size query int true "请求数" minimum(1) maximum(50) default(10)
// @param title query string false "标题模糊搜索"
// @Success 200 {object} schema.SuccessResponse
// @Failure 400 {object} schema.ErrorResponse
// @Router /task/list [get]
func TaskList(c *fiber.Ctx) error {
	var s schema.TaskList
	if r, e := schema.Validator(&s, c, schema.Query); e != nil {
		return r
	}
	return c.JSON(schema.SuccessResponse{Status: true, Data: service.TaskList(s)})
}

// @Summary 新建
// @Description 新建任务
// @Tags 任务管理
// @Accept application/json
// @Produce application/json
// @param param body schema.TaskCreate true "请求参数,字段说明点击model"
// @Success 200 {object} schema.SuccessResponse
// @Failure 400 {object} schema.ErrorResponse
// @Router /task/create [post]
func TaskCreate(c *fiber.Ctx) error {
	var s schema.TaskCreate
	if r, e := schema.Validator(&s, c, schema.Json); e != nil {
		return r
	}
	data, err := service.TaskCreate(s)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.ErrorResponse{Message: err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(schema.SuccessResponse{Status: true, Data: data})
}

// @Summary 开始
// @Description 标记任务开始
// @Tags 任务管理
// @Accept application/json
// @Produce application/json
// @param id query int true "任务id" minimum(1) default(1)
// @Success 200 {object} schema.SuccessResponse
// @Failure 400 {object} schema.ErrorResponse
// @Router /task/start [patch]
func TaskStart(c *fiber.Ctx) error {
	var s schema.TaskIDSchema
	if r, e := schema.Validator(&s, c, schema.Query); e != nil {
		return r
	}
	data, err := service.TaskStart(s)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.ErrorResponse{Message: err.Error()})
	}
	return c.JSON(schema.SuccessResponse{Status: true, Data: data})
}

// @Summary 完成
// @Description 标记任务完成
// @Tags 任务管理
// @Accept application/json
// @Produce application/json
// @param id query int true "任务id" minimum(1) default(1)
// @Success 200 {object} schema.SuccessResponse
// @Failure 400 {object} schema.ErrorResponse
// @Router /task/done [patch]
func TaskDone(c *fiber.Ctx) error {
	var s schema.TaskIDSchema
	if r, e := schema.Validator(&s, c, schema.Query); e != nil {
		return r
	}
	data, err := service.TaskDone(s)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.ErrorResponse{Message: err.Error()})
	}
	return c.JSON(schema.SuccessResponse{Status: true, Data: data})
}

// @Summary 删除
// @Description 删除任务
// @Tags 任务管理
// @Accept application/json
// @Produce application/json
// @param param body schema.TaskIDSchema true "请求参数,字段说明点击model"
// @Success 200 {object} schema.SuccessResponse
// @Failure 400 {object} schema.ErrorResponse
// @Router /task/delete [delete]
func TaskDelete(c *fiber.Ctx) error {
	var s schema.TaskIDSchema
	if r, e := schema.Validator(&s, c, schema.Json); e != nil {
		return r
	}
	if err := service.TaskDelete(s); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.ErrorResponse{Message: err.Error()})
	}
	return c.Status(fiber.StatusNoContent).JSON(schema.SuccessResponse{Status: true, Data: nil})
}
`
