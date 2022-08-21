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

package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

func GetUserDBPath() string {
	userDb := "/.user.json"
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("something went wrong while getting $HOME! Reason %v", err)
	}
	configPath := filepath.Join(dirname, ".meanii/sync/config")
	createWorkingDir(configPath, &userDb)
	return configPath + userDb
}

func GetSymlinkPath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("something went wrong while getting $HOME! Reason %v", err)
	}
	configPath := filepath.Join(dirname, ".meanii/sync/symlink")
	createWorkingDir(configPath, nil)
	return configPath
}

func GetSyncDBPath() string {
	syncDb := "/.sync.json"
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("something went wrong while getting $HOME! Reason %v", err)
	}
	configPath := filepath.Join(dirname, ".meanii/sync")
	createWorkingDir(configPath, &syncDb)
	return configPath + syncDb
}

func GetWorkingPath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("something went wrong while getting $HOME! Reason %v", err)
	}
	working := filepath.Join(dirname, ".meanii/sync")
	createWorkingDir(working, nil)
	return working
}

func GetHistoryPath() string {
	historyDbPath := "/.history.json"
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("something went wrong while getting $HOME! Reason %v", err)
	}
	history := filepath.Join(dirname, ".meanii/sync")
	createWorkingDir(history, &historyDbPath)
	return history + historyDbPath
}

// for creating nested dir and file fi not exists
func createWorkingDir(folderPath string, filename *string) {

	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Something went wrong while create dir! Reason %v", err)
	}
	/* handling for creating new file if weather if exist or not */
	if filename != nil {
		filePath := folderPath + *filename
		_, err = os.ReadFile(filePath)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				err = os.WriteFile(filePath, []byte{}, 0644)
				if err != nil {
					log.Fatal("something went wrong while writing data!")
				}
			}
		}
	}
}
