package logger

import (
	"context"
	"reflect"
	"testing"

	"github.com/elioli1991/app-logger/output"

	"github.com/elioli1991/app-infra/abstract"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name string
		want abstract.Logger
	}{
		{},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLogger()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogger() = %v, want %v", got, tt.want)
			}
			got.Info(11111)
		})
	}
}

func Test_logger_Info(t *testing.T) {
	type fields struct {
		ctx       context.Context
		Logger    abstract.Logger
		stdLogger abstract.StdLogger
		valuer    []Valuer
		outputer  output.LogOutputer
	}
	type args struct {
		a []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				a: []interface{}{"1", "45444", "22222"},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLogger()
			l.Info(tt.args.a...)
		})
	}
}
