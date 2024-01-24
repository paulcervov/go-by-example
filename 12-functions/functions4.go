package main

import "fmt"

// closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// just anonymous function
	square := func(val int) int {
		return val * val
	}
	fmt.Println(square(4)) // 16

	// using closure
	nextInt := intSeq()
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	nextInt2 := intSeq()
	fmt.Println(nextInt2()) // 1
}
