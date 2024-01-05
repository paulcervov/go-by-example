package main

import (
	"fmt"
	"time"
)

func main() {
	timeStamp := time.Now().UnixNano()
	randNum := (timeStamp % 10000) / 1000
	fmt.Println(randNum)
}
