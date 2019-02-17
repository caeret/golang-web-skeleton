package main

import (
	"fmt"
	"os"

	"github.com/caeret/golang-web-skeleton/resource"

	"github.com/caeret/golang-web-skeleton/app"
	"github.com/caeret/golang-web-skeleton/routing"
)

func main() {
	if err := app.LoadConfig(); err != nil {
		panic(fmt.Errorf("fail to load config: %s", err))
	}
	logger := app.NewLogger()
	logger.Info("config loaded.", "config", app.Config)
	if err := app.InitDB(); err != nil {
		logger.Crit("fail to init db.", "error", err)
		os.Exit(1)
	}
	logger.Info("db initialized.")
	if n, err := app.Migrate(resource.SQLBox(), true); err != nil {
		logger.Crit("fail to migrate db.", "error", err)
		os.Exit(1)
	} else {
		logger.Info("db migrated.", "n", n)
	}

	routing.Serve(logger)
}
