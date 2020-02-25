package main

import "fmt"

func main() {
	var sum int = 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum)

	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
	}

	var testFloat float64 = 3.3334422
	for {
		testFloat /= 2
		fmt.Println(testFloat)
		if testFloat <= 0.5 {
			break
		}
	}
}
