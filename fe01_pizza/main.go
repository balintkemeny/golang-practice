/*
Create a function that takes a string(word) and a number(width) as a parameter,
and creates a string array / list. If you write the array elements below eachother,
you should get a rectangle shaped form with the width of the number parameter and if
you start reading in the top left corner of it moving right or down,
in the end you should always read the original word.

Example cases:
wordReactangle('pizza', 3);
Should return
[
['piz'],
['izz'],
['zza']
]
*/

package main

import "fmt"

func pizzaParser(input string, size int) []string {
	var output []string
	if size > len(input) {
		return output
	}
	for i := 0; i < size; i++ {
		output = append(output, input[i:i+size])
	}
	return output
}

func main() {
	var test []string
	test = pizzaParser("blobfishy", 5)

	for i := 0; i < len(test); i++ {
		fmt.Println(test[i])
	}
}
