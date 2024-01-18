package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	x, y := vals()
	fmt.Println(x, y) // 3 7

	z, _ := vals()
	fmt.Println(z) // 3
}
