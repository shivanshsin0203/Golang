package main

import (
	"encoding/json"
	"fmt"
)
func main() {
	fmt.Println("Generate json in go")
    handleJson()
}
type course struct{
	Name string
	Price int
	Platform string
	Password string
	Tags []string
}
func handleJson(){
	courses := []course{
		{ "Go",  0,  "Youtube",  "1234",  []string{"Go", "Programming"}},
		{ "Python", 0, "Youtube",  "1234",  []string{"Python", "Programming"}},
		{ "C++",  0,  "Youtube",  "1234", nil },
	}
	finalJson,err := json.MarshalIndent(courses, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}