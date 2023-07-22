package routes

import (
	"appLau/pkg/logger"
	"appLau/pkg/strategy"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type strategyRouter struct {
	l   logger.Logger
	log *logrus.Logger
}

func StrategyRouteHandler(r *gin.RouterGroup, l logger.Logger, logrusLog *logrus.Logger) {
	s := &strategyRouter{
		l:   l,
		log: logrusLog,
	}

	r.Group("/strategy")
	r.POST("/:writer", s.PostRouteStrategy)

}

func (sr *strategyRouter) PostRouteStrategy(ctx *gin.Context) {
	writer := ctx.Param("writer")
	if strategyWriter, err := createStrategyWrite(writer); err != nil {
		ctx.JSON(400, "bad writer match")
	} else {
		sr.l.SetWriterType(strategyWriter)
		sr.l.Write(sr.log)
	}

}

func createStrategyWrite(writer string) (strategy.StrategyWriter, error) {
	switch writer {
	case "info":
		return &strategy.StrategyInfo{}, nil
	case "debug":
		return &strategy.StrategyDebug{}, nil
	case "error":
		return &strategy.StrategyError{}, nil
	default:
		return nil, fmt.Errorf("error when I tried match type of writer")
	}

}
