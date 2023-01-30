package main

import "fmt"

func main() {

	var number, left, right, a, b, c int

	fmt.Scan(&number)

	right = number % 1000
	a = right / 100       //x.yx -> x
	b = (right / 10) % 10 //x.y -> x
	c = right % 10        //x
	right = a + b + c

	left = number / 1000 // xyz.xyz -> xyz
	a = left / 100       //x.yx -> x
	b = (left / 10) % 10 //x.y -> x
	c = left % 10
	left = a + b + c

	if right == left {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
