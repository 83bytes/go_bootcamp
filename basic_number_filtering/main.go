package main

import "fmt"

func Hello() string {
	return "Hello, world!!"
}

func main() {
	fmt.Println(Hello())

	tt := BoolFilterList([]int{0, 1, 2, 3}, IsEven)
	fmt.Println(tt)
}

func IsEven(x int) (int, bool) {
	res := float64(x) / 2
	// we divide the number by 2.
	// if the quotient is a fraction, then the number is not an integer
	if res == float64(int(res)) {
		return x, true
	}
	return 0, false
}

func BoolFilterList(in []int, filterFunc func(x int) (int, bool)) []int {
	out := []int{}
	for _, v := range in {
		if val, ok := filterFunc(v); ok {
			out = append(out, val)
		}
	}
	return out
}

func IsOdd(x int) (int, bool) {
	if _, ok := IsEven(x); ok {
		return 0, false
	}
	return x, true
}
