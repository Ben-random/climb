package execCmd

import (
	"climb/src/utils"
	"fmt"
	"os"
	"path/filepath"
)

func installToLocalBin(pathToBin string, alias string, dryRun bool, useSymlink bool) {
	var binDir = utils.GetBinDir()

	err := os.MkdirAll(binDir, 0755)
	if err != nil {
		utils.FormatErrorMsg(err)
	}

	absPath, err := filepath.Abs(pathToBin)
	if err != nil {
		utils.FormatErrorMsg(err)
	}

	destPath := filepath.Join(binDir, alias)

	if useSymlink {
		if dryRun {
			fmt.Printf("DRY_RUN: Create symlink from %s to %s\n", destPath, absPath)
			return
		}
		err := utils.CreateSymlink(absPath, destPath)
		if err != nil {
			utils.FormatErrorMsg(err)
		}
		fmt.Printf("Successfully created symlink: %s\n", destPath)
		return
	}

	// Default: create copy
	input, err := os.ReadFile(absPath)
	if err != nil {
		utils.FormatErrorMsg(err)
	}

	if dryRun {
		fmt.Printf("DRY_RUN: Write file from %s to %s\n", absPath, destPath)
		return
	}

	err = os.WriteFile(destPath, input, 0755)
	if err != nil {
		utils.FormatErrorMsg(err)
	}

	fmt.Printf("Successfully installed to: %s\n", destPath)
}

func CreateOrUpdate(alias string, pathToBin string, isUpdate bool, dryRun bool, useSymlink bool) {
	if !utils.IsValidAliasName(alias) {
		utils.NewErrorFromMsg("Error: Invalid alias name '" + alias + "'. Alias must start with a letter or underscore, and contain only alphanumeric characters, underscores, and hyphens")
	}
	validatePathToBin(pathToBin)

	binDir := utils.GetBinDir()
	destPath := filepath.Join(binDir, alias)

	if utils.AliasExists(alias) {
		if !isUpdate {
			// Create command but alias exists
			fmt.Printf("Error: Alias already exists\nDid you mean to use 'update' to update an existing alias?\nUsage: climb update <command> <script-path>")
			os.Exit(1)
		}

		isCurrentSymlink := utils.IsSymlink(destPath)
		if isCurrentSymlink != useSymlink {
			// Switching link types
			typeFrom := "symlink"
			typeTo := "copy"
			if useSymlink {
				typeFrom = "copy"
				typeTo = "symlink"
			}
			fmt.Printf("Alias currently exists as a %s. Converting to %s.\n", typeFrom, typeTo)
		}

		var msg = "Do you want to override alias: " + alias
		if utils.ShouldOverrideFile(msg) {
			if !dryRun {
				os.Remove(destPath)
			}
			installToLocalBin(pathToBin, alias, dryRun, useSymlink)
		} else {
			fmt.Printf("Overwrite of command %s has been aborted\n", alias)
		}
	} else if isUpdate {
		// Update command but alias doesn't exist
		fmt.Printf("Error: Alias doesn't exist\nDid you mean to use 'create' to create a new alias?\nUsage: climb create <command> <script-path>")
		os.Exit(1)
	} else {
		// Create new alias
		fmt.Printf("Creating new alias: %s\n", alias)
		installToLocalBin(pathToBin, alias, dryRun, useSymlink)
	}
}

func Create(alias string, pathToBin string, dryRun bool, useSymlink bool) {
	CreateOrUpdate(alias, pathToBin, false, dryRun, useSymlink)
}

func Update(alias string, pathToBin string, dryRun bool, useSymlink bool) {
	CreateOrUpdate(alias, pathToBin, true, dryRun, useSymlink)
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
