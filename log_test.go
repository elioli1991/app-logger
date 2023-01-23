package logger

import (
	"reflect"
	"testing"

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
