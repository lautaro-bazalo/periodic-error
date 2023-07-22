package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"periodic-error/pkg/api"
	"periodic-error/pkg/logger"
	"periodic-error/pkg/strategy"
)

type strategyRouter struct {
	l              logger.Logger
	log            *logrus.Logger
	strategyWriter map[api.WriterType]strategy.StrategyWriter
}

func StrategyRouteHandler(r *gin.RouterGroup, l logger.Logger, logrusLog *logrus.Logger) {

	strategyWriter := map[api.WriterType]strategy.StrategyWriter{
		api.ErrorW: &strategy.StrategyError{},
		api.InfoW:  &strategy.StrategyInfo{},
		api.DebugW: &strategy.StrategyDebug{},
	}

	s := &strategyRouter{
		l:              l,
		log:            logrusLog,
		strategyWriter: strategyWriter,
	}

	r.Group("/strategy").POST("/:writer", s.PostRouteStrategy)

}

func (sr *strategyRouter) PostRouteStrategy(ctx *gin.Context) {
	var writer api.WriterType
	writer = api.WriterType(ctx.Param("writer"))

	if strategyWriter := sr.strategyWriter[writer]; strategyWriter == nil {
		_ = ctx.AbortWithError(400, fmt.Errorf("unkown strategy writer"))
		return
	} else {
		sr.l.SetWriterType(strategyWriter)
		sr.l.Write(sr.log)
	}

}
