package util

import (
	"fmt"
	"log/syslog"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/config"
)

func GetSysLogger(c config.LoggerOutputConfig, prefix string) (w *syslog.Writer, err error) {
	if c.Type != "syslog" {
		err = fmt.Errorf("invalid syslog config")

		return nil, err
	}

	w, err = syslog.New(c.Priority, prefix)

	if err != nil {
		w = nil
	}

	return w, err
}
