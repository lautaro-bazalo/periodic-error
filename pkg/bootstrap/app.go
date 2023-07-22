package bootstrap

import (
	"net/http"
	"periodic-error/pkg/logger"
	"periodic-error/pkg/server"
	"periodic-error/pkg/strategy"

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

	ginHandler := server.NewGinHandler(app.logger, app.log)

	s := &http.Server{
		Addr:    ":9290",
		Handler: ginHandler,
	}
	go server.StartServer(s, app.log)

	server.ListenShutdownSignal(s, app.log)
}
