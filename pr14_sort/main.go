package main

import (
	"fmt"
	"sort"
)

func main() {
	test := []int{2, 9, 34, 5, 65432, -3}
	sort.Ints(test)
	fmt.Println(test)
}
