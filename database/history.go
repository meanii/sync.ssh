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

type Histories []model.History

var historyPath = config.GetHistoryPath()

func (h *Histories) Find() Histories {
	_history, err := os.ReadFile(historyPath)
	var history Histories
	if err != nil {
		log.Fatalf("Something went while loading history db! Reason %v", err)
	}
	err = json.Unmarshal(_history, &history)
	if err != nil {
		log.Fatalf("Something went while unmarshalling history db! Reason %v", err)
	}
	return history
}

func (h *Histories) Save(history model.History) {
	*h = append(*h, history)
	data, err := json.Marshal(*h)
	if err != nil {
		log.Fatalf("something went wrong while marshaling data! Reason %v", err)
	}
	err = os.WriteFile(historyPath, data, 0644)
	if err != nil {
		log.Fatalf("something went wrong while writing data! Reason %v", err)
	}
}

func (h *Histories) FindByIdAndUpdate(Id string, history model.History) {
	var histories Histories
	_history := h.Find()
	for _, s := range _history {
		if s.Id == Id {
			s.Owner = history.Owner
			s.Type = history.Type
			s.Target = history.Target
			s.SymlinkAddress = history.SymlinkAddress
			s.Status = history.Status
			s.UpdatedAt = history.UpdatedAt
			s.RemarkMessage = history.RemarkMessage
		}
		histories = append(histories, s)
	}
	h.Store(histories)
}

func (h *Histories) Load() error {
	file, err := os.ReadFile(historyPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, h)
	if err != nil {
		return err
	}

	return nil
}

func (h *Histories) Store(histories Histories) {
	data, err := json.Marshal(histories)
	if err != nil {
		log.Fatal("something went wrong while marshaling data!")
	}
	err = os.WriteFile(historyPath, data, 0644)
	if err != nil {
		log.Fatal("something went wrong while writing data!")
	}
}
