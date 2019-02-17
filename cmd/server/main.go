package main

import (
	"fmt"

	"github.com/caeret/golang-web-skeleton/app"
	"github.com/caeret/golang-web-skeleton/routing"
)

func main() {
	if err := app.LoadConfig(); err != nil {
		panic(fmt.Errorf("fail to load config: %s", err))
	}
	logger := app.NewLogger()
	if err := app.InitDB(); err != nil {
		panic(fmt.Errorf("fail to init db: %s", err))
	}

	routing.Serve(logger)
}
