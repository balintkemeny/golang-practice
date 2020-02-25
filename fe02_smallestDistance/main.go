package main

import (
	"fmt"
	"sort"
)

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func fact(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fact(n-1)
}

func findSmallestDistanceFact(input []int) int {
	sort.Ints(input)
	minDist := abs(input[1] - input[0])
	for i := 1; i < len(input); i++ {
		if abs(input[i]-input[i-1]) < minDist {
			minDist = abs(input[i] - input[i-1])
		}
	}
	return fact(minDist)
}

func main() {
	fmt.Println(fact(4))
	fmt.Println(findSmallestDistanceFact([]int{-5, 99, 3425, 5, 8}))
}
