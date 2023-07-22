package strategy

import (
	"github.com/sirupsen/logrus"
	"time"
)

type StrategyError struct{}

func (w *StrategyError) WritePeriodic(l *logrus.Logger) {
	for range time.Tick(time.Second * 10) {
		l.Error("Log Error")
	}
}
