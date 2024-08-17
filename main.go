package main

import (
	
	"fmt"
	
)

func main() {
	fmt.Println("Slices in go")
	var fruits = []string{"apple", "orange", "mango", "banana", "grapes"}
	fmt.Println(fruits)
	fruits=append(fruits,"kiwi")
	fmt.Println(fruits)
	fruits=append(fruits[1:3])
	fmt.Println(fruits)
    // removing values from slices
	var index int = 2
	fruits = append(fruits[:index], fruits[index+1:]...)
	fmt.Println(fruits)
}