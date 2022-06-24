package main

import (
	"fmt"
)

func main() {
	var name string = "Deniz"
	var greeting = createGreet(name)
	fmt.Printf("%s\n", greeting)
}

func createGreet(name string) string {
	greeting := "Selam " + name + " :)"
	return greeting
}
