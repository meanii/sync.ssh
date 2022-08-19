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

package model

import "time"

type User struct {
	Id           string    `json:"id"`
	Owner        string    `json:"owner"`
	Cronjob      int       `json:"cronjob"`
	CreatedAt    time.Time `json:"createdAt"`
	TotalSync    int       `json:"totalSync"`
	Auth         bool      `json:"auth"`
	Token        string    `json:"token"`
	Repo         string    `json:"repo"`
	Github       string    `json:"github"`
	Branch 		 string    `json:"branch"`
	EmailAddress string    `json:"emailAddress"`
	Name         string    `json:"name"`
	Health       string    `json:"health"`
	UserDBPath   string    `json:"userDBPath"`
	SyncDBPath   string    `json:"syncDBPath"`
	SymlinkPath  string    `json:"symlinkPath"`
}
