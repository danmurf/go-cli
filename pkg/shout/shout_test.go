package shout_test

import (
	"github.com/danmurf/go-cli/pkg/shout"
	"testing"
)

func TestShout_Shout(t *testing.T) {
	type args struct {
		in0   string
		level int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"regular shout 1",
			args{in0: "hello", level: 5},
			"HELLO!!!!!",
		},
		{
			"regular shout 2",
			args{in0: "YES", level: 1},
			"YES!",
		},
		{
			"regular shout 3",
			args{in0: "NO", level: 3},
			"NO!!!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := shout.NewShout()
			if got := l.Shout(tt.args.in0, tt.args.level); got != tt.want {
				t.Errorf("Shout() = %v, want %v", got, tt.want)
			}
		})
	}
}
