package utils

import "os"

func IsDir(file os.FileInfo) string {
	if file.IsDir() {
		return "dir"
	}
	return "file"
}
