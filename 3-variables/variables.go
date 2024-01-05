package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a) // initial

	var b, c int = 1, 2
	fmt.Println(b, c) // 1 2

	var d = true
	fmt.Println(d) // true

	var e int
	fmt.Println(e) // 0

	f := "apple"
	fmt.Println(f) // "apple"

	var x string
	fmt.Println(x) // ""

	var sunny = true
	fmt.Printf("type is %T, value is %v\n", sunny, sunny)
}
