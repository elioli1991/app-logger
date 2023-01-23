package logger

import (
	"sync"

	"github.com/elioli1991/app-infra/abstract"
)

var global = &loggerContainer{}

type loggerContainer struct {
	abstract.Logger
	lock sync.Mutex
}

func init() {
	global.SetLogger(DefaultLogger)
}

// SetLogger current import init
func (c *loggerContainer) SetLogger(l abstract.Logger) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.Logger = l
}

// SetLogger global function to set logger
func SetLogger(l abstract.Logger) {
	global.SetLogger(l)
}

// GetLogger global function to get logger
func GetLogger() abstract.Logger {
	return global
}

// Debug default init global logger show debug msg
func Debug(a ...interface{}) {
	global.Debug(a...)
}

// Debugf default init global logger show debugf format msg
func Debugf(format string, a ...interface{}) {
	global.Debugf(format, a...)
}

// Debugw default init global logger show debugw msg
func Debugw(keyvals ...interface{}) {
	global.Debugw(keyvals)
}

// Info default init global logger show info msg
func Info(a ...interface{}) {
	global.Info(a...)
}

// Infof default init global logger show infof format msg
func Infof(format string, a ...interface{}) {
	global.Infof(format, a...)
}

// Infow default init global logger show infow msg
func Infow(keyvals ...interface{}) {
	global.Infow(keyvals...)
}

// Warn defaule init global logger show warn msg
func Warn(a ...interface{}) {
	global.Warn(a...)
}

// Warnf default init global logger show warnf msg
func Warnf(format string, a ...interface{}) {
	global.Warnf(format, a...)
}

// Warnw default init global logger show warnw msg
func Warnw(a ...interface{}) {
	global.Warnw(a...)
}

// Error default init global logger show error msg
func Error(a ...interface{}) {
	global.Error(a)
}

// Errorf default init global logger show errorf msg
func Errorf(format string, a ...interface{}) {
	global.Errorf(format, a)
}

// Errorw default init global logger show errorw msg
func Errorw(keyvals ...interface{}) {
	global.Errorw(keyvals)
}

// Fatal default init global logger show fatal msg
func Fatal(a ...interface{}) {
	global.Fatal(a)
}

// Fatalf default init global logger show fatalf msg
func Fatalf(format string, a ...interface{}) {
	global.Fatalf(format, a)
}

// Fatalw default init global logger show fatalw msg
func Fatalw(keyvals ...interface{}) {
	global.Fataw(keyvals)
}
