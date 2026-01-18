package cli

import (
	"climb/src/execCmd"
	"climb/src/utils"
)

func Cmd(args []string, dryRun bool) {
	strippedArgs := utils.StripFlagsFromArgs(args)
	utils.ValidateArgs(strippedArgs)

	var command = args[0]
	var alias = args[1]

	switch command {
	case "delete":
		execCmd.Delete(alias, dryRun)
	case "create":
		execCmd.Create(alias, args[2], dryRun)
	case "update":
		execCmd.Update(alias, args[2], dryRun)
	}
}
