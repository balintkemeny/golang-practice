package main

import (
	"fmt"
	"strings"
)

func substrList(keyStr string, strList []string) int {
	for i, v := range strList {
		if strings.Contains(v, keyStr) {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(substrList("ching", []string{"this", "is", "what", "I'm", "searching", "in"}))
	fmt.Println(substrList("foo", []string{"this", "is", "what", "I'm", "searching", "in"}))
}
