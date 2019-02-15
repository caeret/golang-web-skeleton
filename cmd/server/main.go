package main

import (
	"github.com/caeret/golang-web-skeleton/routing"
	"github.com/inconshreveable/log15"
)

func main() {
	routing.Serve(log15.New(), nil)
}
