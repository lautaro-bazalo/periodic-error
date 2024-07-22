package strategy

import (
	"time"
	"github.com/sirupsen/logrus"
)

type StrategyDebug struct{}

func (w *StrategyDebug) WritePeriodic(l *logrus.Logger) {
	for range time.Tick(time.Second * 10) {
		l.Debug("Log Debug")
	}
}
