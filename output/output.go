package output

var DefaultOutputer = NewLogOutputer(Default)

type LogOutputer interface {
	Format(keyvals ...interface{}) string
}

func NewLogOutputer(kind Kind) LogOutputer {
	switch kind {
	case Default:
		return newDefaultOutputer()
	}
	return nil
}
