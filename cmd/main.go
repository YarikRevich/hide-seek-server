package main

import (
	"flag"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"

	"github.com/YarikRevich/go-demonizer/pkg/demonizer"

	"github.com/YarikRevich/hide-seek-server/internal/server"
	"github.com/YarikRevich/hide-seek-server/tools/params"
	"github.com/YarikRevich/hide-seek-server/tools/printer"
	"github.com/sirupsen/logrus"
)

func init() {
	rand.Seed(time.Now().Unix())

	flag.Parse()

	if params.IsDemon() {
		demonizer.DemonizeThisProcess()
	}

	logrus.SetFormatter(logrus.StandardLogger().Formatter)

	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.InfoLevel)

	printer.PrintWelcomeMessage()
}

func main() {
	if params.IsProfileCPU() {
		runtime.SetBlockProfileRate(1)
		go func() {
			logrus.Fatalln(http.ListenAndServe(":9909", nil))
		}()
	}

	server.Run()
}
