package app_logger

import (
	"github.com/elioli1991/app-infra/abstract"
	"sync"
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

func Error(a ...interface{}) {
	global.Error(a)
}

func Errorf(format string, a ...interface{}) {
	global.Errorf(format, a)
}

func Errorw(keyvals ...interface{}) {
	global.Errorw(keyvals)
}

func Fatal(a ...interface{}) {
	global.Fatal(a)
}

func Fatalf(format string, a ...interface{}) {
	global.Fatalf(format, a)
}

func Fatalw(keyvals ...interface{}) {
	global.Fataw(keyvals)
}
