package config

import (
	"time"

	"github.com/sirupsen/logrus"
)

func InitLogging() {
	if Env.Logging.Level == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else if Env.Logging.Level == "info" {
		logrus.SetLevel(logrus.InfoLevel)
	}

	if Env.Logging.Format == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else if Env.Logging.Format == "text" {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}


	// For catching Fatal errors, we are waiting before exist for sending error to sentry successfully.
	logrus.RegisterExitHandler(func() {
		time.Sleep(3 * time.Second)
	})
}

