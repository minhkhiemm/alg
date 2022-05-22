package codes

const (
	// ExitCodeKMDInvalidArgs is returned if any cli arguments are invalid
	ExitCodeKMDInvalidArgs = 1

	// ExitCodeKMDLogError is returned if we can't open the log file
	ExitCodeKMDLogError = 2

	// ExitCodeKMDError is the catch-all exit code for most kmd errors
	ExitCodeKMDError = 3

	// ExitCodeKMDAlreadyRunning is returned if an instance of kmd is
	// already running in the given data directory
	ExitCodeKMDAlreadyRunning = 4
)
