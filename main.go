package main

import (
	"os"
	"periodic-error/pkg/bootstrap"

	"github.com/sirupsen/logrus"
)

func main() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	app := bootstrap.NewApplication(l)
	app.RunServer()
}