package slog

import (
	"errors"
	"fmt"
	"github.com/tianxinbaiyun/practice/try/mysql/core/util"
	"os"
	"runtime"
	"strings"
	"time"
)

var log *Logger

//Logger 异步日志
type Logger struct {
	console bool
	warn    bool
	info    bool
	tformat func() string
	file    chan interface{}
}

//NewLog 创建异步日志
func NewLog(level string, console bool, buf int) (*Logger, error) {
	log = &Logger{console: console, tformat: format}
	logName := util.GetExecpath() + "/../server.log"
	File, _ := os.OpenFile(logName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if File != nil {
		FileInfo, err := File.Stat()
		if err != nil {
			return nil, err
		}
		mode := strings.Split(FileInfo.Mode().String(), "-")
		if strings.Contains(mode[1], "w") {
			strChan := make(chan interface{}, buf)
			log.file = strChan
			go func() {
				for {
					fmt.Fprintln(File, <-strChan)
				}
			}()
		} else {
			return nil, errors.New("can't write")
		}
	}

	switch level {
	case "Warn":
		log.warn = true
		return log, nil
	case "Info":
		log.warn = true
		log.info = true
		return log, nil
	}
	err := errors.New("level must be Warn or Info")
	return nil, err
}

// Fatal Fatal
func Fatal(info ...interface{}) {
	if log.console {
		fmt.Println("Fatal", log.tformat(), info)
	}
	if log.file != nil {
		log.file <- fmt.Sprintln("Fatal", log.tformat(), info)
	}
	os.Exit(1)
}

//Error 错误级别
func Error(info ...interface{}) {
	if log.console {
		fmt.Println("Error", log.tformat(), info)
	}
	if log.file != nil {
		log.file <- fmt.Sprintln("Error", log.tformat(), info)
	}
}

// ErrorDB ErrorDB
func ErrorDB(info ...interface{}) {
	if log.console {
		fmt.Println("Error", log.tformat(), info)
	}
	if log.file != nil {
		log.file <- fmt.Sprintln("Error", log.tformat(), info)
	}
}

// Warn Warn
func Warn(info ...interface{}) {
	if log.warn && log.console {
		fmt.Println("Warn", log.tformat(), info)
	}
	if log.file != nil {
		log.file <- fmt.Sprintln("Warn", log.tformat(), info)
	}
}

// Info Info
func Info(info ...interface{}) {
	if log.info && log.console {
		fmt.Println("Info", log.tformat(), info)
	}
	if log.file != nil {
		log.file <- fmt.Sprintln("Info", log.tformat(), info)
	}
}

// Close Close
func Close() {
	for len(log.file) > 0 {
		time.Sleep(1e8)
	}
}
func format() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	return time.Now().Format("2006-01-02 15:04:05") + fmt.Sprintf(" %s:%d ", file, line)
}
