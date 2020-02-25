package main

import "fmt"

func checkEnding(input, ending string) bool {
	if len(ending) == 0 {
		return true
	} else if len(input) == 0 {
		return false
	}

	var endCheck bool = true
	for i := len(input) - len(ending); i < len(input); i++ {
		fmt.Println("Input: ", input[i])
		fmt.Println("Ending: ", ending[i-len(input)+len(ending)])
		if input[i] != ending[i-len(input)+len(ending)] {
			endCheck = false
		}
	}
	return endCheck
}

func main() {
	/*
		len(input) = 6
		len(ending) = 3
		start iteration at input index: 3
		iterate until index: 6

		i = input - ending
		j = 0
	*/
	fmt.Println(checkEnding("aladar", "abar"))
}
