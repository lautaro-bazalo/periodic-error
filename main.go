package main

import (
	"appLau/pkg/bootstrap"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	app := bootstrap.NewApplication(l)
	app.RunServer()
}
