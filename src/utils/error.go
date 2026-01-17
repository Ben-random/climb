package utils

import (
	"fmt"
	"log"
)

func FormatErrorMsg(err error) {
	fmt.Print("Unexpected Error")
	log.Fatal(err)
}

// Create and format a new error from a message and the error object
func NewEror(msg string, err error) {
	var newErr = fmt.Errorf(msg, err)
	FormatErrorMsg(newErr)
}

// Create and format a new Error from message without an error object
func NewErrorFromMsg(msg string) {
	var newErr = fmt.Errorf(msg, "")
	FormatErrorMsg(newErr)
}
