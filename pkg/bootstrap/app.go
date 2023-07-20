package bootstrap

import (
	"appLau/pkg/logger"
	"appLau/pkg/server"
	"net/http"

	"github.com/sirupsen/logrus"
)

type application struct {
	log    *logrus.Logger
	logger logger.Logger
}

type Application interface {
	RunServer()
	RunLogger()
}

func NewApplication(l *logrus.Logger) Application {

	newLogger := logger.NewLogger(l)
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

func (app *application) RunLogger() {
	app.logger.WritePeriodicError()
}
