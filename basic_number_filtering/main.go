package main

import "fmt"

func Hello() string {
	return "Hello, world!!"
}

func main() {
	fmt.Println(Hello())
}

func IsEven(x int) bool {
	res := float64(x) / 2
	// we divide the number by 2.
	// if the quotient is a fraction, then the number is not an integer
	return res == float64(int(res))
}
