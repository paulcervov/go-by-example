package main

import "fmt"

// The sequence consists of natural numbers and ends with the number 0.
// Determine the number of elements of this sequence that are equal to its largest element.
func main() {

	var num, maxNum, maxNumCnt int

	for fmt.Scan(&num); num > 0; fmt.Scan(&num) {

		if num == maxNum {
			maxNumCnt += 1
		}

		if num > maxNum {
			maxNum = num
			maxNumCnt = 1
		}
	}

	fmt.Println(maxNumCnt)
}
