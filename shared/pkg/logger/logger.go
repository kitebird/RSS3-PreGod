package logger

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/logger/model"
)

var log model.LoggerCore
var isSetup = false

func Logger() model.LoggerCore {
	// Only possible when calling for initialization,
	// if an exception occurs, the process will be closed directly
	if !isSetup {
		config := defaultLoggerConfig()
		logger, err := GetZapLogger(config)

		if err != nil {
			panic(err)
		}

		log = &ZapLogger{SugaredLogger: *logger}
	}

	return log
}

// Some simple encapsulations are made for the upper layer.
// The Sugare mode of the zap library is used by default.
// You can customize the encapsulation here to use other log libraries.
func Setup() error {
	// Since there is no suitable nested configuration,
	// the configuration of the logger is writtn to death first,
	// and then called using the configuration file.
	config := LoggerConfig{
		AppName:    "PreGod",
		LoggerType: "zap",
		Level:      "debug",
		Encoding:   "json",
		Output: []LoggerOutputConfig{
			{
				OutputType: "stdout",
			},
			// {
			// 	OutputType: "file",
			// 	Filepath:   "/root/rss3_dev/RSS3-PreGod/PreGod.log",
			// },
			// {
			// 	OutputType: "syslog",
			// },
		},
	}

	if config.LoggerType == "" {
		config.LoggerType = "zap"
	}

	if config.LoggerType == "zap" {
		logger, err := GetZapLogger(config)
		if err != nil {
			return err
		}

		log = &ZapLogger{SugaredLogger: *logger}
	}

	isSetup = true

	return nil
}

func defaultLoggerConfig() LoggerConfig {
	return LoggerConfig{
		AppName:    "PreGod",
		LoggerType: "zap",
		Level:      "debug",
		Encoding:   "json",
		Output: []LoggerOutputConfig{
			{
				OutputType: "stdout",
			},
		},
	}
}
