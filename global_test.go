package app_logger

import (
	"testing"
)

func TestInfo(t *testing.T) {
	type args struct {
		a []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			args: args{
				a: []interface{}{"1111", 222, 4444, 66666},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.a...)
		})
	}
}

func TestInfof(t *testing.T) {
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				format: "%v---%v----%v",
				a:      []interface{}{111, 222, 333},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Infof(tt.args.format, tt.args.a...)
		})
	}
}

func TestInfow(t *testing.T) {
	type args struct {
		keyvals []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "double",
			args: args{
				keyvals: []interface{}{"v1", "111", "v2", "2222", "v4", 5555555},
			},
		},
		{
			name: "single",
			args: args{
				keyvals: []interface{}{"v1", "111", "v2", "2222", "v4", 5555555, 666666},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Infow(tt.args.keyvals...)
		})
	}
}
