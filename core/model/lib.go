package model

const LibOrgLoggerFile = `package libs

import (
	"context"
	"errors"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"gorm.io/gorm/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GormLogger struct {
	ZapLogger                 *zap.Logger
	LogLevel                  logger.LogLevel
	SlowThreshold             time.Duration
	SkipCallerLookup          bool
	IgnoreRecordNotFoundError bool
}

type GormLoggerConfig struct {
	SlowThreshold time.Duration
	LogLevel      logger.LogLevel
}

func GormLoggerNew(zapLogger *zap.Logger, config *GormLoggerConfig) GormLogger {
	if config.SlowThreshold == 0 {
		config.SlowThreshold = 500 * time.Millisecond
	}
	return GormLogger{
		ZapLogger:                 zapLogger,
		LogLevel:                  config.LogLevel,
		SlowThreshold:             config.SlowThreshold,
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: true,
	}
}

func (l GormLogger) SetAsDefault() {
	logger.Default = l
}

func (l GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return GormLogger{
		ZapLogger:                 l.ZapLogger,
		SlowThreshold:             l.SlowThreshold,
		LogLevel:                  level,
		SkipCallerLookup:          l.SkipCallerLookup,
		IgnoreRecordNotFoundError: l.IgnoreRecordNotFoundError,
	}
}

func (l GormLogger) Info(_ context.Context, str string, args ...interface{}) {
	if l.LogLevel < logger.Info {
		return
	}
	l.logger().Sugar().Debugf(str, args...)
}

func (l GormLogger) Warn(_ context.Context, str string, args ...interface{}) {
	if l.LogLevel < logger.Warn {
		return
	}
	l.logger().Sugar().Warnf(str, args...)
}

func (l GormLogger) Error(_ context.Context, str string, args ...interface{}) {
	if l.LogLevel < logger.Error {
		return
	}
	l.logger().Sugar().Errorf(str, args...)
}

func (l GormLogger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}
	cost := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= logger.Error && (!l.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := fc()
		l.logger().Error(sql, zap.Error(err), zap.Duration("cost", cost), zap.Int64("rows", rows))
	case l.SlowThreshold != 0 && cost > l.SlowThreshold && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		l.logger().Warn(sql, zap.Duration("cost", cost), zap.Int64("rows", rows))
	case l.LogLevel >= logger.Info:
		sql, rows := fc()
		l.logger().Debug(sql, zap.Duration("cost", cost), zap.Int64("rows", rows))
	}
}

var gormPackage = filepath.Join("gorm.io", "gorm")

func (l GormLogger) logger() *zap.Logger {
	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		if ok && !strings.Contains(file, gormPackage) {
			return l.ZapLogger.WithOptions(zap.AddCallerSkip(i))
		}
	}
	return l.ZapLogger
}
`
