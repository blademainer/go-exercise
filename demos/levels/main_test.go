package main

import (
	"testing"
)

func Test_minIcons(t *testing.T) {
	type args struct {
		m int
		s int
		c int
		l int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				m: 4,
				s: 3,
				c: 2,
				l: 84,
			},
			want: 4,
		},
		{
			name: "case2",
			args: args{
				m: 12,
				s: 34,
				c: 56,
				l: 1000000000,
			},
			want: 43812,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := minIcons(tt.args.m, tt.args.s, tt.args.c, tt.args.l); got != tt.want {
					t.Errorf("minIcons() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
