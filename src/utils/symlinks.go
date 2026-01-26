package utils

import "os"

func CreateSymlink(source, target string) error {
	return os.Symlink(source, target)
}

func IsSymlink(path string) bool {
	info, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return info.Mode()&os.ModeSymlink == os.ModeSymlink
}
