package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	CmdCreate = "create"
	CmdUpdate = "update"
	CmdDelete = "delete"
	CmdHelp   = "help"
)

func usage() {
	fmt.Println("Usage: climb <command> <alias> [script-path]")
	fmt.Println("\nCommands:")
	fmt.Println("  create <alias> <script-path>  Create a new alias for a script")
	fmt.Println("  update <alias> <script-path>  Update an existing alias")
	fmt.Println("  delete <alias>                Delete an existing alias")
	fmt.Println("  help                          Show this help message")
	fmt.Println("\nOptions:")
	fmt.Println("  --dry-run                     Preview changes without modifying files")
	fmt.Println("  --symlink                     Create a symlink instead of copying a script/binary")
}

func ValidateArgs(args []string) {
	argsLength := len(args)

	if argsLength < 1 {
		log.Fatal("Error: No command provided.\nRun 'climb help' for usage information.")
	}

	command := strings.TrimSpace(strings.ToLower(args[0]))

	if command == "" {
		log.Fatal("Error: Command cannot be empty.\nRun 'climb help' for usage information.")
	}

	if command == CmdHelp {
		usage()
		os.Exit(0)
	}

	validCommands := map[string]bool{
		CmdCreate: true,
		CmdUpdate: true,
		CmdDelete: true,
	}

	if !validCommands[command] {
		log.Fatalf("Error: Unknown command '%s'.\nValid commands are: create, update, delete, help\nRun 'climb help' for more information.", args[0])
	}

	if command == CmdDelete {
		if argsLength < 2 {
			log.Fatal("Error: 'delete' command requires an alias.\nUsage: climb delete <alias>")
		}
		if argsLength > 2 {
			log.Fatalf("Error: 'delete' command accepts only one argument (alias), but received %d arguments.\nUsage: climb delete <alias>", argsLength-1)
		}

		alias := strings.TrimSpace(args[1])
		if alias == "" {
			log.Fatal("Error: Alias cannot be empty or whitespace.")
		}

		return
	}

	if command == CmdCreate || command == CmdUpdate {
		if argsLength < 2 {
			log.Fatalf("Error: '%s' command requires an alias.\nUsage: climb %s <alias> <script-path>", command, command)
		}
		if argsLength < 3 {
			log.Fatalf("Error: '%s' command requires a script path.\nUsage: climb %s <alias> <script-path>", command, command)
		}
		if argsLength > 3 {
			log.Fatalf("Error: '%s' command accepts exactly 2 arguments (alias and script-path), but received %d arguments.\nUsage: climb %s <alias> <script-path>",
				command, argsLength-1, command)
		}

		alias := strings.TrimSpace(args[1])
		if alias == "" {
			log.Fatal("Error: Alias cannot be empty or whitespace.")
		}

		scriptPath := strings.TrimSpace(args[2])
		if scriptPath == "" {
			log.Fatal("Error: Script path cannot be empty or whitespace.")
		}

		return
	}
}
