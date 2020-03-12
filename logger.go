package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Logger struct {
	MyNewLog *log.Logger
	m        *sync.Mutex
}

var (
	fileName  = filepath.Join("source", "logs", "error.log")
	inFile, _ = os.OpenFile(fileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModePerm)
	LogThis   = log.New(inFile, fmt.Sprint(time.Now().Format("02Jan 15:04:05 ")), 0)
	MyLogger  = Logger{MyNewLog: LogThis}
)

func (l *Logger) CheckErr(prefix string, err error) {
	if err != nil {
		l.m.Lock()
		l.MyNewLog.Printf("%s: %v\n", prefix, err)
		l.m.Unlock()
	}
}
