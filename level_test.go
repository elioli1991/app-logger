package app_logger

import (
	"github.com/elioli1991/app-infra/abstract"
	"testing"
)

func TestParseLevel(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want abstract.Level
	}{
		{
			name: "info",
			args: args{s: "INFO"},
			want: LevelInfo,
		},
		{
			name: "error",
			args: args{s: "Error"},
			want: LevelError,
		},
		{
			name: "fatal",
			args: args{s: "fatal"},
			want: LevelFatal,
		},
		{
			name: "other",
			args: args{s: "sssss"},
			want: LevelInfo,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseLevel(tt.args.s); got != tt.want {
				t.Errorf("ParseLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLevelString(t *testing.T) {
	type args struct {
		l abstract.Level
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "info",
			args: args{l: LevelInfo},
			want: "INFO",
		},
		{
			name: "info",
			args: args{l: LevelError},
			want: "ERROR",
		},
		{
			name: "fatal",
			args: args{l: LevelFatal},
			want: "FATAL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLevelString(tt.args.l); got != tt.want {
				t.Errorf("GetLevelString() = %v, want %v", got, tt.want)
			}
		})
	}
}
