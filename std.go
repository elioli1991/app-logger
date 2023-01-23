package logger

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/elioli1991/app-infra/abstract"
)

var _ abstract.StdLogger = (*stdLogger)(nil)

func NewStdLogger(w io.Writer) abstract.StdLogger {
	return &stdLogger{
		log: log.New(w, "", 0),
		pool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}

type stdLogger struct {
	log  *log.Logger
	pool *sync.Pool
}

// Log print log
func (l *stdLogger) Log(level abstract.Level, vals string) error {
	buf := l.pool.Get().(*bytes.Buffer)
	_ = l.log.Output(4, buf.String()) //nolint:gomnd
	buf.WriteString(fmt.Sprintf("[%v]", GetLevelString(level)))
	if len(vals) > 0 {
		buf.WriteString(vals)
	}
	_ = l.log.Output(4, buf.String()) //nolint:gomnd
	buf.Reset()
	l.pool.Put(buf)
	return nil
}
