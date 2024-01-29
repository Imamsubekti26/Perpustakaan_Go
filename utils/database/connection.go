package database

import (
	"fmt"
	"os"

	xlogger "github.com/Imamsubekti26/Perpustakaan_Go/utils/XLogger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, xlogger.Error(err)
	}

	var dbInstance dbInstance
	dbInstance.Connection = db
	return &dbInstance, nil
}
