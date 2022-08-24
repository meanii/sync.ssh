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
	"github.com/meanii/sync.ssh/service"
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
		gitPath, _ := cmd.Flags().GetString("path")

		/* initializing database and loading data */
		db := &database.Database{}
		history := database.Histories{}
		err := db.Load()
		if err != nil {
			log.Fatalf("something went wrong while laoding sync db! Reason %v", err)
		}
		err = history.Load()
		if err != nil {
			log.Fatalf("something went wrong while laoding hitory db! Reason %v", err)
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

		service.Cleaner(sync)

		/* create symlink */
		symlinkAddress := utils.CreateSymlink(target)

		/* inserting data to the db */
		_uuid := uuid.New().String()
		typ := utils.IsDir(fileInfo)
		createdAt := time.Now()
		owner := currentUser.Username
		GitRootPath := utils.GitRootPathHandler(gitPath)
		err = db.InsertOne(model.Sync{
			Id:             _uuid,
			Target:         target,
			Type:           typ,
			Status:         "active",
			CreatedAt:      createdAt,
			GitRootPath:    GitRootPath,
			Owner:          owner,
			SymlinkAddress: symlinkAddress,
		})
		if err != nil {
			fmt.Println("Something went wrong whiling inserting into the database!")
		}

		/* inserting in history */
		history.Save(model.History{
			Id:             uuid.New().String(),
			SyncId:         _uuid,
			Target:         target,
			GitRootPath:    GitRootPath,
			Type:           typ,
			Status:         "active",
			CreatedAt:      createdAt,
			UpdatedAt:      createdAt,
			Owner:          owner,
			SymlinkAddress: symlinkAddress,
			Action:         "created new sync",
			RemarkMessage:  "added to new sync!",
		})

		fmt.Printf("the %v '%v' has been added for sync!\n", utils.IsDir(fileInfo), fileInfo.Name())
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
	syncCmd.Flags().StringP("path", "p", "", "your custom git path, where you want to push the file!")
}
