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
	"github.com/meanii/sync.ssh/config"
	"github.com/meanii/sync.ssh/database"
	"github.com/meanii/sync.ssh/utils"
	"github.com/spf13/cobra"
	"log"
	osuser "os/user"
	"time"
)

/* initCmd represents the init command */
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "this init command, for init all processes which needed to start and init!",
	Long:  `this command for init all process, which is needed to start!`,
	Run: func(cmd *cobra.Command, args []string) {
		repo, _ := cmd.Flags().GetString("repo")

		user := database.User{}
		_ = user.Load()

		/* getting the owner  */
		currentUser, err := osuser.Current()
		if err != nil {
			log.Fatalf(err.Error())
		}

		user.Id = uuid.New().String()
		user.Owner = currentUser.Username
		user.CreatedAt = time.Now()
		user.Cronjob = 30
		user.Repo = repo

		/* checking, if user has entered a valid repo or not */
		utils.CheckRepo(repo)

		/* creating working dir */
		workingDir := config.GetWorkingDir()
		utils.CreateWorkingDir(workingDir)
		user.UserDBPath = config.GetUserDBPath()

		_ = user.Save(user)
		fmt.Println("Your sync.ssh has been established!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("repo", "r", "", "your github repositorie name, which your created for sync!")
}
