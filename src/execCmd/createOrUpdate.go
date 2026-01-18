package execCmd

import (
	"climb/src/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func getBinDir(home string) string {
	if runtime.GOOS == "windows" {
		return filepath.Join(home, "AppData", "local", "Microsoft", "WindowsApps")
	}
	return filepath.Join(home, ".local", "bin")
}

func installToLocalBin(pathToBin string, isUpdate bool, dryRun bool) {
	home, err := os.UserHomeDir()
	if err != nil {
		utils.FormatErrorMsg(err)
	}

	var binDir = getBinDir(home)

	if isUpdate == false {
		err = os.MkdirAll(binDir, 0755)
		if err != nil {
			utils.FormatErrorMsg(err)
		}
	}

	input, err := os.ReadFile(pathToBin)
	if err != nil {
		utils.FormatErrorMsg(err)
	}

	destPath := filepath.Join(binDir, filepath.Base(pathToBin))
	if dryRun == true {
		fmt.Printf("DRY_RUN: Write file from %s to %s\n", pathToBin, destPath)
		return
	}
	err = os.WriteFile(destPath, input, 0755)
	if err != nil {
		utils.FormatErrorMsg(err)
	}

	fmt.Printf("Successfully installed to: %s\n", destPath)
}

func CreateOrUpdate(alias string, pathToBin string, canOverrideExisting bool, dryRun bool) {
	validatePathToBin(pathToBin)

	var newPath = "user/local/bin/" + alias

	if utils.AliasExists(alias) {
		if canOverrideExisting == true {
			// Update alias for new bin
			var msg = "Do you want to override alias: " + alias

			if utils.ShouldOverrideFile(msg) == true {
				installToLocalBin(pathToBin, true, dryRun)
			} else {
				fmt.Printf("Overwrite of command %s with binary at %s has been aborted\n", alias, pathToBin)
			}
		} else {
			log.Fatal("Error: Alias already exists")
		}
	} else if canOverrideExisting == false {
		// Create new alias for bin
		fmt.Printf("Moving bin %s to %s for alias %s\n", pathToBin, newPath, alias)
		installToLocalBin(pathToBin, false, dryRun)
	} else {
		log.Fatal("Error: Alias doesn't exist")
	}
}

func Create(alias string, pathToBin string, dryRun bool) {
	CreateOrUpdate(alias, pathToBin, false, dryRun)
}

func Update(alias string, pathToBin string, dryRun bool) {
	CreateOrUpdate(alias, pathToBin, true, dryRun)
}

func validatePathToBin(pathToBin string) {
	_, err := os.Stat(pathToBin)
	if err == nil {
		return
	}
	fmt.Printf("Error: Failed to find bin at path: %s", pathToBin)
	os.Exit(1)
}
