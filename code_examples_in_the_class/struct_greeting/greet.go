package main

import "fmt"

type Person struct {
	name string
}

func (p Person) greet() string {
	return "Selam " + p.name + " :)"
}

func (p Person) greetSomeBody(q Person) string {
	return "Selam " + p.name + " to " + q.name + " :)"
}

func main() {
	greeter := Person{"Deniz"}
	friend := Person{"Alper"}

	var greeting = greeter.greet()
	var greetings = greeter.greetSomeBody(friend)
	fmt.Printf("%s\n", greeting)
	fmt.Printf("%s\n", greetings)
}
