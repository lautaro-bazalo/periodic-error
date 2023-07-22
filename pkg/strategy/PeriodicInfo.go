package strategy

import (
	"github.com/sirupsen/logrus"
	"time"
)

type StrategyInfo struct{}

func (w *StrategyInfo) WritePeriodic(l *logrus.Logger) {
	for range time.Tick(time.Second * 10) {
		l.Infof("Log Info")
	}
}
