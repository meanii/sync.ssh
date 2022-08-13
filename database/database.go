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
