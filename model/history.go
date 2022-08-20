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

type History struct {
	Id             string    `json:"id"`
	SyncId         string    `json:"syncId"`
	Target         string    `json:"target"`
	GitRootPath    string    `json:"gitRootPath"`
	SymlinkAddress string    `json:"symlink_address"`
	Type           string    `json:"fileType"`
	Status         string    `json:"status"`
	Owner          string    `json:"owner"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Repo           string    `json:"repo"`
	GithubUsername string    `json:"githubUsername"`
	RemarkMessage  string    `json:"remarkMessage"`
	Action         string    `json:"action"`
}
