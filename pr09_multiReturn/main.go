package main

import (
	"fmt"
	"math"
)

func solveQuEq(a, b, c int) (int, int) {
	var out1, out2 int
	out1 = (-b + int(math.Pow(float64((b*b-4*a*c)), 0.5))) / (2 * a)
	out2 = (-b - int(math.Pow(float64((b*b-4*a*c)), 0.5))) / (2 * a)

	return out1, out2
}

func main() {
	res1, res2 := solveQuEq(1, 0, -4)
	fmt.Println(res1, ", ", res2)
}
