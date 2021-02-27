package utils

import (
	"os"
)

func FileExists(fileName string) bool {

	fileInfo, err := os.Lstat(fileName)

	if fileInfo != nil || (err != nil && !os.IsNotExist(err)) {
		return true
	}
	return false
}
