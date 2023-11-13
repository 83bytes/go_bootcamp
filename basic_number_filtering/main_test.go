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

func TestEvenFilter(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"Filter even nos 1", args{[]int{0, 1, 2, 3, 4}}, []int{0, 2, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvenFilter(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EvenFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOddFilter(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"Filter odd", args{[]int{0, 1, 2, 3, 4}}, []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OddFilter(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OddFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrimeFilter(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"Filter Prime", args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}, []int{2, 3, 5, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrimeFilter(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrimeFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOddPrimeFilter(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"Filter Prime", args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, []int{3, 5, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OddPrimeFilter(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OddPrimeFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultipleOf(t *testing.T) {
	type args struct {
		in int
		x  int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		// TODO: Add test cases.
		{"basic yes", args{20, 5}, 20, true},
		{"basic no", args{5, 4}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MultipleOf(tt.args.in, tt.args.x)
			if got != tt.want {
				t.Errorf("MultipleOf() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MultipleOf() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEvenAndMultipleOf5(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"basic yes", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}}, []int{10, 20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvenAndMultipleOf5(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EvenAndMultipleOf5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_story6(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"basic yes", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}}, []int{15}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := story6(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("story6() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_greaterthan3(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		// TODO: Add test cases.
		{"basic negetive", args{2}, 0, false},
		{"basic positive", args{5}, 5, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := greaterthan3(tt.args.x)
			if got != tt.want {
				t.Errorf("greaterthan3() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("greaterthan3() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
