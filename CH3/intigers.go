package main

import (
	"fmt"
)

func main() {

	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
	var i int8 = 127
	fmt.Println(i, i+1, i*i)

	var apples int32 = 1
	var oranges int16 = 2

	var compote = int(apples) + int(oranges)

	fmt.Println(compote)

	f := 3.141 // a float64
	z := int(f)
	fmt.Println(f, z) // "3.141 3"
	f = 1.99
	fmt.Println(int(f)) // "1"
}
