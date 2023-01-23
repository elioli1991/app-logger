package logger

import (
	"context"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	DefaultCaller    = Caller(4)
	DefaultTimestamp = Timestamp(time.RFC3339)
)

type Valuer func(ctx context.Context) (string, interface{})

// Caller returns a Valuer that returns a pkg/file:line description of the caller.
func Caller(depth int) Valuer {
	return func(context.Context) (string, interface{}) {
		_, file, line, _ := runtime.Caller(depth)
		idx := strings.LastIndexByte(file, '/')
		if idx == -1 {
			return "caller", file[idx+1:] + ":" + strconv.Itoa(line)
		}
		idx = strings.LastIndexByte(file[:idx], '/')
		return "caller", file[idx+1:] + ":" + strconv.Itoa(line)
	}
}

// Timestamp get time
func Timestamp(layout string) Valuer {
	return func(ctx context.Context) (string, interface{}) {
		return "ts", time.Now().Format(layout)
	}
}
