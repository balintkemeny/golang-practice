package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("There is a cavern entrance here and a path to the east. What do you do?")
	command, _ := reader.ReadString('\n')
	command = command[:len(command)-1]

	switch command {
	case "go east":
		fmt.Println("You head up the mountain path to the east.")
	case "enter cave", "go inside":
		fmt.Println("You enter the dimly lit cavern.")
		fallthrough
	case "read sign":
		fmt.Println("Danger! Here be trolls.")
	default:
		fmt.Println("Unavailable command!")
	}
}
