package generic

import (
	"testing"
)

func TestSumIntOrFloat(t *testing.T) {
	type args struct {
		m map[string]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				m: map[string]int{
					"a": 1,
					"b": 2,
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := SumIntOrFloat(tt.args.m); got != tt.want {
					t.Errorf("Sum() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestSumNumber(t *testing.T) {
	type args struct {
		m map[string]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				m: map[string]int{
					"a": 1,
					"b": 2,
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := SumNumber(tt.args.m); got != tt.want {
					t.Errorf("Sum() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestSumNumberExt(t *testing.T) {
	type args struct {
		m map[string]IntType
	}
	tests := []struct {
		name string
		args args
		want IntType
	}{
		{
			name: "t1",
			args: args{
				m: map[string]IntType{
					"a": 1,
					"b": 2,
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := SumNumber(tt.args.m); got != tt.want {
					t.Errorf("Sum() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestServer_GetKey(t *testing.T) {
	server := NewServer[string]()
	server.PrintKey("123")
}
