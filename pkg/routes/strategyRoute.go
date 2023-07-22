package routes

import (
	"appLau/pkg/api"
	"appLau/pkg/logger"
	"appLau/pkg/strategy"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	r.Group("/strategy")
	r.POST("/:writer", s.PostRouteStrategy)

}

func (sr *strategyRouter) PostRouteStrategy(ctx *gin.Context) {
	var writerRequest api.WriterRequest

	if err := ctx.ShouldBindJSON(writerRequest); err != nil {
		_ = ctx.AbortWithError(500, err)
		return
	}

	sr.l.SetWriterType(sr.strategyWriter[writerRequest.Writer])
	sr.l.Write(sr.log)

}
