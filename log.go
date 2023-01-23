package logger

import (
	"context"
	"log"
	"os"

	"github.com/elioli1991/app-infra/abstract"
	"github.com/elioli1991/app-logger/output"
)

// DefaultLogger default logger to use
var DefaultLogger = NewLogger()

// DefaultMessageKey default message key.
var DefaultMessageKey = "msg"

func NewLogger() abstract.Logger {
	l := &logger{
		ctx:       context.Background(),
		stdLogger: NewStdLogger(log.Writer()),
		outputer:  output.DefaultOutputer,
	}
	return l
}

// logger core struct
type logger struct {
	ctx context.Context
	abstract.Logger
	stdLogger abstract.StdLogger
	valuer    []Valuer
	outputer  output.LogOutputer
}

func WithValuer(l abstract.Logger, valuer ...Valuer) abstract.Logger {
	c, ok := l.(*logger)
	if !ok {
		return &logger{stdLogger: NewStdLogger(log.Writer()), valuer: valuer, outputer: output.NewLogOutputer(output.Default)}
	}

	return &logger{
		ctx:       c.ctx,
		Logger:    c.Logger,
		stdLogger: c.stdLogger,
		valuer:    append(c.valuer, valuer...),
		outputer:  c.outputer,
	}
}

// Info print info msg
func (l *logger) Info(a ...interface{}) {
	_ = l.stdLogger.Log(LevelInfo, l.outputer.Format(DefaultMessageKey, a...))
}

// Infof printf info msg
func (l *logger) Infof(format string, a ...interface{}) {
	_ = l.stdLogger.Log(LevelInfo, l.outputer.Formatf(format, a...))
}

// Infow print info keyvals msg
func (l *logger) Infow(keyvals ...interface{}) {
	_ = l.stdLogger.Log(LevelInfo, l.outputer.FormatKeyvals(keyvals...))
}

// Error print error msg
func (l *logger) Error(a ...interface{}) {
	_ = l.stdLogger.Log(LevelError, l.outputer.Format(DefaultMessageKey, a...))
}

// Errorf printf error msg
func (l *logger) Errorf(format string, a ...interface{}) {
	_ = l.stdLogger.Log(LevelError, l.outputer.Formatf(format, a...))
}

// Errorw print error keyvals msg
func (l *logger) Errorw(keyvals ...interface{}) {
	_ = l.stdLogger.Log(LevelError, l.outputer.FormatKeyvals(keyvals...))
}

// Fatal print fatal msg
func (l *logger) Fatal(a ...interface{}) {
	_ = l.stdLogger.Log(LevelFatal, l.outputer.Format(DefaultMessageKey, a...))
	os.Exit(1)
}

// Fatalf printf fatal msg
func (l *logger) Fatalf(format string, a ...interface{}) {
	_ = l.stdLogger.Log(LevelFatal, l.outputer.Formatf(format, a...))
	os.Exit(1)
}

// Fatalw print fatal keyvals msg
func (l *logger) Fatalw(a ...interface{}) {
	_ = l.stdLogger.Log(LevelFatal, l.outputer.FormatKeyvals(a...))
	os.Exit(1)
}
