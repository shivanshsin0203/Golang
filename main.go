package main

import (
	
	"fmt"
	
)

func main() {
    fmt.Println("Maps in Golang")
	 languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	languages["GO"] = "Golang"

	fmt.Println("Languages: ", languages)

	delete(languages, "RB")
	fmt.Println("Languages: ", languages["PY"])

	// Iterating over a map
	for key, value := range languages {
		fmt.Printf("%s: %s\n", key, value)
	}
	
}