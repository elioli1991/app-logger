package app_logger

import (
	"github.com/elioli1991/app-infra/abstract"
)

var DefaultLogger = NewLogger()

func NewLogger() abstract.Logger {
	l := &logger{}
	return l
}

type logger struct {
	abstract.Logger
}

func (l *logger) Info(a ...interface{}) {

}

func (l *logger) Infow(keyvals ...interface{}) {

}

func (l *logger) Infof(format string, a ...interface{}) {

}
