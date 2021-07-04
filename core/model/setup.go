package model

const SetupConfigFile = `package setup

import (
	"app/public"
	"github.com/BurntSushi/toml"
)

func Config(filepath string) (err error) {
	if filepath == "" {
		filepath = "config.toml"
	}
	_, err = toml.DecodeFile(filepath, &public.Cfg)
	return err
}
`

const SetupDatabaseFile = `package setup

import (
	"app/config"
	"app/model"
	"app/public"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func migrate(db *gorm.DB) error {
	if !db.Migrator().HasTable("task") {
		return db.AutoMigrate(
			&model.Task{},
		)
	}
	return nil
}

func Database(cfg config.Database) (err error) {
	var sqlCfg = postgres.Config{
		DSN: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port,
		),
		PreferSimpleProtocol: true,
	}
	var gormCfg = &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.Lshortfile), logger.Config{
			SlowThreshold: cfg.SlowThreshold, LogLevel: cfg.LogLevel, Colorful: true,
		}),
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	}
	if public.Db, err = gorm.Open(postgres.New(sqlCfg), gormCfg); err != nil {
		return err
	}
	db, _ := public.Db.DB()
	db.SetMaxIdleConns(cfg.MaxIdleConn)
	db.SetMaxOpenConns(cfg.MaxOpenConn)
	db.SetConnMaxLifetime(30 * time.Second)
	return migrate(public.Db)
}
`

const SetupLoggerFile = `package setup

import (
	"app/config"
	"io"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// setupLogger 初始化日志配置
func setupLogger(level string, writer io.Writer, encoder zapcore.Encoder) (err error) {
	var loggerLevel = new(zapcore.Level)
	err = loggerLevel.UnmarshalText([]byte(level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, zapcore.AddSync(writer), loggerLevel)
	// 替换zap包中全局的logger实例，后续直接调 global.Logger
	zap.ReplaceGlobals(zap.New(core, zap.AddCaller()))
	return err
}

// DebugLogger 日志设置为控制台标准输出
func DebugLogger(cfg config.Logger) (err error) {
	return setupLogger(
		cfg.Level,
		os.Stdout,
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
	)
}

// ReleaseLogger 日志格式化为json并输出到日志
func ReleaseLogger(cfg config.Logger) (err error) {
	return setupLogger(
		cfg.Level,
		&lumberjack.Logger{
			Filename:   cfg.Filepath,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackup,
			MaxAge:     cfg.MaxAge,
		},
		zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}),
	)
}
`

const GinSetupRouterFile = `package setup

import (
	"app/api"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	// 跟路由
	engine.GET("/ping", api.PingApi)

	// 任务路由组
	var task = engine.Group("/task")
	{
		task.GET("/list", api.TaskList)
		task.POST("/create", api.TaskCreate)
		task.PATCH("/start", api.TaskStart)
		task.PATCH("/done", api.TaskDone)
		task.DELETE("/delete", api.TaskDelete)
	}
}
`

const FiberSetupRouterFile = `package setup

import (
	"app/api"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	// 跟路由
	app.Get("/ping", api.Ping)

	// 任务路由组
	task := app.Group("/task")
	task.Get("/list", api.TaskList)
	task.Post("/create", api.TaskCreate)
	task.Patch("/start", api.TaskStart)
	task.Patch("/done", api.TaskDone)
	task.Delete("/delete", api.TaskDelete)
}
`

const GinSetupEngineFile = `package setup

import (
	"app/docs"
	"app/middleware"
	"app/public"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// InitEngine 初始化服务引擎
func InitEngine() *gin.Engine {
	var err error
	var engine = gin.New()
	gin.SetMode(public.Cfg.Server.Mode)
	// 运行模式差异
	if public.Cfg.Server.Mode != gin.ReleaseMode {
		// debug模式下日志配置
		if err = DebugLogger(public.Cfg.Logger); err != nil {
			log.Panicf("初始化日志系统异常: %s", err)
		}
		// 启用swagger api文档
		docs.SwaggerInfo.Host = public.Cfg.Common.SwaggerHost + ":" + public.Cfg.Server.Port
		engine.GET("/swagger/*any", middleware.Swagger(fmt.Sprintf("http://%s/swagger/doc.json", docs.SwaggerInfo.Host)))
		log.Println("===============Config===============")
		log.Printf("运行模式: %s\n", public.Cfg.Server.Mode)
		log.Printf("运行地址: %s\n", fmt.Sprintf("http://%s", docs.SwaggerInfo.Host))
		log.Printf("请求超时: %s\n", public.Cfg.Server.ReadTimeout*time.Second)
		log.Printf("响应超时: %s\n", public.Cfg.Server.WriteTimeout*time.Second)
		if strings.HasPrefix(public.Cfg.Common.SentryDSN, "http") {
			log.Printf("Sentry: %s\n", public.Cfg.Common.SentryDSN)
		}
		log.Printf("Swagger: %s\n", fmt.Sprintf("http://%s/swagger/index.html", docs.SwaggerInfo.Host))
		log.Println("====================================")
	} else {
		// release模式下日志配置
		if err = ReleaseLogger(public.Cfg.Logger); err != nil {
			log.Panicf("初始化日志系统异常: %s", err)
		}
	}

	// 初始化数据库连接
	if err = Database(public.Cfg.Database); err != nil {
		log.Panicf("初始化日志系统异常: %s", err)
	}

	// 引用中间件
	engine.Use(middleware.Cors())
	engine.Use(middleware.GinLogger())
	engine.Use(middleware.GinRecovery(true))

	// 初始化路由
	Router(engine)
	return engine
}
`

const FiberSetupEngineFile = `package setup

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/zap"
	"os"
	"app/public"
)

func Engine() *fiber.App {
	var err error
	switch public.Cfg.Server.Mode {
	case "debug", "DEBUG":
		err = DebugLogger(public.Cfg.Logger)
	default:
		err = ReleaseLogger(public.Cfg.Logger)
	}
	if err != nil {
		zap.L().Error("setup logger", zap.Error(err))
		os.Exit(0)
	}
	if err = Database(public.Cfg.Database); err != nil {
		zap.L().Error("setup database", zap.Error(err))
		os.Exit(0)
	}
	var engine = fiber.New()
	engine.Use(logger.New(), cors.New(), cache.New())
	Router(engine)
	return engine
}
`
