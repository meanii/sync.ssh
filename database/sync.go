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

var syncPath = config.GetSyncDBPath()

type Database []model.Sync

func (d *Database) InsertOne(sync model.Sync) error {
	*d = append(*d, sync)
	data, err := json.Marshal(*d)
	if err != nil {
		log.Fatal("something went wrong while marshaling data!")
		return err
	}
	err = os.WriteFile(syncPath, data, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}

	/* adding the working into the user db */
	user := User{}
	user.Load()
	user.SyncDBPath = syncPath
	user.Save(user)

	return nil
}

func (d *Database) FindByIdAndUpdate(Id string, sync model.Sync) {
	var database Database
	_database, _ := d.Find()
	for _, s := range _database {
		if s.Id == Id {
			s.Owner = sync.Owner
			s.Type = sync.Type
			s.Target = sync.Target
			s.Status = sync.Status
			s.SymlinkAddress = sync.SymlinkAddress
		}
		database = append(database, s)
	}
	d.Store(database)
}

func (d *Database) FindByIdAndDelete(Id string) {
	var database Database
	_database, _ := d.Find()
	for _, s := range _database {
		if s.Id != Id {
			database = append(database, s)
		}
	}
	d.Store(database)
}

func (d *Database) Find() (Database, error) {
	var sync Database
	data, err := os.ReadFile(syncPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
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
	file, err := os.ReadFile(syncPath)
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

func (d *Database) Store(database Database) {
	data, err := json.Marshal(database)
	if err != nil {
		log.Fatal("something went wrong while marshaling data!")
	}
	err = os.WriteFile(syncPath, data, 0644)
	if err != nil {
		log.Fatal("something went wrong while writing data!")
	}
}
