package main

import (
	"flag"
	"fmt"
	"log"
)

func usage() {
	fmt.Printf("Usage: climb <create|update|delete> <COMMAND> <path/to/script NOTE: create & update ONLY> ")
}

func validateArgs(args []string) {
	var argsLength = len(args)

	if argsLength < 1 {
		log.Fatal("Not enough args -- use climb help to see usage")
	}

	if argsLength > 3 {
		log.Fatal("Too many args -- use climb help to see usage")
	}

	var firstArg = args[0]

	if firstArg == "help" {
		usage()
		return
	}

	if firstArg == "delete" {
		if argsLength != 2 {
			log.Fatal("Too many args for delete command -- use climb help to see usage")
		} else {
			return
		}
	}

	if firstArg != "create" && firstArg != "update" {
		log.Fatal("Not a valid command -- use climb help to see usage")
	}
}

func cmd(args []string) {
	var command = args[0]
	var alias = args[1]

	switch command {
	case "delete":
		fmt.Printf("Delete %s", alias)
	case "create":
		fmt.Printf("Create alias %s for path %s", alias, args[2])
	case "update":
		fmt.Printf("Update alias %s with path %s", alias, args[2])
	}
}

func main() {
	flag.Parse()

	var args = flag.Args()

	validateArgs(args)
}
