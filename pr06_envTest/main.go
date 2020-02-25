package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("An error has occurred during config reading")
	}

	value, ok := viper.Get("OUTPUT").(string)

	fmt.Println(value, ok)
}
