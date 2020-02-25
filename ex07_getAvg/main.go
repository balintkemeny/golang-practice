package main

import "fmt"

func getAvg(input []int) int {
	var sum int
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}
	return sum / len(input)
}

func getAvg2(input []int) int {
	var sum int
	for _, v := range input {
		sum += v
	}
	return sum / len(input)
}

func main() {
	testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(getAvg(testSlice), getAvg2(testSlice))
}
