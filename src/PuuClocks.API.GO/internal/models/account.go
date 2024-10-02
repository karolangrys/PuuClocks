package models

import "time"

type Accout struct {
	ID        int64
	Email     string
	Nickname  string
	Hash      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
