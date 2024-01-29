package xlogger

import (
	"log"
	"os"
)

func WriteToFile(message, path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileLogger := log.New(file, "", log.Ldate|log.Ltime)
	fileLogger.Printf(message)
	return nil
}
