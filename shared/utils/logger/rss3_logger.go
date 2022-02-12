package logger

// Some simple encapsulations are made for the upper layer.
// The Sugare mode of the zap library is used by default.
// You can customize the encapsulation here to use other log libraries.
func Logger(config LoggerConfig) (log LoggerCore, err error) {
	log = nil
	err = nil

	if config.LoggerType == "" {
		config.LoggerType = "zap"
	}

	if config.LoggerType == "zap" {
		logger, err := GetZapLogger(config)
		if err == nil {
			log = &ZapLogger{
				SugaredLogger: *logger,
			}
		}
	}

	return log, err
}
