package model

import "time"

type Sync struct {
	Id        string    `json:"id"`
	Target    string    `json:"target"`
	Cronjob   int       `json:"cronjob"`
	Type      string    `json:"fileType"`
	CreatedAt time.Time `json:"createdAt"`
	Status    string    `json:"status"`
	Owner     string    `json:"owner"`
}
