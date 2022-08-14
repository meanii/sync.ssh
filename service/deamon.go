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
	"fmt"
	"github.com/meanii/sync.ssh/database"
)

func Deamon() {
	_database := database.Database{}
	_ = _database.Load()
	sync, _ := _database.Find()
	for _, s := range sync {
		fmt.Println(s.Target)
		fmt.Println(s.Type)
	}
}
