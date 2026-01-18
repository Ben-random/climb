package utils

import (
	"strings"
)

func getFlags() map[string]struct{} {
	flags := make(map[string]struct{})

	flags["dry-run"] = struct{}{}

	return flags
}

func checkFlagsForValue(val string) bool {
	flags := getFlags()

	// Remove '-' or '--' from any flags
	strippedVal := strings.TrimLeft(val, "-")
	_, exists := flags[strippedVal]

	return exists
}

func StripFlagsFromArgs(args []string) []string {
	var strippedArgs []string
	for _, val := range args {
		if checkFlagsForValue(val) == false {
			strippedArgs = append(strippedArgs, val)
		}
	}
	return strippedArgs
}
