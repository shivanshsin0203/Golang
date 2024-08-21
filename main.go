package main

import (
	"fmt"
	
)

func main() {
	fmt.Println("Meathods in Go")
     shivam := User{"Shivam","Kumar","awd@gmail.com",21,true}
	 fmt.Println("User is",shivam)
	 shivam.GetStatus()
}

type User struct{
	FirstName string
	LastName string
	Email string
	Age int
	isActive bool
}
func (u User) GetStatus() {
	fmt.Println("Is user active",u.isActive)
}