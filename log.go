package app_logger

import (
	"fmt"
	"github.com/elioli1991/app-infra/abstract"
	"github.com/elioli1991/app-logger/output"
	"log"
)

var DefaultLogger = NewLogger()

func NewLogger() abstract.Logger {
	l := &logger{}
	std := NewStdLogger(log.Writer())
	l.stdLogger = std
	l.outputer = output.DefaultOutputer
	return l
}

type logger struct {
	abstract.Logger
	stdLogger abstract.StdLogger
	outputer  output.LogOutputer
}

// Info
func (l *logger) Info(a ...interface{}) {
	_ = l.stdLogger.Log(LevelInfo, l.outputer.Format(a...))
}

// Infow
func (l *logger) Infow(keyvals ...interface{}) {
	_ = l.stdLogger.Log(LevelInfo, l.outputer.FormatKeyvals(keyvals...))
}

// Infof
func (l *logger) Infof(format string, a ...interface{}) {
	_ = l.stdLogger.Log(LevelInfo, l.outputer.Formatf(format, a...))
}

// Error
func (l *logger) Error(a ...interface{}) {
	fmt.Println(a...)
}

// Errorf
func (l *logger) Errorf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

// Errorw
func (l *logger) Errorw(keyvals ...interface{}) {

}

// Fatal
func (l *logger) Fatal(a ...interface{}) {
	fmt.Println(a...)
}

// Fatalf
func (l *logger) Fatalf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

// Fatalw
func (l *logger) Fatalw(a ...interface{}) {

}
