package app_logger

import "testing"

func TestParseLevel(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want Level
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
