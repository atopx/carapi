package model

const ServerConfigFile = `package server

import (
	"app/public"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// LoadConfig 加载配置文件
func LoadConfig(config string) (err error) {
	file, err := ioutil.ReadFile(config)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(file, &public.Config)
	if err != nil {
		return err
	}
	return err
}
`

const ServerDatabaseFile = `package server

import (
	"app/config"
	"app/libs"
	"app/models"
	"app/public"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDatabaseData 初始化开发人员数据
func InitDatabaseData() {
	// TODO develop mock user data
	public.Db.Model(models.DevelopUser).FirstOrCreate(models.DevelopUser, "uuid = ?", models.DevelopUser.UUID)
	public.Logger.Info("Create develop user develop(c66f19b6-569d-4c85-94ae-5f5b45de18f0)")
}

// InitDatabase 初始化数据库连接
func InitDatabase(cfg *config.DatabaseConfig) (err error) {
	var psqlCfg = postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			cfg.Host, cfg.Username, cfg.Password, cfg.DbName, cfg.Port, cfg.SSLMode, cfg.TimeZone),
		PreferSimpleProtocol: cfg.PreferSimpleProtocol,
	}

	var gormCfg = &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: libs.GormLoggerNew(zap.L(), &libs.GormLoggerConfig{
			LogLevel: cfg.LogLevel, SlowThreshold: cfg.SlowThreshold * time.Millisecond}),
	}

	public.Db, err = gorm.Open(postgres.New(psqlCfg), gormCfg)
	if err != nil {
		return
	}
	sql, err := public.Db.DB()
	if err != nil {
		return
	}
	sql.SetMaxIdleConns(cfg.MaxIdleConns)
	sql.SetMaxOpenConns(cfg.MaxOpenConns)
	sql.SetConnMaxIdleTime(1 * time.Hour)
	InitDatabaseData()
	return err
}
`

const ServerEngineFile = `package server

import (
	"app/docs"
	"app/middleware"
	"app/public"
	"app/routers"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	sentrygin "github.com/getsentry/sentry-go/gin"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// InitEngine 初始化服务引擎
func InitEngine() *gin.Engine {
	var err error
	var engine = gin.New()
	gin.SetMode(public.Config.Server.RunMode)
	// 运行模式差异
	if public.Config.Server.RunMode != gin.ReleaseMode {
		// debug模式下日志配置
		err = InitDebugLogger(&public.Config.Logger)
		if err != nil {
			public.Logger.Error("初始化日志系统异常", zap.Any("err", err))
			os.Exit(0)
		}
		// 启用swagger api文档
		docs.SwaggerInfo.Host = public.Config.Common.SwaggerHost + ":" + public.Config.Server.RunPort
		engine.GET("/swagger/*any", middleware.Swagger(fmt.Sprintf("http://%s/swagger/doc.json", docs.SwaggerInfo.Host)))
		log.Println("===============Config===============")
		log.Printf("运行模式: %s\n", public.Config.Server.RunMode)
		log.Printf("运行地址: %s\n", fmt.Sprintf("http://%s", docs.SwaggerInfo.Host))
		log.Printf("请求超时: %s\n", public.Config.Server.ReadTimeout*time.Second)
		log.Printf("响应超时: %s\n", public.Config.Server.WriteTimeout*time.Second)
		if strings.HasPrefix(public.Config.Common.SentryDSN, "http") {
			log.Printf("Sentry: %s\n", public.Config.Common.SentryDSN)
		}
		log.Printf("Swagger: %s\n", fmt.Sprintf("http://%s/swagger/index.html", docs.SwaggerInfo.Host))
		log.Println("====================================")
	} else {
		// release模式下日志配置
		err = InitReleaseLogger(&public.Config.Logger)
		if err != nil {
			public.Logger.Error("初始化日志系统异常", zap.Any("err", err))
			os.Exit(0)
		}
		// release模式启用sentry连接
		err = InitSentry(public.Config.Common.SentryDSN)
		if err != nil {
			public.Logger.Error("初始化Sentry异常", zap.Any("err", err))
			os.Exit(0)
		}
		engine.Use(sentrygin.New(sentrygin.Options{}))
	}

	// 初始化数据库连接
	err = InitDatabase(&public.Config.Database)
	if err != nil {
		public.Logger.Error("数据库连接异常", zap.Any("err", err))
		os.Exit(0)
	}

	// 引用中间件
	engine.Use(middleware.Cors())
	engine.Use(middleware.GinLogger())
	engine.Use(middleware.GinRecovery(true))

	// 初始化路由
	routers.InitPingRouter(engine)

	return engine
}
`

const ServerLoggerFile = `package server

import (
	"app/config"
	"app/public"
	"io"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// initLogger 初始化日志配置
func initLogger(level string, writer io.Writer, encoder zapcore.Encoder) (err error) {
	var loggerLevel = new(zapcore.Level)
	err = loggerLevel.UnmarshalText([]byte(level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, zapcore.AddSync(writer), loggerLevel)
	// 替换zap包中全局的logger实例，后续直接调 global.Logger
	public.Logger = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(public.Logger)
	return err
}

// InitDebugLogger 日志设置为控制台标准输出
func InitDebugLogger(cfg *config.LoggerConfig) (err error) {
	return initLogger(
		cfg.Level,
		os.Stdout,
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
	)
}

// InitReleaseLogger 日志格式化为json并输出到日志
func InitReleaseLogger(cfg *config.LoggerConfig) (err error) {
	return initLogger(
		cfg.Level,
		NewReleaseWriter(cfg.Filepath, cfg.MaxSize, cfg.MaxBackup, cfg.MaxAge),
		zapcore.NewJSONEncoder(NewReleaseEncoderConfig()),
	)
}

// NewReleaseEncoderConfig 生产环境下的日志环境配置
func NewReleaseEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

}

// NewReleaseWriter 生产环境下，将日志输出到文件(自动分块)
func NewReleaseWriter(filename string, maxSize, maxBackup, maxAge int) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
}
`

const ServerSentryFile = `package server

import (
	"github.com/getsentry/sentry-go"
)

func InitSentry(dsn string) (err error) {
	return sentry.Init(sentry.ClientOptions{Dsn: dsn, Debug: false})
}
`
