package execCmd

import (
	"climb/src/utils"
	"fmt"
	"os"
	"os/exec"
)

func Delete(alias string) {
	var msg = "Are you sure you want to delete alias: " + alias + "?"
	var shouldDelete = utils.ShouldOverrideFile(msg)

	if shouldDelete == false {
		fmt.Printf("Deletion of alias %s has been aborted", alias)
		return
	}

	path, err := exec.LookPath(alias)
	if err != nil {
		utils.NewErrorFromMsg("Command '" + alias + "' not found in PATH")
	}

	fmt.Printf("Found command at: %s\n", path)

	err = os.Remove(path)
	if err != nil {
		utils.NewEror("Failed to delete file\n", err)
	}

	fmt.Printf("Successfully deleted alias %s\n", alias)
}
