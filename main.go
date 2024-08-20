package main

import (
	"fmt"
	
)

func main() {
     fmt.Println("Lopps in golang")
	 days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	//  for i := 0; i < len(days); i++ {
	// 	 fmt.Println(days[i])
	//  }
	//  for i:= range days {
	// 	 fmt.Println(days[i])
	//  }
	 for index, value := range days {
		 fmt.Printf("Index: %v, Value: %v\n", index, value)
	 }
	  value :=1

	 for value < 10 {
        if value ==3{
			goto lco
		}
		if(value == 5){
			value++
			continue
		}
		 fmt.Println(value)
		 value++
	 }
	 lco:
	     fmt.Println("Jumping to LCO")
}
