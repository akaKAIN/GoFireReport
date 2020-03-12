package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (fileName = filepath.Join("source", "logs", "error.log")
	inFile, _ = os.OpenFile(fileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModePerm)
	LogThis   = log.New(inFile, fmt.Sprint(time.Now().Format("02Jan 15:04:05 ")), 0)
)

func CheckErr(prefix string, err error) {
	var m sync.Mutex
	if err != nil {
		m.Lock()
		LogThis.Printf("%s: %v\n", prefix, err)
		m.Unlock()
	}
}
