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

package service

import (
	"fmt"
	"github.com/meanii/sync.ssh/database"
	"github.com/meanii/sync.ssh/model"
	"os"
)

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// change status to deleted, if found deleted file
func Cleaner(sync []model.Sync) {
	for _, s := range sync {
		if s.Status != "deleted" {
			file, _ := exists(s.SymlinkAddress)
			if !file {
				fmt.Printf("found deleted file! %v\n", s.Target)
				_database := database.Database{}
				s.Status = "deleted"
				_database.FineByIdAndUpdate(s.Id, s)
			}
		}
	}
}
