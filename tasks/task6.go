package main

import "fmt"

// It is required to determine whether a given year is a leap year
func main() {

	var year int
	fmt.Scan(&year)

	if (year%400 == 0) || ((year%4 == 0) && (year%100 != 0)) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
