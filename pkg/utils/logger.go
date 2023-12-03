package utils

import (
	"time"

	"github.com/sirupsen/logrus"
)

func LoggerInit() *logrus.Logger {
	log := logrus.New()

	if GetEnv() == "production" {
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC1123,
		})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
			TimestampFormat: time.RFC1123,
		})
	}

	return log
}