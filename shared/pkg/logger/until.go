package logger

import (
	"fmt"
	"log/syslog"
)

func GetSysLogger(config LoggerOutputConfig, appName string) (w *syslog.Writer, err error) {
	if config.OutputType != "syslog" {
		err = fmt.Errorf("invalid syslog config")

		return nil, err
	}

	w, err = syslog.New(config.Priority, appName)

	if err != nil {
		w = nil
	}

	return w, err
}
