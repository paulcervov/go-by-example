package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []string{"a", "b", "c"}
	fmt.Println(s)      // [a b c]
	fmt.Println(len(s)) // 3

	s2 := make([]string, len(s))
	copy(s2, s)

	fmt.Println(s2)      // [a b c]
	fmt.Println(len(s2)) // 3

	if slices.Equal(s, s2) {
		fmt.Println("s == s2")
	} //s == s2

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD) // [[0] [1 2] [2 3 4]]
	fmt.Printf("%T\n", twoD)  //[][]int
}
