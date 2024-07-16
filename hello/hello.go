package main

import (
	"fmt"

	"log"

	"phantola.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
			log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.

	for _, message := range messages {
		fmt.Println(message)
	}

	// Test the format specifiers
	fmt.Printf("%q", "Hello\tworld.") // Show literal with escape characters
	fmt.Println()
	fmt.Printf("%#q", "Hello\tworld.") // Show literal but evaluate escape characters with backtic
	fmt.Println()
	fmt.Printf("%v", "Hello\tworld.") // Show value
}
