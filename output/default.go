package output

import (
	"fmt"
)

var _ LogOutputer = (*DefaultOutput)(nil)

func newDefaultOutputer() LogOutputer {
	return &DefaultOutput{}
}

// DefaultOutput default output format
type DefaultOutput struct{}

// Format
func (o *DefaultOutput) Format(keyvals ...interface{}) string {
	var s string
	if len(keyvals)%2 > 0 {
		keyvals = keyvals[:len(keyvals)-1]
	}
	for i := 0; i < len(keyvals)-1; i += 2 {
		s += fmt.Sprintf(" %v=%v", keyvals[i], keyvals[i+1])
	}
	return s
}
