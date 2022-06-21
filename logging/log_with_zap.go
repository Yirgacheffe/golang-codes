package main

import (
	"log"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
}

func main() {
	log.Print("Logging in Go!")

	lg := zap.NewExample().Sugar()
	defer lg.Sync()

	lg.Info("fetch the url from some web site.", "data-dir", "dir-type")
	lg.Infof("%d %s", 10, "this is error.")

	logger.Info("this is from prod logger")
}
