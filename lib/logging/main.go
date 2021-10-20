package logging

import (
	"go.uber.org/zap"
	"tawny/lib/configuration"
)

func InitLogger(c configuration.Configuration) (logger *zap.Logger) {
	if c.IsDev {
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()

	}
	return
}
