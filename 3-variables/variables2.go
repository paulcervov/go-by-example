package main

import "fmt"

func main() {
	var (
		name = "Paul"
		age  = 39
	)

	fmt.Printf("Name: %s, age: %d.\n", name, age) // Name: Paul, age: 39.

	var a int = 10 / 6
	fmt.Println(a) // 1 (!)

	var b float32 = 10 / 6
	fmt.Println(b) // 1

	// at least one operand must be float number
	var c float32 = 10.0 / 6
	fmt.Println(c) // 1.6666666

	// var x int = 10.0 / 6 // cannot use 10.0 / 6 (untyped float constant 1.66667) as int value

	var y int = 10 % 3 // integers only
	fmt.Println(y)     // 1
	// var z float32 = 10.0 % 3 // invalid operation: operator % not defined on 10.0
}
