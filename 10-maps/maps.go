package main

import "fmt"

func main() {
	myMap := make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2

	fmt.Println(myMap)      // map[one:1 two:2]
	fmt.Println(len(myMap)) // 2

	v1 := myMap["one"]
	fmt.Println(v1) // 1

	v3 := myMap["three"]
	fmt.Println(v3) // 0 (! int zero value)

	delete(myMap, "two")
	fmt.Println(myMap)      // map[one:1]
	fmt.Println(len(myMap)) // 1

	clear(myMap)
	fmt.Println(myMap)      // map[]
	fmt.Println(len(myMap)) // 0
}
