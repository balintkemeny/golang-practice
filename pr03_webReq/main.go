package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

func main() {
	resp, _ := http.Get("https://www.google.com/")
	fmt.Println(resp.Status)
	fmt.Printf("%T\n", resp.Status)
	fmt.Println(resp.StatusCode)
	fmt.Printf("%T\n", resp.StatusCode)
	fmt.Println(reflect.TypeOf(resp.StatusCode).String())

	fmt.Println(resp.Header)

	rawResponseBody, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	stringifiedResponseBody := string(rawResponseBody)
	fmt.Println(stringifiedResponseBody)
}
