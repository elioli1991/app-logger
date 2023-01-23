package logger

import (
	"strings"

	"github.com/elioli1991/app-infra/abstract"
)

type LevelKey string

const (
	LevelKeyDebug LevelKey = "DEBUG"
	LevelKeyInfo  LevelKey = "INFO"
	LevelKeyWarn  LevelKey = "WARN"
	LevelKeyError LevelKey = "ERROR"
	LevelKeyFatal LevelKey = "FATAL"
)

func (k LevelKey) String() string {
	return string(k)
}

const (
	LevelDebug abstract.Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

func GetLevelString(l abstract.Level) string {
	switch l {
	case LevelDebug:
		return LevelKeyDebug.String()
	case LevelInfo:
		return LevelKeyInfo.String()
	case LevelWarn:
		return LevelKeyWarn.String()
	case LevelError:
		return LevelKeyError.String()
	case LevelFatal:
		return LevelKeyFatal.String()
	default:
		return ""
	}
}

func ParseLevel(s string) abstract.Level {
	switch strings.ToUpper(s) {
	case LevelKeyDebug.String():
		return LevelDebug
	case LevelKeyInfo.String():
		return LevelInfo
	case LevelKeyWarn.String():
		return LevelWarn
	case LevelKeyError.String():
		return LevelError
	case LevelKeyFatal.String():
		return LevelFatal
	default:
		return LevelInfo
	}
}
