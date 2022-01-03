package main

import (
	"github.com/jili/pkg-practice/log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	dir, _ := os.Getwd()
	logFile, _ := os.OpenFile(filepath.Join(dir, "a1234.log"), os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)
	logger := log.NewLogger(log.NewBase(logFile))
	logger.SetLevel(log.DebugLevel)

	for {
		logger.Debug("fwfew")
		logger.Info("fdwfw")
		time.Sleep(2 * time.Second)
	}

}
