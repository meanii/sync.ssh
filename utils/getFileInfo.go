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

func GetFileInfo(path string) os.FileInfo {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Something went wrong while getting file info! Reason: %v", err)
	}

	/* closing the opened file, in the end of the process */
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Something went wrong, while closing the db!")
		}
	}(file)

	/* This returns an *os.FileInfo type */
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("something went wrong while getting stat! Reason %v", err)
	}

	return fileInfo
}
