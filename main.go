package main

import (
	
	"fmt"
	
	"time"
)

func main() {
	fmt.Println("Learn Go - 1.1")
	date := time.Now()
	fmt.Println("Today is", date)
	fmt.Println(date.Format("02-01-2006 15:04:05 Monday"))
}