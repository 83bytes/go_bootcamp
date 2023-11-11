package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world!!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestIsEven(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"Two is even!", args{2}, true},
		{"Negetive Two is even!", args{-2}, true},
		{"Zero is even!", args{0}, true},
		{"One is not even!", args{1}, false},
		{"Negetive One is not even!", args{1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.args.x); got != tt.want {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
