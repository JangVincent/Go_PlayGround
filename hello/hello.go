package main

import (
	"fmt"

	"phantola.com/greetings"

	"log"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	message, err := greetings.Hello("Phantola")

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(message)
}
