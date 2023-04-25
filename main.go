package app_lau

import (
	"app-lau/pkg/bootstrap"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	app := bootstrap.NewApplication(l)
	app.RunServer()
	app.RunLogger()
}
