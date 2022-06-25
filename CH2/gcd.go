package main

import (
	"fmt"
)

func main() {
	fmt.Printf("GCD of 84 and 142 is %d\n", gcd(84, 142))

}

// greatest common divisor
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
