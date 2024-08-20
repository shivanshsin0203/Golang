package main

import (
	"fmt"
	
)

func main() {
     fmt.Println("Lopps in golang")
	 days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	 for i := 0; i < len(days); i++ {
		 fmt.Println(days[i])
	 }
	 for i:= range days {
		 fmt.Println(days[i])
	 }
}
