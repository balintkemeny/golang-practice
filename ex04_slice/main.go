package main

import "fmt"

func intModifier(numberPtr *int) {
	*numberPtr = 11
}

func arrayModifier(arrayPtr *[5]int) {
	(*arrayPtr)[4] = 0
}

func bestPracticeModifier(inputSlice []int) {
	length := len(inputSlice)
	inputSlice[length-1] = 66
}

func main() {
	testInt := 6
	fmt.Println(testInt)
	intModifier(&testInt)
	fmt.Println(testInt)

	testArray := [5]int{1, 2, 3, 4, 5}

	fmt.Printf("Type %T\n", testArray)
	fmt.Println(testArray)
	arrayModifier(&testArray)
	fmt.Println(testArray)

	var testSlice []int = testArray[0:5]
	testSlice = append(testSlice, 99)
	fmt.Println(testSlice)
	bestPracticeModifier(testSlice)
	fmt.Println(testSlice)
	fmt.Println(testArray)

}
