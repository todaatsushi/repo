package repoconf

import (
	"os"
	"path/filepath"
)

func GetRootPath() (string, bool) {
	home, _ := os.UserHomeDir()
	exists := true

	if _, err := os.Stat(home); os.IsNotExist(err) {
		exists = false
	}
	return filepath.Join(home, ".config", "repo"), exists
}

func GetLibPath() (string, bool) {
	root, exists := GetRootPath()
	filename := filepath.Join(root, "lib.json")

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exists = false
	}
	return filename, exists
}
