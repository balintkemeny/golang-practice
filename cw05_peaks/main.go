package main

import (
	"fmt"
)

//PosPeaks data structure containing an array of peak positions and an array of peak heights
type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func pickPeaks(input []int) PosPeaks {
	var output PosPeaks
	if len(input) == 0 {
		output.Pos = []int{}
		output.Peaks = []int{}
		return output
	}
	for i := 1; i < len(input)-1; i++ {
		if input[i] > input[i-1] {
			if input[i] > input[i+1] {
				output.Pos = append(output.Pos, i)
				output.Peaks = append(output.Peaks, input[i])
			} else {
				j := 0
				for input[i+j] == input[i+j+1] {
					if i+j+1 < len(input)-1 {
						j++
					} else {
						break
					}
				}
				if input[i] > input[i+j+1] && i+j+1 < len(input) {
					output.Pos = append(output.Pos, i)
					output.Peaks = append(output.Peaks, input[i])
				}
			}
		}
	}
	if len(output.Pos) == 0 {
		output.Pos = []int{}
		output.Peaks = []int{}
		return output
	}
	return output
}

func main() {
	fmt.Println(pickPeaks([]int{1, 1, 1}))
}
