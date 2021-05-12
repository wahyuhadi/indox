package services

import (
	"flag"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func SetLogger(mode, message string) {
	flag.Parse()
	log := logrus.New()
	if !*debug {
		// -- disable logger where logger is false
		log.Out = ioutil.Discard
	}
	switch mode {

	// - error log
	case "error":
		log.Errorf(message)

	// - warning log
	case "warning":
		log.Warningf(message)

	// - info log
	case "info":
		log.Infof(message)
	}
}
