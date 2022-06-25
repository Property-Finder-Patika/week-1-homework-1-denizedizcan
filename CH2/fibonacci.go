package main

import (
	"fmt"
)

func main() {
	fmt.Printf("14th fibonacci number is %d\n", fib(14))

}

// greatest common divisor
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
