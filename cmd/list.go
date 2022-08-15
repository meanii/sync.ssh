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
	"github.com/alexeyco/simpletable"
	"github.com/meanii/sync.ssh/database"
	"github.com/spf13/cobra"
	"time"
)

/* listCmd represents the list command */
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list of all your syncing files!",
	Long: `This command shows all the syncing files and dir, which you have added!
try with -a, --all flag to show all list including deleted one!`,
	Run: func(cmd *cobra.Command, args []string) {
		_syncs := database.Database{}
		syncs, _ := _syncs.Find()

		/* starting table array */
		var cells [][]*simpletable.Cell

		/* init table instance */
		table := simpletable.New()

		/* setting up table header */
		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: "ID"},
				{Align: simpletable.AlignCenter, Text: "Target"},
				{Align: simpletable.AlignCenter, Text: "Type"},
				{Align: simpletable.AlignCenter, Text: "Status"},
				{Align: simpletable.AlignCenter, Text: "CreatedAt"},
			},
		}

		/* appending items in the table array */
		for _, s := range syncs {
			cells = append(cells, *&[]*simpletable.Cell{
				{Text: s.Id},
				{Text: s.Target},
				{Text: s.Type},
				{Text: s.Status},
				{Text: s.CreatedAt.Format(time.RFC822)},
			})
		}
		table.Body = &simpletable.Body{Cells: cells}
		table.Println()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("add", "a", false, "add flag for some extra information, it shows shows deleted lists too!")
}
