package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if (8%2 == 0) || (7%2 == 0) {
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 {
		fmt.Println("is negative")
	} else if num < 10 {
		fmt.Println("has 1 digit")
	} else {
		fmt.Println("has multiple digits")
	}

	if foo := 1; foo < 2 {
		//the foo variable is only available in this block
		fmt.Println(foo)
	}
	// fmt.Println(foo) // Error: undefined: foo
}
