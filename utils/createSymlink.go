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
	"github.com/meanii/sync.ssh/database"
	"log"
	"os"
)

func CreateSymlink(file string) {
	fileInfo := GetFileInfo(file)

	/* loading database */
	user := database.User{}
	_ = user.Load()

	/* handling symlink for file */
	target := user.SymlinkPath + "/" + fileInfo.Name()
	/* creating symlink file */
	err := os.Symlink(file, target)
	if err != nil {
		log.Fatalf("Something went wrong while creatign symlink! Reason: %v", err)
	}
}
