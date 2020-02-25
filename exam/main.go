package main

import (
	"fmt"
	"sort"
)

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func sortDescInt(input []int) []int {
	for i := 0; i < len(input); i++ {
		swapped := false
		for j := 1; j < len(input)-i; j++ {
			if input[j-1] < input[j] {
				swap(&(input[j-1]), &(input[j]))
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return input
}

func pyramidSort(input []int) []int {
	var evens []int
	var odds []int

	for _, v := range input {
		if v%2 == 0 {
			evens = append(evens, v)
		} else {
			odds = append(odds, v)
		}
	}

	sort.Ints(odds)
	evens = sortDescInt(evens)

	var out []int
	out = odds

	for i := 0; i < len(evens); i++ {
		out = append(out, evens[i])
	}

	return out
}

func main() {
	testSlice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(pyramidSort(testSlice))
}
