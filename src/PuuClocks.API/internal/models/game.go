package models

import "github.com/google/uuid"

type Game struct {
	ID             uuid.UUID
	Rules          []Rule
	LastPlayedCard *Card
	DiscardedCards []Card

	AreRulesBroken bool
	Turn           int
	Direction      GameDirection

	LastCalledTime   *float64
	LastActionCaller *Player

	ExpectedTime            float64
	ExpectedSynchronization bool
	Synchronization         map[uuid.UUID]bool

	Players    []*Player
	State      GameState
	Scoreboard map[*Player]int
}

type GameState int

const (
	GameStateReportTime GameState = iota
	GameStateAction
	GameStateSynchronization
)

func (g GameState) ToString() string {
	switch g {
	case GameStateAction:
		return "Action"
	case GameStateReportTime:
		return "Report Time"
	case GameStateSynchronization:
		return "Synchronization"
	}
	return ""
}

type GameDirection bool

const (
	GameDirectionClockWise        GameDirection = true
	GameDirectionCounterClockWise GameDirection = false
)

func NewGame() Game {
	return Game{
		ID:    uuid.New(),
		Rules: DefaultRules(),
	}
}
