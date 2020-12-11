package quark

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func (a *App) SetupLogger() {
	// Setup logger
	logFile := os.Getenv("LOG_FILE")
	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal("Logger: ", err)
		}

		log.SetFormatter(&log.JSONFormatter{})

		mw := io.MultiWriter(os.Stdout, file)
		logrus.SetOutput(mw)
	}
}
