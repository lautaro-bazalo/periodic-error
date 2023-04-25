package logger

import (
	"github.com/sirupsen/logrus"
	"time"
)

type logger struct {
	log *logrus.Logger
}

type Logger interface {
	WritePeriodicError()
}

func NewLogger(log *logrus.Logger) Logger {
	return &logger{
		log: log,
	}
}

func (l *logger) WritePeriodicError() {
	for range time.Tick(time.Second * 10) {
		l.log.Error("Log Error")
	}
}
