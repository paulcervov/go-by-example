package main

import (
	"fmt"
	"time"
)

func myFunc(num int) {
	switch num {
	case 16:
		fmt.Println(16)
		fallthrough //fallthrough in to next block without checking
	case 32:
		fmt.Println(32)
	case 64:
		fmt.Println(64)
	case 65:
		fmt.Println(65)
	default:
		fmt.Println("Default")
	}
}

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

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend!")
	case time.Wednesday:
		fmt.Println("It's the middle of the week")
	default:
		fmt.Println("Work hard!")
	}

	myFunc(16) // 16 32
	myFunc(32) // 32
	myFunc(57) // Default
}
