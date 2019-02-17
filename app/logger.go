package app

import "github.com/inconshreveable/log15"

type Logger log15.Logger

func NewLogger() Logger {
	logger := log15.New()

	var lvl log15.Lvl
	if Config.Debug {
		lvl = log15.LvlDebug
	} else {
		lvl = log15.LvlInfo
	}
	logger.SetHandler(log15.LvlFilterHandler(lvl, logger.GetHandler()))
	return logger
}
