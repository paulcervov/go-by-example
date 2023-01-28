package main

import "fmt"

func main() {
	var xyz, yz, x, z, y int

	fmt.Scan(&xyz)

	yz = xyz % 100
	x = (xyz - yz) / 100

	z = yz % 10
	y = (yz - z) / 10

	if (x != y) && (y != z) && (x != z) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
