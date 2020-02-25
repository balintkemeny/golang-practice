package main

import (
	"fmt"
)

var romanMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanConverter(input string) int {
	rawSlice := make([]int, len(input))
	for i, v := range input {
		rawSlice[i] = romanMap[string(v)]
	}

	if len(rawSlice) <= 1 {
		return rawSlice[0]
	}

	var output int
	for i := 0; i < len(rawSlice); i++ {
		if i != len(rawSlice)-1 {
			if rawSlice[i] >= rawSlice[i+1] {
				output += rawSlice[i]
			} else {
				output += rawSlice[i+1] - rawSlice[i]
				i++
			}
		} else {
			output += rawSlice[i]
		}

	}
	return output
}

func main() {
	fmt.Println(romanConverter("MCMIX"))
}
