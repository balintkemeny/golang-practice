package main

import (
	"fmt"
)

func main() {
	randomSlice := []int{5, 4, 3, 2, 1}
	fmt.Println(randomSlice)
	randomSlice = append(randomSlice, 3)
	fmt.Println(randomSlice)
	randomSlice = randomSlice[:len(randomSlice)-1]
	fmt.Println(randomSlice)

	randomSlice2 := []string{"random", "cheese", "brie"}
	fmt.Println(randomSlice2[1][1:5])
}
