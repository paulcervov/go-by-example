package main

import "fmt"

// recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7)) // 5040

	var fib func(n int) int

	// Recursive closures require explicit declaration
	// of typed variables before definition
	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7)) // 13
}
