package server

import (
	"golang.org/x/net/context"
	"net/http"
	"os"
	"os/signal"
	"periodic-error/pkg/logger"
	"periodic-error/pkg/routes"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewGinHandler(logg logger.Logger, loggerLogrus *logrus.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	e.Group("/").GET("health-check", routes.HealthCheck)
	strategyRouterGroup := e.Group("")
	routes.StrategyRouteHandler(strategyRouterGroup, logg, loggerLogrus)

	return e
}

func StartServer(s *http.Server, log *logrus.Logger) {
	log.Infof("Starting Server on %s\n", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Warnf("Server shutted down [%s]", err)
	}
}

func ListenShutdownSignal(s *http.Server, log *logrus.Logger) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infof("Shutting down server")

	ctx, cancelFunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelFunc()

	if err := s.Shutdown(ctx); err != nil {
		log.Infof("Server shutdown error [%s]", err)
	} else {
		log.Infof("Server exited")
	}

}
