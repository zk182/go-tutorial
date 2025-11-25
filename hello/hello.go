package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
    message := greetings.Hello("Nico")
    fmt.Println(message)
}