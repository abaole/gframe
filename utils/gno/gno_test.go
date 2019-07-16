package gno

import "testing"

func TestGenID(t *testing.T) {
	type args struct {
		mode string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"测算",
			args{
				"k",
			},
			"k777",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenID(tt.args.mode); got != tt.want {
				t.Errorf("GenID() = %v, want %v", got, tt.want)
			}
		})
	}
}
