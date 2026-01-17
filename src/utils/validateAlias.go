package utils

import "os/exec"

func AliasExists(alias string) bool {
	_, err := exec.LookPath(alias)
	if err != nil {
		return false
	}
	return true
}
