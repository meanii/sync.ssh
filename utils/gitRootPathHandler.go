package utils

import "strings"

// handling gitRootPath for valid path
func GitRootPathHandler(path string) string {
	var validString string
	if strings.HasPrefix(path, "/") {
		validString = delChar(path, 0)
	}
	if !strings.HasSuffix(path, "/") {
		validString = path + "/"
	}
	return validString
}

// this function helps to remove a char from index wise
func delChar(_string string, index int) string {
	s := []rune(_string)
	return string(append(s[0:index], s[index+1:]...))
}
