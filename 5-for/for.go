package main

import (
	"fmt"
	"time"
)

func main() {
	// loop with one expression
	i := 10
	for i <= 30 {
		fmt.Println("Loop with one expression, number is: ", i)
		i = i + 10
	}

	// classic loop
	for j := 1; j <= 3; j++ {
		fmt.Println("Classic loop, number is:", j)
	}

	// infinity loop
	for {
		timeStamp := time.Now().UnixNano()
		randNum := (timeStamp % 10000) / 1000
		fmt.Println("Infinity loop while true, random number is:", randNum)
		if randNum > 5 {
			break
		}
	}

	// classic loop with continue by conditions
	for k := 1; k <= 10; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Println("Classic loop with continue by conditions, number is: ", k)
	}
}
