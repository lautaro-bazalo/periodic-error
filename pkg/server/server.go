package server

import (
	"appLau/pkg/routes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewGinHandler() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	e.Group("/").GET("health-check", routes.HealthCheck)

	return e
}

func StartServer(s *http.Server, log *logrus.Logger) {
	log.Infof("Starting Server on %s\n", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Warnf("Server shutted down [%s]", err)
	}
}
