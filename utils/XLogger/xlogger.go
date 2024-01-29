package xlogger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Log struct {
	isFatal bool
	message string
}

// Error implements error.
func (*Log) Error() string {
	panic("unimplemented")
}

func Info(args ...interface{}) *Log {
	return &Log{
		isFatal: false,
		message: fmt.Sprint(args...),
	}
}

func Error(args ...interface{}) *Log {
	return &Log{
		isFatal: true,
		message: fmt.Sprint(args...),
	}
}

func Show(err error) {
	xlog, ok := err.(*Log)
	if !ok {
		xlog = Error(err.Error())
	}

	if xlog.isFatal {
		log.Fatal(xlog.message)
		panic(xlog.message)
	}
	log.Print(xlog.message)
}

func Write(err error) {
	xlog, ok := err.(*Log)
	if !ok {
		xlog = Error(err.Error())
	}

	today := time.Now().Format("2006-01-02")
	filePath := fmt.Sprintf("logs/%s.log", today)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileLogger := log.New(file, "", log.Ldate|log.Ltime)
	fileLogger.Printf(xlog.message)

	if xlog.isFatal {
		panic(xlog.message)
	}
}

func WriteAndShow(err error) {
	Write(err)
	Show(err)
}
