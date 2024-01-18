package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func main() {
	res := plus(4, 3)
	fmt.Println(res) // 7
}
