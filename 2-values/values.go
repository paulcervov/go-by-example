package main

import "fmt"

func main() {
	fmt.Println("Go" + "lang")

	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// strings
	var s1 = "Hello\table"
	fmt.Println(s1) //Hello   able

	var s2 = `Hello\table`
	fmt.Println(s2) //Hello\table

	var s3 = 'H'            //rune
	fmt.Println(s3)         //72
	fmt.Println(string(s3)) //H

	fmt.Println(s1[0])         //72
	fmt.Println(string(s1[0])) //H

	// string in apostrophes can be multiline
	var s4 = `
One
two
three`
	fmt.Println(s4)
	//One
	//two
	//three
}
