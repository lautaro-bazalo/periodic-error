package logger

import (
	"periodic-error/pkg/strategy"

	"github.com/sirupsen/logrus"
)

type logger struct {
	writerType strategy.StrategyWriter
}

type Logger interface {
	Write(log *logrus.Logger)
	SetWriterType(writerType strategy.StrategyWriter)
}

func InitLogger(writerType strategy.StrategyWriter) Logger {
	return &logger{
		writerType: writerType,
	}
}

func (l *logger) Write(log *logrus.Logger) {
	go l.writerType.WritePeriodic(log)
}
func (l *logger) SetWriterType(writerType strategy.StrategyWriter) {
	l.writerType = writerType
}
