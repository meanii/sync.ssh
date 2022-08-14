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

package database

import (
	"encoding/json"
	"errors"
	"github.com/meanii/sync.ssh/config"
	"github.com/meanii/sync.ssh/model"
	"log"
	"os"
)

var userPath = config.GetUserDBPath()

type User model.User

func (u *User) Save(user User) error {
	/* marshalling the object */
	data, err := json.Marshal(user)
	if err != nil {
		log.Fatal("something went wrong while marshaling data!")
	}

	/* writing the object to the db */
	err = os.WriteFile(userPath, data, 0644)
	if err != nil {
		log.Fatal("something went wrong, while writing to db!")
	}
	return nil
}

func (u *User) SaveToken(token string) {
	_ = u.Load()
	u.Token = token
	_ = u.Save(*u)
}

func (u *User) GetToken() string {
	_ = u.Load()
	return u.Token
}

func (u *User) Load() error {
	/* reading the user db file, and if not exist then return blank object */
	file, err := os.ReadFile(userPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	/* unmarshalling the object */
	err = json.Unmarshal(file, u)
	if err != nil {
		return err
	}

	return nil
}
