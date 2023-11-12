package main

import (
	"reflect"
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
			got, ok := IsEven(tt.args.x)
			if (got != tt.args.x) && (ok != tt.want) {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}

func doNothing(x int) (int, bool) {
	return x, true
}

func TestBoolFilterList(t *testing.T) {
	type args struct {
		in         []int
		filterFunc func(x int) (int, bool)
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"Filter nothing", args{[]int{0, 1, 2, 3}, doNothing}, []int{0, 1, 2, 3}},
		{"Filter Even", args{[]int{0, 1, 2, 3}, IsEven}, []int{0, 2}},
		{"Empty List", args{[]int{}, IsEven}, []int{}},
		{"Filter Odd", args{[]int{0, 1, 2, 3}, IsOdd}, []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolFilterList(tt.args.in, tt.args.filterFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoolFilterList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsOdd(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"Three is odd!", args{3}, true},
		{"Five is odd!", args{5}, true},
		{"Six is odd!", args{6}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := IsOdd(tt.args.x)
			if (got != tt.args.x) && (ok != tt.want) {
				t.Errorf("IsOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
	}{
		// TODO: Add test cases.
		{"1 is not prime!!", args{1}, false, 0},
		{"8 is NOT prime!!", args{1}, false, 0},
		{"7 is prime!!", args{7}, true, 7},
		{"19 is prime!!", args{19}, true, 19},
		{"18 is NOT prime!!", args{18}, false, 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got := IsPrime(tt.args.x)
			if (got != tt.want) && (got1 != tt.want1) {
				t.Errorf("IsPrime() got = %v, want %v", got, tt.want)
			}
		})
	}
}
