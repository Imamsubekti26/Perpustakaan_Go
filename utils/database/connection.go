package database

import (
	"fmt"
	"log"
	"os"
	"time"

	xlogger "github.com/Imamsubekti26/Perpustakaan_Go/utils/XLogger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbInstance struct {
	Connection *gorm.DB
}

func Connection() (*dbInstance, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	gormConfig, err := dbLogger()
	if err != nil {
		return nil, xlogger.Error(err)
	}

	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return nil, xlogger.Error(err)
	}

	var dbInstance dbInstance
	dbInstance.Connection = db
	return &dbInstance, nil
}

func dbLogger() (*gorm.Config, error) {
	today := time.Now().Format("2006-01-02")
	filePath := fmt.Sprintf("logs/%s-db.log", today)
	logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, xlogger.Error(err)
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
