package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
    log.SetFlags(0)

	names := []string{"Nico", "Matías", "Gabi"}
	
    message, err := greetings.Hellos(names)
	for _, greet := range message {
		fmt.Println(greet)
	}
	if err != nil {
		log.Fatal(err)
	}
}