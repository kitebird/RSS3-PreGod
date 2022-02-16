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
