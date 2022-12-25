package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)           //6e+11
	fmt.Println(int64(d))    //600000000000
	fmt.Println(math.Sin(n)) //-0.28470407323754404

	const pi float64 = 3.1415
	//pi = 2.7182 //Error: cannot assign to pi

	const (
		aa int = 0
		bb
		cc
	)
	fmt.Println(aa, bb, cc) // 0 0 0

	const (
		aaa int = iota + 1
		bbb
		ccc
		ddd = 33
		fff
	)
	fmt.Println(aaa, bbb, ccc, ddd, fff) // 1 2 3 33 33

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		_ // пропускаем 7
		Add
		Add2 = iota + 2
		Add3
	)
	fmt.Println(Saturday)
	fmt.Println(Add)

	const (
		u         = iota * 42 //index 0 (0 * 42)
		v float64 = iota * 42 //index 1.0 (1.0 * 42)
		w         = iota * 42 //index 2 (2 * 42)
	)
	fmt.Println(u) //0
	fmt.Println(v) //42
	fmt.Println(w) //84

	// the constants below are not in the same block
	const x = iota // x == 0
	const y = iota // y == 0
}
