package main

import (
	"fmt"
	"math"
)

func main() {
	const s string = "constant"
	fmt.Println(s) // "constant"

	const ch = 'a'          // rune type
	fmt.Println(ch)         // 97
	fmt.Println(string(ch)) // "a"

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)           //6e+11
	fmt.Println(int64(d))    //600000000000
	fmt.Println(math.Sin(n)) //-0.28470407323754404

	const pi float64 = 3.1415
	//pi = 2.7182 //Error: cannot assign to pi

	fmt.Println(math.Sin(ch)) //0.37960773902752165
	const ch2 = "b"           // string type
	//fmt.Println(math.Sin(ch2)) //Error: cannot use ch2 (untyped string constant "b") as float64 value in argument to math.Sin
}
