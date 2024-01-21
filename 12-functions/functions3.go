package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {

	total := sum(2, 4)
	fmt.Println(total) // 6

	total = sum(2, 4, 6, 8)
	fmt.Println(total) // 20

	myArgs := []int{1, 2, 3}
	total = sum(myArgs...)
	fmt.Println(total) // 6
}
