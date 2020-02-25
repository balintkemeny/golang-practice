package main

import (
	"fmt"
)

func stringReverser(input string) string {
	output := ""

	for i := len(input) - 1; i >= 0; i-- {
		//output += string(input[i])
		output += input[i : i+1]
	}

	return output
}

func main() {
	sample := "blobfish"
	fmt.Println(stringReverser(sample))
}
