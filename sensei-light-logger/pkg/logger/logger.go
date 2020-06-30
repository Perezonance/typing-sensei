package logger

import (
	"fmt"
	"time"
)

const (
	timeFormatLayout = "01/02/2006 15:04:05"

	colorReset = "\033[0m"
	colorRed = "\033[31m"
	colorCyan = "\033[32m"
	colorYellow = "\033[33m"
)

func InfoLog(msg string) {
	currTime := time.Now().Format(timeFormatLayout)
	fmt.Printf(colorCyan + "%v::INFO LOG  >> " + colorReset + msg + "\n", currTime)
}

func ErrorLog(msg string, err error) {
	currTime := time.Now().Format(timeFormatLayout)
	fmt.Printf(colorRed+ "%v::ERROR LOG >> " + colorReset + msg + "\n", currTime)
}

func DebugLog(msg string) {
	currTime := time.Now().Format(timeFormatLayout)
	fmt.Printf(colorYellow + "%v::DEBUG LOG >> " + colorReset + msg + "\n", currTime)
}
