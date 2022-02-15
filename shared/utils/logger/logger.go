package logger

import "log/syslog"

type LoggerOutputConfig struct {
	OutputType string

	// The outputType is "file"
	Filepath string

	// The outputType is "syslog"
	Priority syslog.Priority
}

type LoggerConfig struct {
	AppName string

	// Basic log configuration
	LoggerType string
	Level      string
	Encoding   string

	// Configurable to options such as syslog or stdout
	Output []LoggerOutputConfig
}

type StandardLogger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
}

type LoggerCore interface {
	StandardLogger
}
