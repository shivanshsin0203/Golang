package main

import (
	"fmt"
	
)

func main() {
     fmt.Println("Functions in Go")
     preset()
     fmt.Println(adder(2,3))
	 proAddresult,proAddStatus := proadder(1,2,3,4,5)
	 fmt.Println(proAddresult,proAddStatus)
}
func proadder(values ... int) (int,string){
	 result :=0
	 for _,v := range values{
		 result += v
	 }
	 return result,"Success"
}

func adder(a int, b int) int {
	return a + b
}
func preset(){
	fmt.Println("This is a preset function")
}
