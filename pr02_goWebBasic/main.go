package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I can use formatted text here. Hello, World by the way!\n")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", nil)
}
