package logger

import (
	"context"
	"errors"
	"time"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/logger"
	logger_model "github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/logger/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

type Logger struct {
	CoreLogger                logger_model.LoggerCore
	LogLevel                  gorm_logger.LogLevel
	SlowThreshold             time.Duration
	SkipCallerLookup          bool
	IgnoreRecordNotFoundError bool
}

func New() Logger {
	return Logger{
		CoreLogger:                logger.Logger,
		LogLevel:                  gorm_logger.Warn,
		SlowThreshold:             100 * time.Millisecond,
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: false,
	}
}

func (l Logger) SetAsDefault() {
	gorm_logger.Default = l
}

func (l Logger) LogMode(level gorm_logger.LogLevel) gorm_logger.Interface {
	return Logger{
		CoreLogger:                l.CoreLogger,
		SlowThreshold:             l.SlowThreshold,
		LogLevel:                  level,
		SkipCallerLookup:          l.SkipCallerLookup,
		IgnoreRecordNotFoundError: l.IgnoreRecordNotFoundError,
	}
}

func (l Logger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gorm_logger.Info {
		return
	}

	l.CoreLogger.Debugf(str, args...)
}

func (l Logger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gorm_logger.Warn {
		return
	}

	l.CoreLogger.Warnf(str, args...)
}

func (l Logger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gorm_logger.Error {
		return
	}

	l.CoreLogger.Errorf(str, args...)
}

func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}

	elapsed := time.Since(begin)

	switch {
	case err != nil && l.LogLevel >= gorm_logger.Error &&
		(!l.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := fc()
		l.CoreLogger.Error("trace",
			zap.Error(err), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.LogLevel >= gorm_logger.Warn:
		sql, rows := fc()
		l.CoreLogger.Warn("trace", zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	case l.LogLevel >= gorm_logger.Info:
		sql, rows := fc()
		l.CoreLogger.Debug("trace", zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	}
}
