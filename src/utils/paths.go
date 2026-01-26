package utils

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func GetBinDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	if runtime.GOOS == "windows" {
		return filepath.Join(home, "AppData", "local", "Microsoft", "WindowsApps")
	}
	return filepath.Join(home, ".local", "bin")
}
