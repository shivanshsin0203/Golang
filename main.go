package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to pizza app")
	fmt.Println("Give ratinng to pizza")
	reader :=bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
}