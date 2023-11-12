package main

import (
	"fmt"
	"math/big"
)

func Hello() string {
	return "Hello, world!!"
}

func main() {
	fmt.Println(Hello())
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

func IsPrime(x int) (int, bool) {
	if big.NewInt(int64(x)).ProbablyPrime(0) {
		return x, true
	}
	return 0, false
}
