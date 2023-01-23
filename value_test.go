package logger

import (
	"reflect"
	"testing"
)

func TestCaller(t *testing.T) {
	type args struct {
		depth int
	}
	tests := []struct {
		name string
		args args
		want Valuer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Caller(tt.args.depth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Caller() = %v, want %v", got, tt.want)
			}
		})
	}
}
