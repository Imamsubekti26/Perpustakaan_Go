package xlogger

import (
	"fmt"
	"log"
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

func Infof(args ...interface{}) *Log {
	return &Log{
		isFatal: false,
		message: fmt.Sprintf(args[0].(string), args[1:]...),
	}
}

func Errorf(args ...interface{}) *Log {
	return &Log{
		isFatal: true,
		message: fmt.Sprintf(args[0].(string), args[1:]...),
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
	WriteToFile(filePath, xlog.message)

	if xlog.isFatal {
		filePath = fmt.Sprintf("logs/%s-error.log", today)
		WriteToFile(filePath, xlog.message)
		panic(xlog.message)
	}
}

func WriteAndShow(err error) {
	Write(err)
	Show(err)
}
