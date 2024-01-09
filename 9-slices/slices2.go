package main

import "fmt"

func main() {
	s := []string{"a", "b", "c"}
	fmt.Println(s)      // [a b c]
	fmt.Println(len(s)) // 3

	s2 := make([]string, len(s))
	fmt.Println(s2)      // [   ]
	fmt.Println(len(s2)) // 3

	copy(s2, s)
	fmt.Println(s2)      // [a b c]
	fmt.Println(len(s2)) // 3

	s3 := s2[0:1]        // slice
	fmt.Println(s3)      // [a]
	fmt.Println(len(s3)) // 1

	s4 := s2[2:3]
	fmt.Println(s4)      // [c]
	fmt.Println(len(s4)) // 1

	//s5 := s2[3:4] //error: slice bounds out of range [:4] with capacity 3
}
