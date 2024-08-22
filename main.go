package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
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
	content := string(data)
	file, err := os.Create("nith.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	length,err := io.WriteString(file,content)
	if err != nil {
		panic(err)
	}
	fmt.Printf("File created with %v characters",length)
}