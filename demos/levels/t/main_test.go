package main

import (
	"testing"
)

func Test_minFriends(t *testing.T) {
}

func Test_minFriends1(t *testing.T) {
	type args struct {
		comments [][]int
		friends  map[int]struct{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				comments: [][]int{
					{6, 3},
					{1, 3},
					{3, 1},
					{1, 3},
					{2, 4},
					{1, 5},
					{3, 3},
				},
				friends: map[int]struct{}{
					1: {},
					4: {},
					2: {},
				},
			},
			want: 5,
		},
		{
			name: "case2",
			args: args{
				comments: [][]int{
					{1234567, 7654321},
					{1234567, 8888888},
					{8888888, 1234567},
				},
				friends: map[int]struct{}{
					1234567:   {},
					8888888:   {},
					7654321:   {},
					10007:     {},
					998244353: {},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := minFriends(tt.args.comments, tt.args.friends); got != tt.want {
					t.Errorf("minFriends() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
