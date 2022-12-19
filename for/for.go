package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// loop with one expression
	i := 10
	for i <= 30 {
		fmt.Println(i)
		i = i + 10
	}

	// classic loop
	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}

	// infinity loop
	rand.Seed(time.Now().UnixNano())
	for {
		fmt.Println("Loop while true")
		randNum := rand.Intn(10-1) + 1
		if randNum > 5 {
			break
		}
	}

	// classic loop with continue by conditions
	for k := 1; k <= 10; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Println(k)
	}
}
