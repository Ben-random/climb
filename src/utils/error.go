package utils

import (
	"fmt"
	"log"
)

func FormatErrorMsg(err error) {
	fmt.Print("Unexpected Error\n")
	log.Fatal(err)
}

// Create and format a new Error from message without an error object
func NewErrorFromMsg(msg string) {
	fmt.Print("Unexpected Error\n")
	log.Fatal(msg)
}
