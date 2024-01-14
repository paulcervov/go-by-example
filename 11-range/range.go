package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3}
	for _, num := range nums {
		fmt.Println(num)
	}
	// 1
	// 2
	// 3

	info := map[string]any{"name": "Paul", "age": 39, "married": true}
	for key, value := range info {
		fmt.Print(key, ": ", value, "\n")
	}
	// name: Paul
	// age: 39
	// married: true

	for i, c := range "Go-去" {
		fmt.Println(i, c, string(c))
	}
	// 0 71 G
	// 1 111 o
	// 2 45 -
	// 3 21435 去
}
