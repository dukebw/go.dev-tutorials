package main

import (
	"fmt"
	"log"

	"github.com/dukebw/go.dev-tutorials/create-module/greetings"
)

func main() {
	// Set properties of the predefined Logger, including the log entry prefix
	// and a flag to disable printing the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "John", "Mary"}

	// Request a greeting message.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
