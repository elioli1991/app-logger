package output

var DefaultOutputer = NewLogOutputer(Default)

type LogOutputer interface {
	Format(a ...interface{}) string
	Formatf(format string, a ...interface{}) string
	FormatKeyvals(keyvals ...interface{}) string
}

func NewLogOutputer(kind Kind) LogOutputer {
	switch kind {
	case Default:
		return NewDefaultOutputer()
	}
	return nil
}
