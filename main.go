package main

import (
	"fmt"
	"io"
	
	"os"
)

func main() {
	fmt.Println("Files in golang")
	content := "Hello, World!"
	file,err := os.Create("./test.txt")
	if err != nil {
		panic(err)
	}
	length,err := io.WriteString(file,content)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Length: %d",length)
	readFile("./test.txt")
	defer file.Close()
}
func readFile (fileName string) {
	data,err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("The data read from the file is: ",string(data))
}