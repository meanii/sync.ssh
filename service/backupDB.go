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
	"github.com/meanii/sync.ssh/config"
	"github.com/meanii/sync.ssh/github"
)

func BackupDB() {
	user := config.GetUserDBPath()
	sync := config.GetSyncDBPath()
	history := config.GetHistoryPath()

	_github := github.GitService{}

	/* pushing user db  */
	_github.Push(user, "")

	/* pushing sync db  */
	_github.Push(sync, "")

	/* pushing history db  */
	_github.Push(history, "")
}
