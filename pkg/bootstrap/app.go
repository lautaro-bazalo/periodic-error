package bootstrap

import (
	"appLau/pkg/logger"
	"appLau/pkg/server"
	"appLau/pkg/strategy"
	"net/http"

	"github.com/sirupsen/logrus"
)

type application struct {
	log    *logrus.Logger
	logger logger.Logger
}

type Application interface {
	RunServer()
}

func NewApplication(l *logrus.Logger) Application {

	writerInfo := &strategy.StrategyInfo{}
	newLogger := logger.InitLogger(writerInfo)

	return &application{
		log:    l,
		logger: newLogger,
	}
}

func (app *application) RunServer() {
	s := &http.Server{
		Addr:    "9290",
		Handler: server.NewGinHandler(),
	}
	go server.StartServer(s, app.log)

}
