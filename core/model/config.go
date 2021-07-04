package model

const GinConfigGoFile = `package config

import (
	"time"

	"gorm.io/gorm/logger"
)

type Config struct {
	Server   Server      ` + "`" + `toml:"server"` + "`" + `
	Database Database    ` + "`" + `toml:"database"` + "`" + `
	Logger   Logger      ` + "`" + `toml:"logger"` + "`" + `
	Common   Common      ` + "`" + `toml:"common"` + "`" + `
}

type Common struct {
	SentryDSN   string    ` + "`" + `toml:"sentry_dsn"` + "`" + `
	SwaggerHost string    ` + "`" + `toml:"swagger_host"` + "`" + `
}

type Server struct {
	Mode         string        ` + "`" + `toml:"mode"` + "`" + `
	Port         string        ` + "`" + `toml:"port"` + "`" + `
	ReadTimeout  time.Duration ` + "`" + `toml:"read_timeout"` + "`" + `
	WriteTimeout time.Duration ` + "`" + `toml:"write_timeout"` + "`" + `
}

type Database struct {
	User          string          ` + "`" + `toml:"user"` + "`" + `
	Password      string          ` + "`" + `toml:"password"` + "`" + `
	DBName        string          ` + "`" + `toml:"dbname"` + "`" + `
	Host          string          ` + "`" + `toml:"host"` + "`" + `
	Port          int             ` + "`" + `toml:"port"` + "`" + `
	MaxIdleConn   int             ` + "`" + `toml:"max_idle_conn"` + "`" + `
	MaxOpenConn   int             ` + "`" + `toml:"max_open_conn"` + "`" + `
	LogLevel      logger.LogLevel ` + "`" + `toml:"log_level"` + "`" + `
	SlowThreshold time.Duration   ` + "`" + `toml:"threshold"` + "`" + `
}

type Logger struct {
	MaxSize   int    ` + "`" + `toml:"maxsize"` + "`" + `
	MaxAge    int    ` + "`" + `toml:"maxage"` + "`" + `
	MaxBackup int    ` + "`" + `toml:"backup"` + "`" + `
	Level     string ` + "`" + `toml:"level"` + "`" + `
	Filepath  string ` + "`" + `toml:"filepath"` + "`" + `
}
`

const GinConfigTomlFile = `[server]
mode = "debug"
port = "9000"
read_timeout = 60
write_timeout = 60

[common]
swagger_host = "127.0.0.1"

[database]
user = "postgres"
password = "postgres"
dbname = "todo"
host = "172.20.88.240"
port = 5432
max_idle_conn = 10
max_open_conn = 20
log_level = 2
threshold = 1000

[logger]
maxsize = 100
maxage = 7
backup = 10
level = "debug"
filepath = "logs/app.log"
`

const FiberConfigGoFile = `package config

import (
	"gorm.io/gorm/logger"
	"time"
)

type Config struct {
	Server   Server   ` + "`" + `toml:"server"` + "`" + `
	Database Database ` + "`" + `toml:"database"` + "`" + `
	Logger   Logger   ` + "`" + `toml:"logger"` + "`" + `
}

type Server struct {
	Mode string ` + "`" + `toml:"mode"` + "`" + `
	Port int    ` + "`" + `toml:"port"` + "`" + `
}

type Database struct {
	User          string          ` + "`" + `toml:"user"` + "`" + `
	Password      string          ` + "`" + `toml:"password"` + "`" + `
	DBName        string          ` + "`" + `toml:"dbname"` + "`" + `
	Host          string          ` + "`" + `toml:"host"` + "`" + `
	Port          int             ` + "`" + `toml:"port"` + "`" + `
	MaxIdleConn   int             ` + "`" + `toml:"max_idle_conn"` + "`" + `
	MaxOpenConn   int             ` + "`" + `toml:"max_open_conn"` + "`" + `
	LogLevel      logger.LogLevel ` + "`" + `toml:"log_level"` + "`" + `
	SlowThreshold time.Duration   ` + "`" + `toml:"threshold"` + "`" + `
}

type Logger struct {
	MaxSize   int    ` + "`" + `toml:"maxsize"` + "`" + `
	MaxAge    int    ` + "`" + `toml:"maxage"` + "`" + `
	MaxBackup int    ` + "`" + `toml:"backup"` + "`" + `
	Level     string ` + "`" + `toml:"level"` + "`" + `
	Filepath  string ` + "`" + `toml:"filepath"` + "`" + `
}
`

const FiberConfigTomlFile = `[server]
mode = "debug"
port = 9000

[database]
user = "postgres"
password = "postgres"
dbname = "todo"
host = "172.20.88.240"
port = 5432
max_idle_conn = 10
max_open_conn = 20
log_level = 2
threshold = 1000

[logger]
maxsize = 100
maxage = 7
backup = 10
level = "debug"
filepath = "logs/app.log"
`
