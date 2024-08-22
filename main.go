package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Post request in go")
}
func HandlePostReq(){
	fmt.Println("Post request in go")
	const url = "http://localhost:8080"
	requestBody :=strings.NewReader(`{"title":"New post","content":"This is a new post"}`)
	res,err := http.Post(url,"application/json",requestBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content,err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}