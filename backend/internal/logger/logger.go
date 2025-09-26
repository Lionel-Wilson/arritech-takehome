package logger

import (
	"github.com/Lionel-Wilson/arritech-takehome/internal/config"
	"github.com/sirupsen/logrus"
)

func New(cfg *config.Config) *logrus.Logger {
	logger := logrus.New()

	if cfg.Env == "prod" {
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{})
	}

	return logger
}
