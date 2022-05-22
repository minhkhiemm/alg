package server

import "fmt"

// ErrAlreadyRunning is returned if we failed to start kmd because we couldn't
// acquire its file lock. We export this so that the kmd cli can return a
// different exit code for this situation
var ErrAlreadyRunning = fmt.Errorf("failed to lock kmd.lock; is an instance of kmd already running inthis data directory?")
