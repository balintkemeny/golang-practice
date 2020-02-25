package main

import (
	"fmt"
)

func main() {
	testSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("before: ", testSlice)

	copy(testSlice[2:], testSlice[3:])
	testSlice = testSlice[:len(testSlice)-1]
	fmt.Println("after : ", testSlice)
}
