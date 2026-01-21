package utils

import (
	"os/exec"
	"regexp"
)

func AliasExists(alias string) bool {
	_, err := exec.LookPath(alias)
	if err != nil {
		return false
	}
	return true
}

func IsValidAliasName(alias string) bool {
	// Alias must:
	// 1. Not be empty
	// 2. Start with a letter or underscore
	// 3. Contain only alphanumeric characters, underscores, and hyphens
	pattern := `^[a-zA-Z_][a-zA-Z0-9_-]*$`
	match, _ := regexp.MatchString(pattern, alias)
	return match
}
