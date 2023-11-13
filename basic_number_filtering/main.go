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

// checks if a float is holding an int or not.
// probably there are better ways of doing this.. but wth.. why not.
func checkInteger(r float64) bool {
	return r == float64(int64(r))
}

// this function implements filtering
// accepts a function that takes an int and returns an (int, bool) tuple
func BoolFilterList(in []int, filterFunc func(x int) (int, bool)) []int {
	out := []int{}
	for _, v := range in {
		if val, ok := filterFunc(v); ok {
			out = append(out, val)
		}
	}
	return out
}

// this is a filter which checks if a int is even or not.
// returns the (int, true) if its even. Else returns (0, false)
func IsEven(x int) (int, bool) {
	res := float64(x) / 2
	// we divide the number by 2.
	// if the quotient is a fraction, then the number is not an integer
	if checkInteger(res) {
		return x, true
	}
	return 0, false
}

// this is a filter function that is designed to work with BoolFilterList
// accepts an int
// returns (int, true) if the int odd
// returns (0, false) otherwise
func IsOdd(x int) (int, bool) {
	if _, ok := IsEven(x); ok {
		return 0, false
	}
	return x, true
}

// filter to check for prime
// uses golang built-in bit.Int.ProbablyPrime
func IsPrime(x int) (int, bool) {
	if big.NewInt(int64(x)).ProbablyPrime(0) {
		return x, true
	}
	return 0, false
}

// Takes 1 int and returns a function which
func genMultipleOf(x int) func(int) (int, bool) {
	return func(in int) (int, bool) {
		r := float64(in) / float64(x)

		if checkInteger(r) {
			return in, true
		}
		return 0, false
	}
}

// greater than operator
func genGreaterThan(x int) func(int) (int, bool) {
	return func(in int) (int, bool) {
		if in > x {
			return in, true
		}
		return 0, false
	}
}

// functions that use the above to compose the solution like asked for in the notion doc

func EvenFilter(in []int) []int {
	return BoolFilterList(in, IsEven)
}

func OddFilter(in []int) []int {
	return BoolFilterList(in, IsOdd)
}

func PrimeFilter(in []int) []int {
	return BoolFilterList(in, IsPrime)
}

// story 4
// we can either filter for odds first and then check for primes OR do the opposite
// this choice depends on how much prime numbers we are expecting in the input list
// AND what is the distribution of even/odd numbers.
func OddPrimeFilter(in []int) []int {
	return BoolFilterList(BoolFilterList(in, IsOdd), IsPrime)
}

// story 5
// filter: even + multiple of 5

func EvenAndMultipleOf5(in []int) []int {
	return BoolFilterList(BoolFilterList(in, genMultipleOf(5)), IsEven)
}

func story6(in []int) []int {
	return BoolFilterList(BoolFilterList(BoolFilterList(in, IsOdd), genMultipleOf(3)), genGreaterThan(10))
}
