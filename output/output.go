package output

var DefaultOutputer = NewLogOutputer(Default)

type LogOutputer interface {
	Format(key string, a ...interface{}) string
	FormatKeyvals(keyvals ...interface{}) string
}

func NewLogOutputer(kind Kind) LogOutputer {
	switch kind {
	case Default:
		return newDefaultOutputer()
	}
	return nil
}
