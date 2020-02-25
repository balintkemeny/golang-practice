package main

import "fmt"

type ingredient struct {
	name         string
	needsCooling bool
	inStock      int
}

func ingredientPrinter(ingredients []ingredient) {
	var maxNameLength int = 10
	for _, v := range ingredients {
		if len(v.name) > maxNameLength {
			maxNameLength = len(v.name)
		}
	}

	//Print header
	for j := 0; j < 3; j++ {
		if j%2 == 0 {
			fmt.Printf("+")
			for i := 0; i < maxNameLength+2; i++ {
				fmt.Printf("-")
			}
			fmt.Printf("+---------------+----------+\n")
		} else {
			fmt.Printf("| Ingredient")
			for i := 0; i < maxNameLength-len("Ingredient")+1; i++ {
				fmt.Printf(" ")
			}
			fmt.Printf("| Needs cooling | In stock |\n")
		}
	}

	//Print body
	for _, v := range ingredients {
		fmt.Printf("| %s", v.name)
		for j := 0; j < maxNameLength-len(v.name)+1; j++ {
			fmt.Printf(" ")
		}
		fmt.Printf("| ")
		if v.needsCooling {
			fmt.Printf("Yes           | ")
		} else {
			fmt.Printf("No            | ")
		}
		if v.inStock <= 0 {
			fmt.Printf("-        |\n")
		} else if v.inStock < 10 {
			fmt.Printf("%d        |\n", v.inStock)
		} else {
			fmt.Printf("%d       |\n", v.inStock)
		}
	}

	//Print footer
	fmt.Printf("+")
	for i := 0; i < maxNameLength+2; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("+---------------+----------+\n")
}

func main() {
	ingredientPrinter([]ingredient{
		{"banana", false, 2},
		{"venison", true, 1},
		{"Captain Morgan Dark Rum", false, 15},
	})
}
