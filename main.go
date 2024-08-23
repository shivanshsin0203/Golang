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
	 handleError(err)
	defer res.Body.Close()
	content,err := io.ReadAll(res.Body)
	handleError(err)
	fmt.Println(string(content))
}
func handleError(err error){
	if err != nil{
		fmt.Println("Error:",err)
	}
}