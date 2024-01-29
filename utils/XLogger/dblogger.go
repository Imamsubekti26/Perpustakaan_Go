package xlogger

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DbLogger() (*gorm.Config, error) {
	today := time.Now().Format("2006-01-02")
	filePath := fmt.Sprintf("logs/%s-db.log", today)
	logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, Error(err)
	}
	logWriter := log.New(logFile, "\r\n", log.LstdFlags)

	return &gorm.Config{
		Logger: logger.New(
			logWriter,
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		),
	}, nil
}
