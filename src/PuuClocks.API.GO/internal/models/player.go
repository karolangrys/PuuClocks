package models

import "github.com/google/uuid"

type Player struct {
	AccountID    *int64
	ConnectionID uuid.UUID
	Nickname     string
	PlayingHand  []Card
}
