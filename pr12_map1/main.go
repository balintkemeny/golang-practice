package main

import (
	"fmt"
)

func main() {
	var m map[int]string
	m = make(map[int]string)
	m[42] = "blobfish"
	fmt.Println(m[42])

	m2 := map[int]string{3: "blob", 4: "fish"}
	fmt.Println(m2[3])
	m2[5] = "time"
	fmt.Println(m2[5])
	delete(m2, 5)
	fmt.Println(m2)
}
