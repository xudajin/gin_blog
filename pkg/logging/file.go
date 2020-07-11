package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	prefixpath := getLogFilePath()
	suffixpath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixpath, suffixpath)
}

func openLogFile(filepath string) *os.File {
	_, err := os.Stat(filepath)
	if !os.IsExist(err) {
		mkDir()
	} else if os.IsPermission(err) {
		log.Fatalf("permission: %v", err)

	}

	handle, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail openfile : %v", err)
	}
	return handle
}

func mkDir() {
	dir, _ := os.Getwd()
	fullPath := dir + "/" + getLogFilePath()
	err := os.MkdirAll(fullPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

}
