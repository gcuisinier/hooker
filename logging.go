package main

import (
	"os"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("hooker")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{longfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func initLog() {

	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend1Formatter := logging.NewBackendFormatter(backend1, format)
	logging.SetBackend(backend1Formatter)
	log.ExtraCalldepth = 1

}

func debug(logText ...interface{}) {
	debug := os.Getenv("HOOKER_DEBUG")
	if debug != "" {
		log.Debug(logText)
	}
}

func debugf(format string, arguments ...interface{}) {
	debug := os.Getenv("HOOKER_DEBUG")
	if debug != "" {
		log.Debugf(format, arguments...)
	}
}
