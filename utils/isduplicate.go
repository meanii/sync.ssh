package utils

import "github.com/meanii/sync.ssh/model"

func IsDuplicate(sync []model.Sync, target string) bool {
	for _, s := range sync {
		if s.Target == target {
			return true
		}
	}
	return false
}
