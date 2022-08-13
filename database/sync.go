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
	"github.com/meanii/sync.ssh/model"
	"log"
	"os"
)

var filename = `.sync.json`

type Database []model.Sync

func (d *Database) InsertOne(sync model.Sync) error {
	*d = append(*d, sync)
	data, err := json.Marshal(*d)
	if err != nil {
		log.Fatal("something went wrong while marshaling data!")
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (d *Database) Find() (Database, error) {
	var sync Database
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("something went wrong while reading database!")
		return nil, err
	}

	/* if found empty database then return blank array */
	if len(data) == 0 {
		return Database{}, nil
	}

	err = json.Unmarshal(data, &sync)
	if err != nil {
		log.Fatal("something went wrong while unmarshalling data!")
		return nil, err
	}
	return sync, nil
}

func (d *Database) Load() error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, d)
	if err != nil {
		return err
	}

	return nil
}
