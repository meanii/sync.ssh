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
	"log"
	osuser "os/user"
	"time"

	"github.com/google/uuid"
	"github.com/meanii/sync.ssh/config"
	"github.com/meanii/sync.ssh/database"
	"github.com/meanii/sync.ssh/github"
	"github.com/meanii/sync.ssh/service"
	"github.com/meanii/sync.ssh/utils"
	"github.com/meanii/sync.ssh/validator"
	"github.com/spf13/cobra"
)

/* initCmd represents the init command */
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "this init command, for init all processes which needed to start and init!",
	Long:  `this command for init all process, which is needed to start!`,
	Run: func(cmd *cobra.Command, args []string) {
		repo, _ := cmd.Flags().GetString("repo")
		branch, _ := cmd.Flags().GetString("branch")

		user := database.User{}
		_ = user.Load()

		if len(user.Token) == 0 {
			log.Fatalf("You need to run 'sync.ssh auth --token=<TOKEN> --username=<GITHUB_USERNAME>' first, in order to init sync.ssh!")
		}

		/* getting the owner  */
		currentUser, err := osuser.Current()
		if err != nil {
			log.Fatalf(err.Error())
		}

		user.Id = uuid.New().String()
		user.Owner = currentUser.Username
		user.CreatedAt = time.Now()
		user.Branch = branch
		user.Repo = repo

		/* checking, if user has entered a valid repo or not */
		validator.CheckRepo(repo)

		/* creating working dir */
		workingDir := config.GetWorkingPath()
		utils.CreateWorkingDir(workingDir)
		user.UserDBPath = config.GetUserDBPath()

		/* creating symlink dir */
		symlinkPath := config.GetSymlinkPath()
		utils.CreateWorkingDir(symlinkPath)
		user.SymlinkPath = symlinkPath

		/* saving to database */
		_ = user.Save(user)

		/* creating README.md file while init */
		readme := utils.GetReadme()
		_github := github.GitService{}
		_github.Push(readme.FilePath, "")

		fmt.Println("Your sync.ssh has been established!")

		/* pushing db files */
		service.BackupDB()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("repo", "r", "", "your github repository name, which your created for sync!")
	initCmd.Flags().StringP("branch", "b", "main", "your github repository branch, which you want to sync!")
}
