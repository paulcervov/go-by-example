package main

import (
	"fmt"
	"maps"
)

func main() {
	myMap := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(myMap) //map[bar:2 foo:1]

	_, isPresent := myMap["foo"]
	fmt.Println(isPresent) //true

	userData := map[string]any{"name": "Paul", "age": 39} // map[age:39 name:Paul] 🤔
	fmt.Println(userData)

	// var m map[string]int
	// m["age"] = 39 // panic: assignment to entry in nil map

	myCopy := map[string]int{}
	// maps.Copy(myCopy, userData) // M2 (type map[string]any) does not satisfy ~map[K]V
	maps.Copy(myCopy, myMap)
	fmt.Println(myCopy) // map[bar:2 foo:1]

	isEqual := maps.Equal(myMap, myCopy)
	fmt.Println(isEqual) // true
}
