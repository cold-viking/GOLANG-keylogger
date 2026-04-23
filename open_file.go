package main

import (
	"os"
	"time"
)

func openLogFile() (*os.File, error) {
	nameLogFile := "log_" + time.Now().Format("2006-01-02_15-04-05") + ".txt"
	file, err := os.OpenFile(nameLogFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return file, nil
}
