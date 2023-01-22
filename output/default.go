package output

import (
	"fmt"
)

var _ LogOutputer = (*DefaultOutput)(nil)

func NewDefaultOutputer() LogOutputer {
	return &DefaultOutput{}
}

type DefaultOutput struct {
}

func (o *DefaultOutput) Format(a ...interface{}) string {
	var s string
	s = fmt.Sprint(a...)
	return s
}

func (o *DefaultOutput) Formatf(format string, a ...interface{}) string {
	var s string
	s = fmt.Sprintf(format, a...)
	return s
}

// FormatKeyvals format keyvals
// if keyval len % 2 > 0, loss last single element
func (o *DefaultOutput) FormatKeyvals(keyvals ...interface{}) string {
	var s string
	for i := 1; i < len(keyvals)-1; i += 2 {
		s += fmt.Sprintf(" %v=%v", keyvals[i-1], keyvals[i])
	}
	return s
}
