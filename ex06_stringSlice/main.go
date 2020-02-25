package main

import (
	"fmt"
	"strings"
)

func stringConverter(inputString string) string {
	var i int
	var outputString string

	for i < len(inputString) {
		if inputString[i] == '-' || inputString[i] == '_' {
			outputString += strings.ToUpper(string(inputString[i+1]))
			i++
		} else {
			outputString += string(inputString[i])
		}
		i++
	}
	return outputString
}

func main() {
	test := "convert_this"
	fmt.Println(test)
	test2 := stringConverter(test)
	fmt.Println(test2)

	var bla string
	bla += string(test[1])
	bla += string(test[0])
	bla += string(test[len(test)-1])
	fmt.Println(bla)
}
