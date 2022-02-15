package logger

import (
	"fmt"
	"log/syslog"
)

func GetSysLogger(config LoggerOutputConfig) (w *syslog.Writer, err error) {
	if config.OutputType == "" ||
		config.OutputType != "syslog" ||
		config.OutputType != "file" {
		err = fmt.Errorf("invalid syslog config")
		return nil, err
	}

	w, err = syslog.New(config.Priority, "syslog")

	if err != nil {
		w = nil
	}

	return w, err
}
