package main

import "fmt"

func main() {
	var (
		name   string
		age    int
		famous bool
	)

	name = "Newton"
	age = 84
	famous = true

	name = "Somebody"
	age = 20
	famous = false

	var prevName string
	prevName = name

	name = "Einstein"

	fmt.Println("previous name:", prevName)
	fmt.Println("current name:", name)

	fmt.Println(name, age, famous)
}
