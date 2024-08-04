package service

import (
	"puuclocks/actions"
	"puuclocks/internal/models"

	"github.com/google/uuid"
)

type StartGameParams struct {
	Broadcast chan []byte
	OwnerID   uuid.UUID
	Game      *models.Game
	Action    actions.Action
}

type LobbyHandler interface {
	StartNewGame(params StartGameParams) (*models.Game, error)
}

type lobbyHandler struct {
}

func newLobbyHandler() LobbyHandler {
	return &lobbyHandler{}
}

func (l lobbyHandler) StartNewGame(params StartGameParams) (*models.Game, error) {
	if *params.Action.GetData().ReporterID != params.OwnerID {
		return nil, nil
	}

	if params.Game != nil {
		return nil, nil
	}

	newGame := models.NewGame()

	return &newGame, nil
}
