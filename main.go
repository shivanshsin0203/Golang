package main

import (
	"fmt"
	
)

func main() {
	fmt.Println("Defer in Go")
    
	defer fmt.Println("This is the first defer statement")
	defer fmt.Println("This is the second defer statement")
	fmt.Println("Hello, World!")
}