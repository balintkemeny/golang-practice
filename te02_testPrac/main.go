package main

import (
	"fmt"
	"net/http"
)

func getStatusCode(URL string) int {
	resp, err := http.Get(URL)
	if err != nil {
		return -1
	}
	return resp.StatusCode
}

func main() {
	fmt.Println(getStatusCode("https://www.google.com"))
}
