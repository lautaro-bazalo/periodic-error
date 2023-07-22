package strategy

import (
	"github.com/sirupsen/logrus"
)

type StrategyWriter interface {
	WritePeriodic(l *logrus.Logger)
}
