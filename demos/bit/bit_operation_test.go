package main

import (
	"fmt"
	"testing"
)

func TestIsDirectory(t *testing.T) {
	type args struct {
		flags Flags
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"READ", args{READ}, false},
		{"WRITE", args{WRITE}, false},
		{"EXECUTE", args{EXECUTE}, false},
		{"FILE", args{FILE}, false},
		{"DIRECTORY", args{DIRECTORY}, true},
		{"ALL", args{DIRECTORY | READ | WRITE | EXECUTE | FILE}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDirectory(tt.args.flags); got != tt.want {
				t.Errorf("IsDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsExecute(t *testing.T) {
	type args struct {
		flags Flags
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"READ", args{READ}, false},
		{"WRITE", args{WRITE}, false},
		{"EXECUTE", args{EXECUTE}, true},
		{"FILE", args{FILE}, false},
		{"DIRECTORY", args{DIRECTORY}, false},
		{"ALL", args{DIRECTORY | READ | WRITE | EXECUTE | FILE}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsExecute(tt.args.flags); got != tt.want {
				t.Errorf("IsExecute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFile(t *testing.T) {
	type args struct {
		flags Flags
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"READ", args{READ}, false},
		{"WRITE", args{WRITE}, false},
		{"EXECUTE", args{EXECUTE}, false},
		{"FILE", args{FILE}, true},
		{"DIRECTORY", args{DIRECTORY}, false},
		{"ALL", args{DIRECTORY | READ | WRITE | EXECUTE | FILE}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFile(tt.args.flags); got != tt.want {
				t.Errorf("IsFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRead(t *testing.T) {
	type args struct {
		flags Flags
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"READ", args{READ}, true},
		{"WRITE", args{WRITE}, false},
		{"EXECUTE", args{EXECUTE}, false},
		{"FILE", args{FILE}, false},
		{"DIRECTORY", args{DIRECTORY}, false},
		{"ALL", args{DIRECTORY | READ | WRITE | EXECUTE | FILE}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRead(tt.args.flags); got != tt.want {
				t.Errorf("IsRead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsWrite(t *testing.T) {
	type args struct {
		flags Flags
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"READ", args{READ}, false},
		{"WRITE", args{WRITE}, true},
		{"EXECUTE", args{EXECUTE}, false},
		{"FILE", args{FILE}, false},
		{"DIRECTORY", args{DIRECTORY}, false},
		{"ALL", args{DIRECTORY | READ | WRITE | EXECUTE | FILE}, true},
	}
	fmt.Println(DIRECTORY | READ | WRITE | EXECUTE | FILE)
	fmt.Println(DIRECTORY)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsWrite(tt.args.flags); got != tt.want {
				t.Errorf("IsWrite() = %v, want %v", got, tt.want)
			}
		})
	}
}
