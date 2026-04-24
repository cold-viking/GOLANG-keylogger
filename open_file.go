package main

import (
	"os"
	"time"
)

func openLogFile() (*os.File, error) {
	logFileName := "log_" + time.Now().Format("2006-01-02_15-04-05") + ".txt"
	file, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return nil, err
	}
	return file, nil
}
