package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:3000/")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp.StatusCode)
	}
}
