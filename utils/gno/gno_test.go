package gno

import "testing"

func TestGenID(t *testing.T) {
	type args struct {
		mode int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"测算",
			args{
				22,
			},
			"k777",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenID(tt.args.mode)
			t.Log(got)
			if got != tt.want {
				t.Errorf("GenID() = %v, want %v", got, tt.want)
			}
		})
	}
}
