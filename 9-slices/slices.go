package main

import "fmt"

func main() {
	var s []string

	fmt.Println(s)        // []
	fmt.Println(s == nil) // true
	fmt.Println(len(s))   // 0

	//s[0] = 'test' // error: index out of range [0] with length 0
	s = append(s, "one", "two")

	fmt.Println(s)        // [one two]
	fmt.Println(s == nil) // false
	fmt.Println(len(s))   // 2

	var s2 = make([]string, 3)
	fmt.Println(s2)        // [  ] (3 empty strings)
	fmt.Println(s2 == nil) // false
	fmt.Println(len(s2))   // 3
	s2[0] = "Foo"
	s2[1] = "Bar"
	s2[2] = "Buz"
	fmt.Println(s2)      // [Foo Bar Buz]
	fmt.Println(len(s2)) // 3
	//s2[3] = "Qux" // error: index out of range [3] with length 3
	s2 = append(s2, "Qux")
	fmt.Println(s2)      // [Foo Bar Buz Qux]
	fmt.Println(len(s2)) // 4
}
