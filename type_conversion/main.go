package main

import "fmt"

func main() {
	var apple int
	var orange int32

	apple = int(orange)

	// isDelicious := bool(orange)

	orange = 65
	color := string(orange)
	fmt.Println(color)

	_ = apple
}
