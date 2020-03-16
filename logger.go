package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Logger struct {
	MyNewLog *log.Logger
}

var (
	fileName  = "error.log"
	filePath  = filepath.Join("source", "logs", fileName)
	inFile, _ = os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModePerm)
	LogThis   = log.New(inFile, fmt.Sprint(time.Now().Format("02Jan 15:04:05 ")), 0)
	MyLogger  = Logger{MyNewLog: LogThis}
)

func (l *Logger) CheckErr(prefix string, err error) {
	if err != nil {
		l.MyNewLog.Printf("%s: %v\n", prefix, err)
	}
}
