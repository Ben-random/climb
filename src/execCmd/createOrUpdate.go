package execCmd

import (
	"climb/src/utils"
	"fmt"
	"os"
	"path/filepath"
)

func installToLocalBin(pathToBin string, alias string, isUpdate bool, dryRun bool) {
	var binDir = utils.GetBinDir()

	if isUpdate == false {
		err := os.MkdirAll(binDir, 0755)
		if err != nil {
			utils.FormatErrorMsg(err)
		}
	}

	absPath, err := filepath.Abs(pathToBin)
	if err != nil {
		utils.FormatErrorMsg(err)
	}

	input, err := os.ReadFile(absPath)
	if err != nil {
		utils.FormatErrorMsg(err)
	}

	destPath := filepath.Join(binDir, alias)
	if dryRun == true {
		fmt.Printf("DRY_RUN: Write file from %s to %s\n", absPath, destPath)
		return
	}
	err = os.WriteFile(destPath, input, 0755)
	if err != nil {
		utils.FormatErrorMsg(err)
	}

	fmt.Printf("Successfully installed to: %s\n", destPath)
}

func CreateOrUpdate(alias string, pathToBin string, canOverrideExisting bool, dryRun bool) {
	if !utils.IsValidAliasName(alias) {
		utils.NewErrorFromMsg("Error: Invalid alias name '" + alias + "'. Alias must start with a letter or underscore, and contain only alphanumeric characters, underscores, and hyphens")
	}
	validatePathToBin(pathToBin)

	if utils.AliasExists(alias) {
		if canOverrideExisting == true {
			// Update alias for new bin
			var msg = "Do you want to override alias: " + alias

			if utils.ShouldOverrideFile(msg) == true {
				installToLocalBin(pathToBin, alias, true, dryRun)
			} else {
				fmt.Printf("Overwrite of command %s with binary at %s has been aborted\n", alias, pathToBin)
			}
		} else {
			fmt.Printf("Error: Alias already exists\nDid you mean to use 'update' to update an existing alias?\nUsage: climb update <command> <script-path>")
			os.Exit(1)
		}
	} else if canOverrideExisting == false {
		// Create new alias for bin
		fmt.Printf("Creating new alias: %s\n", alias)
		installToLocalBin(pathToBin, alias, false, dryRun)
	} else {
		fmt.Printf("Error: Alias doesn't exist\nDid you mean to use 'create' to create a new alias?\nUsage: climb create <command> <script-path>")
	}
}

func Create(alias string, pathToBin string, dryRun bool) {
	CreateOrUpdate(alias, pathToBin, false, dryRun)
}

func Update(alias string, pathToBin string, dryRun bool) {
	CreateOrUpdate(alias, pathToBin, true, dryRun)
}

func validatePathToBin(pathToBin string) {
	fileInfo, err := os.Stat(pathToBin)
	if err != nil {
		utils.NewErrorFromMsg("Error: file " + pathToBin + " could not be found\n" + err.Error() + "\n")
	}
	if fileInfo.Mode().Perm()&0111 == 0 {
		utils.NewErrorFromMsg("Error: file " + pathToBin + " is not executable")
	}
}
