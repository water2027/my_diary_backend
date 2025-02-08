package utils

import (
	"log"
	"os"
)

func InitLog() error {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Unable to open log file: %v", err)
		return err
	}

	log.SetOutput(logFile)
	return nil
}

func LogError(err *error) {
	if *err != nil {
		log.Printf("Error: %v", *err)
	}
}