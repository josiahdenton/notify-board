package logging

import (
	"log"
	"os"
)

type LogLevel int

const (
    INFO LogLevel = iota
    ERROR
)

var (
    infoLogger = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
    errLogger =  log.New(os.Stdout, "ERROR ", log.Ldate|log.Ltime)

)

func New(level LogLevel) *log.Logger {
    if (level == INFO) {
        return infoLogger
    } else if (level == ERROR) {
        return errLogger
    }
    panic("Yeah, don't to that with my logger...")
}
