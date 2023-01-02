package main

import "fmt"

func main() {
	i := 2

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}
	// Two

	v := 64
	switch v {
	case 16:
		fmt.Println(16)
		fallthrough //fallthrough in to next block without checking
	case 32:
		fmt.Println(32)
	case 64:
		fmt.Println(32)
	case 65:
		fmt.Println(32)
		fallthrough
	default:
		fmt.Println("Default")
	}
	// 32
}
