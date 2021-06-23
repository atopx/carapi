package model

const PublicConstFile = `package public

const (
	TimeFormatDay    = "2006-01-02"          // 固定format时间，2006-12345
	TimeFormatSecond = "2006-01-02 15:04:05" // 固定format时间，2006-12345
)
`
const PublicVariableFile = `package public

import (
	"app/config"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *config.Cfg
	Logger *zap.Logger
	Db     *gorm.DB
)
`
