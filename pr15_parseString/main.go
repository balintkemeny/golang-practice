package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var gNeedsReplacing = []string{"shit", "crap"}
var gReplacements = []string{"poop", "bad-stuffy-stuff"}

func stringReplacer(input string) string {
	var inSlice []string
	inSlice = strings.Split(input, " ")

	for i := 0; i < len(inSlice); i++ {
		for j := 0; j < len(gNeedsReplacing); j++ {
			if inSlice[i] == gNeedsReplacing[j] {
				inSlice[i] = gReplacements[rand.Intn(2)]
			}
		}
	}

	var output string
	for _, v := range inSlice {
		output += v + " "
	}
	return output
}

func stringReplacer2(input string) string {
	for _, v := range gNeedsReplacing {
		input = strings.Replace(input, v, gReplacements[rand.Intn(2)], -1)
	}
	return input
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(stringReplacer("Please replace the shit in this crap"))
	fmt.Println(stringReplacer2("It's still the same shit that is crap"))
}
