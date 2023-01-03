package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a) // [0 0 0 0 0]

	//a[5] = 200 // Error: invalid argument: index 5 out of bounds [0:5]

	a[4] = 100
	fmt.Println(a)      // [0 0 0 0 100]
	fmt.Println(a[4])   // 100
	fmt.Println(len(a)) // 5

	var b = [3]int{10, 20, 30}
	fmt.Println(b[0]) // 10
}
