package main

import "github.com/algorand/go-algorand/logging"

func tryMlockall(log logging.Logger) {
	log.Infof("running macOS -- not calling mlockall")
}
