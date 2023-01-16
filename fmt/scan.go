package main

import "fmt"

func main() {
	var number int

	fmt.Println("Enter the number and press ↵:")

	fmt.Scan(&number)

	fmt.Printf("Number is: %d\n", number)
}
