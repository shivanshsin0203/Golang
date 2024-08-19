package main

import (
	
	"fmt"
	
)

func main() {
   fmt.Println("If and else in go")

	count:= 10
	if count < 10 {
		fmt.Println("The count is less than 10")
	} else if count == 10 {
		fmt.Println("The count is equal to 10")
	} else {
		fmt.Println("The count is greater than 10")
	}

	// intialize and check the value of a variable
	if num := 10; num < 10 {
		fmt.Println("The number is less than 10")
	} else if num == 10 {
		fmt.Println("The number is equal to 10")
	} else {
		fmt.Println("The number is greater than 10")
	}
}
