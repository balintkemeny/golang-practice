package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%-10s $%4d\n", "Blah", 11)
	fmt.Printf("%-10s $%4d\n", "Blablah", 122)

	var distance, speed float32 = 53.2, 65433.1
	fmt.Println(distance, speed)

	var (
		name1 string = "John"
		name2 string = "Joe"
	)
	fmt.Println(name1, name2)
}
