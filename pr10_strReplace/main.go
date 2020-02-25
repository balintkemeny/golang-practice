package main

import (
	"fmt"
)

func main() {
	wrongURL := "https//www.reddit.com/r/nevertellmethebots"
	rightURL := wrongURL[:5] + ":" + wrongURL[5:len(wrongURL)-4] + "odds"
	fmt.Println(rightURL)
}
