package log

import (
	"fmt"
	"time"
)

type Level int

const (
	INFO Level = iota
	ERROR
)

func Log(log string, level Level) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	if level == INFO {
		fmt.Printf("[+][%s] %s\n", currentTime, log)
	}
	if level == ERROR {
		fmt.Printf("[-][%s] %s\n", currentTime, log)
	}
}
