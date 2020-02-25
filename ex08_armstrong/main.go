package main

import (
	"fmt"
	"math"
)

func isArmstrong(input int) {
	var runner int = input
	var digits []int
	var sum int

	for runner > 0 {
		digits = append(digits, runner%10)
		runner /= 10
	}

	for _, v := range digits {
		sum += int(math.Pow(float64(v), float64(len(digits))))
	}

	if sum == input {
		fmt.Printf("The %d is an Armstrong number.\n", input)
	} else {
		fmt.Printf("The %d is not an Armstrong number.\n", input)
	}
}

func main() {
	isArmstrong(1634)
	isArmstrong(1633)
	isArmstrong(153)
	isArmstrong(154)
	isArmstrong(371)
}
