package logger

import (
	"os"
	"reflect"

	"github.com/charmbracelet/log"
)

type Logger struct {
	opts   log.Options
	prefix string
	file   string
}

// Instantiate a new logger
//
// //default file ".log"j
// l, f := logger.New(logger.Logger{})
//
// defer f.Close()
//
// l.Info("test")
func New(l Logger) (*log.Logger, *os.File) {
	var defaultOpts log.Options
	file := ".log"

	if l.file != "" {
		file = l.file
	}

	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o600)
	if err != nil {
		log.Fatal("Opening error \n", err)
	}

	if reflect.ValueOf(l.opts).IsZero() {
		defaultOpts = log.Options{
			ReportCaller:    true,
			ReportTimestamp: true,
		}
	} else {
		defaultOpts = l.opts
	}

	defaultOpts.Prefix = l.prefix

	return log.NewWithOptions(f, defaultOpts), f
}

// Log value to "debug.log"
//
// logger.Log("value")
func Log(value string) {
	l, f := New(Logger{file: "debug.log"})
	defer f.Close()

	l.Debug(value)
}
