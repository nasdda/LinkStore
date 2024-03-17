package logger

import (
	"log"
	"os"
)

type Logger interface {
	Info(message string)
	Error(message string)
	Errorf(format string, args ...interface{})
	Printf(format string, args ...interface{})
}

type logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func NewLogger(prefix string) Logger {
	return &logger{
		infoLogger:  log.New(os.Stdout, "INFO: "+prefix+" ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: "+prefix+" ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *logger) Info(message string) {
	l.infoLogger.Println(message)
}

func (l *logger) Error(message string) {
	l.errorLogger.Println(message)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.errorLogger.Printf(format, args...)
}

func (l *logger) Printf(format string, args ...interface{}) {
	l.infoLogger.Printf(format, args...)
}
