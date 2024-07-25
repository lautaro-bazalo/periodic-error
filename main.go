package main

import (
	"os"
	"periodic-error/internal/bootstrap"

	"github.com/sirupsen/logrus"
)

func main() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	app := bootstrap.NewApplication(l)
	app.RunServer()
}
