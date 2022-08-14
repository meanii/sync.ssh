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

package validator

import (
	"github.com/meanii/sync.ssh/database"
	"github.com/meanii/sync.ssh/github"
	"log"
)

func CheckRepo(repo string) {
	user := database.User{}
	user.Load()
	ctx, client := github.Github(user.Token)
	_, _, err := client.Repositories.Get(ctx, user.Github, repo)
	if err != nil {
		log.Fatalf("Couldn't get the repo in your account! Reason: %v", err)
	}
}
