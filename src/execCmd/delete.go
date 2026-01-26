package execCmd

import (
	"climb/src/utils"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func isCmdInBinDir(path string) bool {
	var binPath = utils.GetBinDir()
	if strings.HasPrefix(path, binPath) {
		return true
	}
	return false
}

func Delete(alias string, dryRun bool) {
	path, err := exec.LookPath(alias)
	if err != nil {
		utils.NewErrorFromMsg("Command '" + alias + "' not found in PATH")
	}

	if !isCmdInBinDir(path) {
		utils.NewErrorFromMsg("Command '" + alias + "' not found in bin directory - cannot safely delete")
	}

	var msg = "Are you sure you want to delete alias: " + alias + "?"
	var shouldDelete = utils.ShouldOverrideFile(msg)

	if shouldDelete == false {
		fmt.Printf("Deletion of alias %s has been aborted\n", alias)
		return
	}

	fmt.Printf("Found command at: %s\n", path)

	if dryRun == true {
		fmt.Printf("DRY_RUN: File at path %s deleted", path)
		return
	}
	err = os.Remove(path)
	if err != nil {
		fmt.Printf("Failed to delete command at path %s\n", path)
		utils.FormatErrorMsg(err)
	}

	fmt.Printf("Successfully deleted alias %s\n", alias)
}
