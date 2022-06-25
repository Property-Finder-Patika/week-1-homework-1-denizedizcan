package main

import (
	"fmt"
)

func main() {
	a := 1
	b := &a

	fmt.Println(*b)
	*b = 2
	fmt.Println(a)

	fmt.Println(f(b))

}
func f(x *int) int {
	*x++
	return *x
}
