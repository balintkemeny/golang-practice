package main

import (
	"fmt"
)

func parenthesesValidator(input string) bool {
	var counter int

	for i := 0; i < len(input); i++ {
		if string(input[i]) == "(" {
			counter++
		} else if string(input[i]) == ")" {
			counter--
		}

		if counter < 0 {
			break
		}
	}

	return counter == 0
}

func main() {
	fmt.Println(parenthesesValidator(")()()("))
}
