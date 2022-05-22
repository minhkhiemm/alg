package main

import (
	"flag"
	"github.com/algorand/alg/cmd/kmd/codes"
	"github.com/algorand/alg/daemon/kmd"
	"github.com/algorand/alg/daemon/kmd/server"
	"github.com/algorand/go-algorand/logging"
	"golang.org/x/sys/unix"
	"os"
	"os/signal"
	"path/filepath"
	"time"
)

const (
	kmdLogFileName = "kmd.log"
	kmdLogFilePerm = 0640
)

func main() {
	dataDir := flag.String("d", "", "kmd data directory")
	timeoutSecs := flag.Uint("t", 0, "number of seconds after which to kill kmd if there are no requests. 0 means no timeout.")
	flag.Parse()


	log := logging.NewLogger()
	log.SetLevel(logging.Info)

	if *dataDir == "" {
		log.Errorf("dataDir (-d) is a required argument")
		os.Exit(codes.ExitCodeKMDInvalidArgs)
	}

	var timeout *time.Duration
	if *timeoutSecs != 0 {
		t := time.Duration(*timeoutSecs) * time.Second
		timeout = &t
	}

	kmdLogFilePath := filepath.Join(*dataDir, kmdLogFileName)
	kmdLogFileMode := os.O_CREATE | os.O_WRONLY | os.O_APPEND
	logFile, err := os.OpenFile(kmdLogFilePath, kmdLogFileMode, kmdLogFilePerm)
	if err != nil {
		log.Errorf("failed to open log file: %s",err)
		os.Exit(codes.ExitCodeKMDLogError)
	}
	log.SetOutput(logFile)

	// Prevent swapping with mlockall if supported by the platform
	tryMlockall(log)

	// Create a "kill" channel to allow the server to shut down gracefully
	kill := make(chan os.Signal)

	// Timeouts can also send on the kill channel; because signal.Notify
	// will not block, this shouldn't cause an issue. From docs: "Package
	// signal will not block sending to c"
	signal.Notify(kill, os.Interrupt, unix.SIGTERM, unix.SIGINT)
	signal.Ignore(unix.SIGHUP)

	startConfig := kmd.StartConfig{
		DataDir: *dataDir,
		Kill:    kill,
		Log:     log,
		Timeout: timeout,
	}

	died, sock, err := kmd.Start(startConfig)
	if err == server.ErrAlreadyRunning {
		log.Errorf("couldn't start kmd: %s", err)
		os.Exit(codes.ExitCodeKMDAlreadyRunning)
	}

	if err != nil {
		log.Errorf("couldn't start kmd: %s", err)
		os.Exit(codes.ExitCodeKMDError)
	}

	log.Infof("started kmd on sock: %s", sock)


	<-died
	log.Infof("kmd server died. exiting...")
}
