package app_logger

import "strings"

type LevelKey string

const (
	LevelKeyInfo  LevelKey = "INFO"
	LevelKeyError LevelKey = "ERROR"
	LevelKeyFatal LevelKey = "FATAL"
)

func (k LevelKey) String() string {
	return string(k)
}

type Level int8

const (
	LevelInfo Level = iota
	LevelError
	LevelFatal
)

func (l Level) GetString() string {
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

func ParseLevel(s string) Level {
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
