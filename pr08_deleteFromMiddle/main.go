package main

import (
	"fmt"
)

func deleteAt(input string, index int) string {
	output := ""
	output += input[:index]
	output += input[index+1:]
	return output
}

func main() {
	fmt.Println(deleteAt("blobfish", 1))
}
