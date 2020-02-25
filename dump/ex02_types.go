package main

import "fmt"

func main() {
	a := 10
	b := 10.312

	c, python, java := true, false, "no!"

	fmt.Printf("%T, %T, %T, %T, %T\n", a, b, c, python, java)
}