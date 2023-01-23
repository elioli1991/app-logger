package logger

import (
	"bytes"
	"log"
	"sync"
	"testing"

	"github.com/elioli1991/app-infra/abstract"
)

func Test_stdLogger_Log(t *testing.T) {
	type fields struct {
		log  *log.Logger
		pool *sync.Pool
	}
	type args struct {
		level abstract.Level
		vals  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "print info",
			fields: fields{
				log: log.New(log.Writer(), "", 1),
				pool: &sync.Pool{
					New: func() interface{} {
						return new(bytes.Buffer)
					},
				},
			},
			args: args{
				level: LevelInfo,
				vals:  "saggnodfowdnskngopengnmsnront",
			},
			wantErr: false,
		},
		{
			name: "print error",
			fields: fields{
				log: log.New(log.Writer(), "", 1),
				pool: &sync.Pool{
					New: func() interface{} {
						return new(bytes.Buffer)
					},
				},
			},
			args: args{
				level: LevelError,
				vals:  " 5=4 222=5 2123=66",
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &stdLogger{
				log:  tt.fields.log,
				pool: tt.fields.pool,
			}
			if err := l.Log(tt.args.level, tt.args.vals); (err != nil) != tt.wantErr {
				t.Errorf("Log() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
