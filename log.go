package logger

import (
	"context"
	"fmt"
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

func WithStdLogger(l abstract.Logger, std abstract.StdLogger) abstract.Logger {
	c, ok := l.(*logger)
	if !ok {
		return l
	}
	return &logger{
		ctx:       c.ctx,
		Logger:    c.Logger,
		stdLogger: std,
		valuer:    c.valuer,
		outputer:  c.outputer,
	}
}

func WithOutputer(l abstract.Logger, outputer output.LogOutputer) abstract.Logger {
	c, ok := l.(*logger)
	if !ok {
		return l
	}
	return &logger{
		ctx:       c.ctx,
		Logger:    c.Logger,
		stdLogger: c.stdLogger,
		valuer:    c.valuer,
		outputer:  outputer,
	}
}

// Debug print debug msg
func (l *logger) Debug(a ...interface{}) {
	keyvals := make([]interface{}, 0, len(l.valuer)*2+1)
	for _, v := range l.valuer {
		k, value := v(l.ctx)
		keyvals = append(keyvals, k, value)
	}
	keyvals = append(keyvals, DefaultMessageKey, fmt.Sprint(a...))
	_ = l.stdLogger.Log(LevelDebug, l.outputer.Format(keyvals...))
}

// Debugf printf debug msg
func (l *logger) Debugf(format string, a ...interface{}) {
	keyvals := make([]interface{}, 0, len(l.valuer)*2+1)
	for _, v := range l.valuer {
		k, value := v(l.ctx)
		keyvals = append(keyvals, k, value)
	}
	keyvals = append(keyvals, DefaultMessageKey, fmt.Sprintf(format, a...))
	_ = l.stdLogger.Log(LevelDebug, l.outputer.Format(keyvals...))
}

// Debugw print debug keyvals msg
func (l *logger) Debugw(keyvals ...interface{}) {
	newKeyvals := make([]interface{}, 0, len(keyvals)+len(l.valuer)*2)
	for _, v := range l.valuer {
		k, value := v(l.ctx)
		newKeyvals = append(newKeyvals, k, value)
	}
	newKeyvals = append(newKeyvals, keyvals...)
	_ = l.stdLogger.Log(LevelDebug, l.outputer.Format(newKeyvals...))
}

// Info print info msg
func (l *logger) Info(a ...interface{}) {
	keyvals := make([]interface{}, 0, len(l.valuer)*2+1)
	for _, v := range l.valuer {
		k, value := v(l.ctx)
		keyvals = append(keyvals, k, value)
	}
	keyvals = append(keyvals, DefaultMessageKey, fmt.Sprint(a...))
	_ = l.stdLogger.Log(LevelInfo, l.outputer.Format(keyvals...))
}

// Infof printf info msg
func (l *logger) Infof(format string, a ...interface{}) {
	keyvals := make([]interface{}, 0, len(l.valuer)*2+1)
	for _, v := range l.valuer {
		k, value := v(l.ctx)
		keyvals = append(keyvals, k, value)
	}
	keyvals = append(keyvals, DefaultMessageKey, fmt.Sprintf(format, a...))
	_ = l.stdLogger.Log(LevelInfo, l.outputer.Format(keyvals...))
}

// Infow print info keyvals msg
func (l *logger) Infow(keyvals ...interface{}) {
	newKeyvals := make([]interface{}, 0, len(keyvals)+len(l.valuer)*2)
	for _, v := range l.valuer {
		k, value := v(l.ctx)
		newKeyvals = append(newKeyvals, k, value)
	}
	newKeyvals = append(newKeyvals, keyvals...)
	_ = l.stdLogger.Log(LevelInfo, l.outputer.Format(newKeyvals...))
}

// Error print error msg
func (l *logger) Error(a ...interface{}) {
	keyvals := make([]interface{}, 0, len(l.valuer)*2+1)
	for _, v := range l.valuer {
		k, value := v(l.ctx)
		keyvals = append(keyvals, k, value)
	}
	keyvals = append(keyvals, DefaultMessageKey, fmt.Sprint(a...))
	_ = l.stdLogger.Log(LevelError, l.outputer.Format(keyvals))
}

// Errorf printf error msg
func (l *logger) Errorf(format string, a ...interface{}) {
	keyvals := make([]interface{}, 0, len(l.valuer)*2+1)
	for _, v := range l.valuer {
		k, value := v(l.ctx)
		keyvals = append(keyvals, k, value)
	}
	keyvals = append(keyvals, DefaultMessageKey, fmt.Sprintf(format, a...))
	_ = l.stdLogger.Log(LevelError, l.outputer.Format(keyvals))
}

// Errorw print error keyvals msg
func (l *logger) Errorw(keyvals ...interface{}) {
	newKeyvals := make([]interface{}, 0, len(keyvals)+len(l.valuer)*2)
	for _, v := range l.valuer {
		k, value := v(l.ctx)
		newKeyvals = append(newKeyvals, k, value)
	}
	newKeyvals = append(newKeyvals, keyvals...)
	_ = l.stdLogger.Log(LevelError, l.outputer.Format(newKeyvals...))
}

// Fatal print fatal msg
func (l *logger) Fatal(a ...interface{}) {
	keyvals := make([]interface{}, 0, len(l.valuer)*2+1)
	for _, v := range l.valuer {
		k, value := v(l.ctx)
		keyvals = append(keyvals, k, value)
	}
	keyvals = append(keyvals, DefaultMessageKey, fmt.Sprint(a...))
	_ = l.stdLogger.Log(LevelFatal, l.outputer.Format(keyvals...))
	os.Exit(1)
}

// Fatalf printf fatal msg
func (l *logger) Fatalf(format string, a ...interface{}) {
	keyvals := make([]interface{}, 0, len(l.valuer)*2+1)
	for _, v := range l.valuer {
		k, value := v(l.ctx)
		keyvals = append(keyvals, k, value)
	}
	keyvals = append(keyvals, DefaultMessageKey, fmt.Sprintf(format, a...))
	_ = l.stdLogger.Log(LevelFatal, l.outputer.Format(keyvals...))
	os.Exit(1)
}

// Fatalw print fatal keyvals msg
func (l *logger) Fatalw(keyvals ...interface{}) {
	newKeyvals := make([]interface{}, 0, len(keyvals)+len(l.valuer)*2)
	for _, v := range l.valuer {
		k, value := v(l.ctx)
		newKeyvals = append(newKeyvals, k, value)
	}
	newKeyvals = append(newKeyvals, keyvals...)
	_ = l.stdLogger.Log(LevelFatal, l.outputer.Format(newKeyvals...))
	os.Exit(1)
}
