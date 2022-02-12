package logger

import (
	"fmt"
	"log/syslog"
)

func GetSysLogger(config LoggerOutputConfig) (w *syslog.Writer, err error) {
	if config.OutputType == "" ||
		config.OutputType != "syslog" ||
		config.Network == "" ||
		config.Ipv4Addr == "" ||
		config.Port == "" {
		err = fmt.Errorf("invalid syslog config")
		return nil, err
	}

	w, err = syslog.Dial(config.Network,
		config.Ipv4Addr+":"+config.Port,
		syslog.LOG_INFO,
		"")

	if err != nil {
		w = nil
	}

	return w, err
}
