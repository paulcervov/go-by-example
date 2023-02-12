package main

import "fmt"

func main() {
	var a, b, i, sum int

	fmt.Scan(&a, &b)

	for i = a; i <= b; i++ {
		sum += i
	}

	fmt.Println(sum)
}
