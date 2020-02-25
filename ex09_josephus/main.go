package main

import "fmt"

func removeElement(input []int, index int) []int {
	if index < 0 || index >= len(input) {
		return input
	}

	copy(input[index:], input[index+1:])
	input = input[:len(input)-1]
	return input
}

func getJosephusPosition(elements int) int {
	var elemSlice []int

	for i := 1; i <= elements; i++ {
		elemSlice = append(elemSlice, i)
	}

	for {
		for i := 0; i < len(elemSlice); i++ {
			if len(elemSlice) == 1 {
				break
			}
			if i+1 >= len(elemSlice) {
				elemSlice = removeElement(elemSlice, 0)
				break
			} else {
				elemSlice = removeElement(elemSlice, i+1)
			}
		}
		if len(elemSlice) == 1 {
			break
		}
	}

	return elemSlice[0]
}

func main() {
	exampleArray := [6]int{1, 2, 3, 7, 12, 41}
	for _, v := range exampleArray {
		fmt.Println(v, getJosephusPosition(v))
	}
}
