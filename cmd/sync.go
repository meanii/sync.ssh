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

package cmd

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/meanii/sync.ssh/database"
	"github.com/meanii/sync.ssh/model"
	"github.com/meanii/sync.ssh/utils"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/user"
	"time"
)

/* syncCmd represents the sync command */
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "this command adds to sync the dir/file",
	Long:  `this command help to set the dir/file to backup!`,
	Run: func(cmd *cobra.Command, args []string) {

		/* validating if user has passed the dir or file path or not */
		if len(args) <= 0 {
			log.Fatal("Pass dir path OR file path, which you want to sync!")
			return
		}

		var target = args[0]

		/* initializing database and loading data */
		db := &database.Database{}
		err := db.Load()
		if err != nil {
			return
		}

		file, err := os.Open(target)
		if err != nil {
			log.Fatal(err)
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
			log.Fatal(err)
		}

		/* getting the owner  */
		currentUser, err := user.Current()
		if err != nil {
			log.Fatalf(err.Error())
		}

		/* checking if, the dir/file is already in sync or not */
		sync, _ := db.Find()
		if utils.IsDuplicate(sync, target) {
			log.Fatal("You have already syncing this dir/file!")
			return
		}

		/* inserting data to the db */
		err = db.InsertOne(model.Sync{
			Id:        uuid.New().String(),
			Target:    target,
			Type:      utils.IsDir(fileInfo),
			Status:    "active",
			CreatedAt: time.Now(),
			Owner:     currentUser.Username,
		})
		if err != nil {
			fmt.Println("Something went wrong whiling inserting into the database!")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
