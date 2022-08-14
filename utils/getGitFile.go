/*
Copyright © 2022 Anil Chauhan <https://github.com/meanii>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package utils

import (
	"log"
	"os"
)

type GitFile struct {
	Content  string
	FilePath string
	FileName string
}

var GitFiles []GitFile

func getFileData(filePath string) GitFile {
	data, err := os.ReadFile(filePath)
	fileInfo := GetFileInfo(filePath)
	if err != nil {
		log.Fatalf("Something went wrong while getting content of the file! Reason: %v", err)
	}
	return GitFile{
		Content:  string(data),
		FilePath: GetFilePath(filePath),
		FileName: fileInfo.Name(),
	}
}

func getDirData(filePath string) []GitFile {
	files, err := os.ReadDir(filePath)
	if err != nil {
		log.Fatalf("Something went wrong while read %v dir!\n", filePath)
	}

	for _, file := range files {
		if !file.IsDir() {
			GitFiles = append(GitFiles, getFileData(filePath+"/"+file.Name()))
		} else {
			getDirData(filePath + "/" + file.Name())
		}
	}
	return GitFiles
}

func GetGitFiles(filePath string) []GitFile {
	fileInfo := GetFileInfo(filePath)

	if !fileInfo.IsDir() {
		/* handling for file */
		return append(GitFiles, getFileData(filePath))
	}
	return getDirData(filePath)
}
