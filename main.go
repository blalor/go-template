package main

import (
	"fmt"
	"os"
	"syscall"

	flags "github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

var version string = "undef"

type Options struct {
	Debug   bool   `env:"DEBUG"    long:"debug"    description:"enable debug"`
	LogFile string `env:"LOG_FILE" long:"log-file" description:"path to JSON log file"`
}

func main() {
	var opts Options

	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	if opts.Debug {
		log.SetLevel(log.DebugLevel)
	}

	if opts.LogFile != "" {
		logFp, err := os.OpenFile(opts.LogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
		checkError(fmt.Sprintf("error opening %s", opts.LogFile), err)

		defer logFp.Close()

		// ensure panic output goes to log file
		syscall.Dup2(int(logFp.Fd()), 1)
		syscall.Dup2(int(logFp.Fd()), 2)

		// log as JSON
		log.SetFormatter(&log.JSONFormatter{})

		// send output to file
		log.SetOutput(logFp)
	}

	log.Debug("hi there! (tickertape tickertape)")
	log.Infof("version: %s", version)
}
