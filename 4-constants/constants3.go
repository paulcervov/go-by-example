package main

import "fmt"

func main() {

	const (
		a = iota + 1
		_
		b
		c
		d = c * 2 // here the counter is stopped, the next value will be copied
		t
		x = iota // why iota is 6 here?
		y
	)
	fmt.Println(a, b, c, d, t, x, y)
}
