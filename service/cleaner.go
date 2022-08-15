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
	"github.com/google/uuid"
	"github.com/meanii/sync.ssh/config"
	"github.com/meanii/sync.ssh/database"
	"github.com/meanii/sync.ssh/model"
	"log"
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

// Cleaner change status to deleted, if found deleted file
func Cleaner(sync []model.Sync) {
	for _, s := range sync {
		file, _ := exists(s.SymlinkAddress)
		if !file {
			fmt.Printf("found deleted file! %v\n", s.Target)
			_database := database.Database{}
			err := os.RemoveAll(s.SymlinkAddress)
			if err != nil {
				log.Fatalf("Something went wrong while deleted junk file! %v\n", s.SymlinkAddress)
			}
			_database.FindByIdAndDelete(s.Id)
		}
	}

	/* cleaning the file, which isn't exists in the db! */
	symlinkPath := config.GetSymlinkPath()
	symlinkDirCleaner(symlinkPath, sync)
}

func symlinkDirCleaner(symlinkPath string, sync []model.Sync) {
	files, err := os.ReadDir(symlinkPath)
	historyDB := database.Histories{}
	_ = historyDB.Load()
	if err != nil {
		log.Fatalf("Something went wrong while read %v dir!\n", symlinkPath)
	}
	for _, file := range files {
		symPathItem := symlinkPath + "/" + file.Name()
		isValidSymlinkFile := isValidFile(symPathItem, sync)
		if !isValidSymlinkFile {
			fmt.Printf("found not valid file %v\n", symPathItem)
			err := os.RemoveAll(symPathItem)
			if err != nil {
				log.Fatalf("Something went wrong while deleted junk file! %v\n", symPathItem)
			}

			/* adding to history */
			historyDB.Save(model.History{
				Id:            uuid.New().String(),
				Action:        "deleted",
				RemarkMessage: fmt.Sprintf("found not a valid file while cleaing! File: %v", symlinkPath),
			})

		}
	}
}

/* check if the file is exits in the sync db or not */
func isValidFile(symPathItem string, sync []model.Sync) bool {
	for _, s := range sync {
		if symPathItem == s.SymlinkAddress {
			return true
		}
	}
	return false
}
