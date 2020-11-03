package main

import (
	"fmt"
)

func main() {
	name, lastname := "Nicola", "Tesla"

	fmt.Println(name, lastname)

	birth, death := 1856, 1943
	fmt.Println(birth, death)

	on, off := true, false
	fmt.Println(on, off)

	degree, ratio, heat := 10.55, 30.5, 20.
	fmt.Println(degree, ratio, heat)

	nFiles, valid, msg := 10, true, "hello"
	fmt.Println(nFiles, valid, msg)
}
