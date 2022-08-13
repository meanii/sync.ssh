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

	"github.com/spf13/cobra"
)

/* configCmd represents the config command */
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "config your sync.ssh settings",
	Long:  `this command helps us to configure sync.ssh settings such as like cronjob setting!`,
	Run: func(cmd *cobra.Command, args []string) {
		cronjob, _ := cmd.Flags().GetInt("cronjob")
		fmt.Println("config called", cronjob)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().IntP("cronjob", "c", 0, "cronjob timing, so which you gonna set up it will run the deamon service!")
}
