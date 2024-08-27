package main

import (
	"fmt"
)

// Define an interface
type Speaker interface {
	Speak() string
}

// Define a type Dog that implements the Speaker interface
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof! My name is " + d.Name
}

// Define another type Cat that implements the Speaker interface
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow! My name is " + c.Name
}

// Define a function that takes a Speaker interface as a parameter
func SaySomething(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	// Create instances of Dog and Cat
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	// Both Dog and Cat satisfy the Speaker interface, so they can be passed to SaySomething
	SaySomething(dog)
	SaySomething(cat)
}
