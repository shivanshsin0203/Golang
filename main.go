package main

import (
	
	"fmt"
	
)

func main() {
   fmt.Println("Stgructs in go lang")
     shivansh :=User{"Shivansh","singhshivansh12may@gmail.com",21,true}
	fmt.Println(shivansh)
	
}
type User struct {
	Name string
	Email string
	Age int
	Status bool
}