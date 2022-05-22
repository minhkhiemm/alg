package kmd

import (
	"github.com/algorand/go-algorand/logging"
	"os"
	"time"
)

// StartConfig contains configuration information used to starting up kmd
type StartConfig struct {
	// DataDir is the kmd data directory, used to store config info and
	// some kinds of wallets
	DataDir string
	// Kill takes an os.Signal and gracefully shuts down the kmd process
	Kill chan os.Signal
	// Log logs information about the running kmd process
	Log logging.Logger
	// Timeout is the duration of time after which we will kill the kmd
	// process automatically. If Timeout is nul, we will never time out.
	Timeout *time.Duration
}

// Start loads kmd's configuration information, initializes all of its
// services, and starts the API HTTP server
func Start(startConfig StartConfig) (died chan error, sock string, err error) {
	return
}
