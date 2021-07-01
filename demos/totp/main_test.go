package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"hash"
	"testing"
	"time"
)

func Test_flow(t *testing.T) {
	type args struct {
		i    int
		step int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "t1",
			args: args{
				i:    1,
				step: 30,
			},
			want: 0,
		},
		{
			name: "t1",
			args: args{
				i:    2,
				step: 30,
			},
			want: 0,
		},
		{
			name: "t2",
			args: args{
				i:    31,
				step: 30,
			},
			want: 1,
		},
		{
			name: "t3",
			args: args{
				i:    61,
				step: 30,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := flow(tt.args.i, tt.args.step); got != tt.want {
					t.Errorf("flow() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestToken_GenByTime(t1 *testing.T) {
	type fields struct {
		key []byte
		h   func() hash.Hash
		len int
	}

	tn, err := time.Parse(time.RFC3339, "2021-01-02T00:01:00+08:00")
	if err != nil {
		panic(err.Error())
	}

	type args struct {
		now time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "t1",
			fields: fields{
				key: []byte("hello"),
				h: func() hash.Hash {
					return hmac.New(sha256.New, []byte("hello"))
				},
				len: 8,
			},
			args: args{
				now: tn,
			},
			want: "298B0E9F",
		},
	}
	for _, tt := range tests {
		t1.Run(
			tt.name, func(t1 *testing.T) {
				t := &Token{
					h:   tt.fields.h,
					len: tt.fields.len,
				}
				if got := t.GenByTime(tt.args.now); got != tt.want {
					t1.Errorf("GenByTime() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestToken_sum(t1 *testing.T) {
	type fields struct {
		h   func() hash.Hash
		len int
	}
	type args struct {
		f int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "t1",
			fields: fields{
				h: func() hash.Hash {
					return hmac.New(sha256.New, []byte("hello2"))
				},
				len: 8,
			},
			args: args{
				f: 520542,
			},
			want: "B0A3371D05B774D46E668BD5E56C0C181D8E301B83D441422F878A0BE8B186F9",
		},
	}
	for _, tt := range tests {
		t1.Run(
			tt.name, func(t1 *testing.T) {
				t := &Token{
					h:   tt.fields.h,
					len: tt.fields.len,
				}
				got := t.sum(tt.args.f)
				if got != tt.want {
					t1.Errorf("sum() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
