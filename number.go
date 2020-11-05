package main

import (
	"fmt"
	"os"
	"strconv"
)

func feetToMeters() {
	arg := os.Args[1]
	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * 0.3048

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}

func main() {
	feetToMeters()
}
