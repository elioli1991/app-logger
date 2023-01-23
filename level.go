package logger

import (
	"strings"

	"github.com/elioli1991/app-infra/abstract"
)

type LevelKey string

const (
	LevelKeyInfo  LevelKey = "INFO"
	LevelKeyError LevelKey = "ERROR"
	LevelKeyFatal LevelKey = "FATAL"
)

func (k LevelKey) String() string {
	return string(k)
}

const (
	LevelInfo abstract.Level = iota
	LevelError
	LevelFatal
)

func GetLevelString(l abstract.Level) string {
	switch l {
	case LevelInfo:
		return LevelKeyInfo.String()
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
	case LevelKeyInfo.String():
		return LevelInfo
	case LevelKeyError.String():
		return LevelError
	case LevelKeyFatal.String():
		return LevelFatal
	default:
		return LevelInfo
	}
}
