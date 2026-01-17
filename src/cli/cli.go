package cli

import (
	"climb/src/execCmd"
)

func Cmd(args []string) {
	var command = args[0]
	var alias = args[1]

	switch command {
	case "delete":
		execCmd.Delete(alias)
	case "create":
		execCmd.Create(alias, args[2])
	case "update":
		execCmd.Update(alias, args[2])
	}
}
