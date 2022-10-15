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
	"github.com/meanii/sync.ssh/database"
	"github.com/meanii/sync.ssh/validator"
	"github.com/spf13/cobra"
	"log"
)

/* authCmd represents the auth command */
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "sync.ssh allow to login with your github account with specific permissions and repo!",
	Long:  `this command  allow to login with your github account with specific permissions and repo to keep syncing your files/dir!`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		token, _ := cmd.Flags().GetString("token")

		if len(username) <= 0 {
			log.Fatalf("please provide github username!")
		}

		if len(token) <= 0 {
			log.Fatalf("please provide your github access token!")
		}

		user := database.User{}
		_ = user.Load() /* loading all the pre config */

		user.Token = token
		user.Github = username
		user.Branch = "main"
		user.Auth = true

		/* validating, if toke is valid or not */
		user.Name, user.EmailAddress = validator.CheckToken(token, username)
		_ = user.Save(user)

		fmt.Println("you have been logged in now!, please prefer to init --repo=<repo_name>")
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.Flags().StringP("username", "u", "", "your github username!")
	authCmd.Flags().StringP("token", "t", "", "your github token!")
}