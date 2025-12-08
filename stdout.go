package logging

import "log"

func DefaultLogger() Logger {
	return stdlogger{}
}

type stdlogger struct{}

func (l stdlogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (l stdlogger) Println(v ...interface{}) {
	log.Println(v...)
}
