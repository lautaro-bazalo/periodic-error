package strategy

import (
	"github.com/sirupsen/logrus"
	"time"
)

type StrategyDebug struct{}

func (w *StrategyDebug) WritePeriodic(l *logrus.Logger) {
	for range time.Tick(time.Second * 10) {
		l.Debug("Log Error")
	}
}
