package main

import "fmt"

func main() {
	var a, b, c int = 3, 4, 5
	var d [2]string
	d[0] = "moon"
	d[1] = "sun"

	testArray := [6]int{2, 3, 4, 5, 6, 7}

	fmt.Println(d)
	fmt.Println(a, b, c, testArray)
}