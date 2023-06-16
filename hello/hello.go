package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main(){
	// set properties of the  predifined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// request a greeting message.
	message, err := greetings.Hello("")
	// If an error was returned, print it to the console and
	//  exit the program.

	if err != nil {
		log.Fatal(err)
	}


	// if no error was returned, print the returned message
	// to the console
	fmt.Println(message)

}