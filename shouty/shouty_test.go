package shouty_test

import (
	"github.com/danmurf/go-cli/shouty"
	"testing"
)

func Test_loudMouth_Shout(t *testing.T) {
	type args struct {
		in0 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"regular shout 1",
			args{in0:"hello"},
			"HELLO!!!!!",
		},
		{
			"regular shout 2",
			args{in0:"YES"},
			"YES!!!!!",
		},
		{
			"regular shout 3",
			args{in0:"NO"},
			"NO!!!!!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := shouty.NewLoudMouth()
			if got := l.Shout(tt.args.in0); got != tt.want {
				t.Errorf("Shout() = %v, want %v", got, tt.want)
			}
		})
	}
}