package execCmd

import (
	"climb/src/utils"
	"fmt"
	"os"
	"os/exec"
)

func Delete(alias string, dryRun bool) {
	var msg = "Are you sure you want to delete alias: " + alias + "?"
	var shouldDelete = utils.ShouldOverrideFile(msg)

	if shouldDelete == false {
		fmt.Printf("Deletion of alias %s has been aborted\n", alias)
		return
	}

	path, err := exec.LookPath(alias)
	if err != nil {
		utils.NewErrorFromMsg("Command '" + alias + "' not found in PATH")
	}

	fmt.Printf("Found command at: %s\n", path)

	if dryRun == true {
		fmt.Printf("DRY_RUN: File at path %s deleted", path)
		return
	}
	err = os.Remove(path)
	if err != nil {
		utils.NewError("Failed to delete file\n", err)
	}

	fmt.Printf("Successfully deleted alias %s\n", alias)
}
