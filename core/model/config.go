package model

const ConfigGoFile = `package config

import (
	"time"

	"gorm.io/gorm/logger"
)

type Cfg struct {
	Server   ServerConfig   ` + "`" + `yaml:"server"` + "`" + `
	Database DatabaseConfig ` + "`" + `yaml:"database"` + "`" + `
	Logger   LoggerConfig   ` + "`" + `yaml:"logger"` + "`" + `
	Common   CommonConfig   ` + "`" + `yaml:"common"` + "`" + `
}

type CommonConfig struct {
	SentryDSN    string ` + "`" + `yaml:"sentry_dsn"` + "`" + `
	SwaggerHost string ` + "`" + `yaml:"swagger_host"` + "`" + `
}

type ServerConfig struct {
	RunMode      string        ` + "`" + `yaml:"run_mode"` + "`" + `
	RunPort      string        ` + "`" + `yaml:"run_port"` + "`" + `
	ReadTimeout  time.Duration ` + "`" + `yaml:"read_timeout"` + "`" + `
	WriteTimeout time.Duration ` + "`" + `yaml:"write_timeout"` + "`" + `
}

type LoggerConfig struct {
	MaxSize   int    ` + "`" + `yaml:"maxsize"` + "`" + `
	MaxAge    int    ` + "`" + `yaml:"maxage"` + "`" + `
	MaxBackup int    ` + "`" + `yaml:"max_backup"` + "`" + `
	Level     string ` + "`" + `yaml:"level"` + "`" + `
	Filepath  string ` + "`" + `yaml:"filepath"` + "`" + `
}

type DatabaseConfig struct {
	Host                 string          ` + "`" + `yaml:"host"` + "`" + `
	Port                 string          ` + "`" + `yaml:"port"` + "`" + `
	DbName               string          ` + "`" + `yaml:"db_name"` + "`" + `
	SSLMode              string          ` + "`" + `yaml:"sslmode"` + "`" + `
	TimeZone             string          ` + "`" + `yaml:"timezone"` + "`" + `
	Username             string          ` + "`" + `yaml:"username"` + "`" + `
	Password             string          ` + "`" + `yaml:"password"` + "`" + `
	PreferSimpleProtocol bool            ` + "`" + `yaml:"prefer_simple_protocol"` + "`" + `
	MaxIdleConns         int             ` + "`" + `yaml:"max_idle_conns"` + "`" + `
	MaxOpenConns         int             ` + "`" + `yaml:"max_open_conns"` + "`" + `
	LogLevel             logger.LogLevel ` + "`" + `yaml:"log_level"` + "`" + `
	SlowThreshold        time.Duration   ` + "`" + `yaml:"slow_threshold"` + "`" + `
}
`

const ConfigYamlFile = `server:
  run_mode: 'debug'
  run_port: '9404'
  read_timeout: 60
  write_timeout: 60

common:
  sentry_dsn: ''
  swagger_host: '127.0.0.1'

logger:
  maxsize: 100
  max_age: 7
  level: debug
  max_backup: 1
  filepath: logs/app.log

database:
  host: 127.0.0.1
  port: 5432
  db_name: ginhelper
  username: postgres
  password: postgres
  sslmode: disable
  timezone: Asia/Shanghai
  prefer_simple_protocol: true
  max_idle_conns: 10
  max_open_conns: 20
  log_level: 3   # 1:silent, 2:error, 3:warn; 4:info
  slow_threshold: 1000 # 慢SQL记录(毫秒)
`
