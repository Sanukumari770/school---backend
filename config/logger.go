package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogger() {

	log.SetFormatter(&log.JSONFormatter{})

	file, err := os.OpenFile(
		"app.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)

	if err == nil {

		log.SetOutput(file)
	}
}