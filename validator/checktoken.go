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
	"github.com/meanii/sync.ssh/github"
	"log"
)

func CheckToken(token string, username string) (string, string) {
	ctx, client := github.Github(token)
	user, _, err := client.Users.Get(ctx, username)
	if err != nil {
		log.Fatalf("failed to login with this token! Reason: %v", err)
	}
	return *user.Name, *user.Email
}
