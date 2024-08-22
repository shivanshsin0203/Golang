package main

import (
	"fmt"
	"io"
	"net/http"

)
const url = "https://nith.ac.in"
func main() {
	fmt.Println("Web request in Go");
    responce,err :=http.Get(url)
	if err != nil {
		panic(err)
	}
    defer responce.Body.Close()
	data,err := io.ReadAll(responce.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}