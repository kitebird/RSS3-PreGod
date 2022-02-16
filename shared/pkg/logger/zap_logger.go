package logger

import (
	"log/syslog"
	"net/url"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type sysLogSink struct {
	sysLogWriter *syslog.Writer
}

func (s sysLogSink) Sync() error {
	return nil
}

func (s sysLogSink) Close() error {
	return nil
}

func (s sysLogSink) Write(p []byte) (n int, err error) {
	s.sysLogWriter.Write(p)

	return len(p), nil
}

type ZapLogger struct {
	zap.SugaredLogger
}

func GetZapLogger(loggerConfig LoggerConfig) (*zap.SugaredLogger, error) {
	outputPaths, err := parseOutputPaths(loggerConfig)
	if err != nil {
		return nil, err
	}

	if len(outputPaths) == 0 {
		outputPaths = append(outputPaths, "stdout")
	}

	zapConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(getLogLevel(loggerConfig.Level)),
		Development:      true,
		Encoding:         loggerConfig.Encoding,
		EncoderConfig:    getDefaultEncoderCfg(),
		InitialFields:    map[string]interface{}{"app": loggerConfig.AppName},
		OutputPaths:      outputPaths,
		ErrorOutputPaths: []string{},
	}

	zaplog, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}

	sugaredLogger := zaplog.Sugar()

	return sugaredLogger, nil
}

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.DebugLevel
	}
}

func getDefaultEncoderCfg() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "severity",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		FunctionKey:    zapcore.OmitKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func parseOutputPaths(loggerConfig LoggerConfig) ([]string, error) {
	var outputPaths []string

	for _, outputConfig := range loggerConfig.Output {
		if outputConfig.OutputType == "stdout" ||
			outputConfig.OutputType == "stderr" {
			outputPaths = append(outputPaths, outputConfig.OutputType)
		} else if outputConfig.OutputType == "file" {
			outputPaths = append(outputPaths, outputConfig.Filepath)
		} else if outputConfig.OutputType == "syslog" {
			outputPaths = append(outputPaths, "Syslog://127.0.0.1")
			sysLogFactory := func(u *url.URL) (zap.Sink, error) {
				w, err := GetSysLogger(outputConfig, loggerConfig.AppName)

				if err != nil {
					return nil, err
				}

				s := sysLogSink{
					sysLogWriter: w,
				}

				return s, nil
			}
			if err := zap.RegisterSink("Syslog", sysLogFactory); err != nil {
				return nil, err
			}
		}
	}

	return outputPaths, nil
}
