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
	"github.com/meanii/sync.ssh/service"
	"github.com/spf13/cobra"
)

/* deamonCmd represents the deamon command */
var deamonCmd = &cobra.Command{
	Use:   "deamon",
	Short: "this command, is use for deamon service!",
	Long:  `this command is use for background deamon service, which make sures your all files are sync!`,
	Run: func(cmd *cobra.Command, args []string) {
		service.Deamon()
	},
}

func init() {
	rootCmd.AddCommand(deamonCmd)
}
