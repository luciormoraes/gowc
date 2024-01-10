package main

import (
	"testing"
)

func Test_countLines(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{[]byte("Hello\nWorld")}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countLines(tt.args.content); got != tt.want {
				t.Errorf("countLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countWords(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{[]byte("Hello World")}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countWords(tt.args.content); got != tt.want {
				t.Errorf("countWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countMultiBytes(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{[]byte("Hello World")}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMultiBytes(tt.args.content); got != tt.want {
				t.Errorf("countMultiBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
