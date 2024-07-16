package service

import (
	"puuclocks/internal/models"
	"puuclocks/internal/models/actions"

	"github.com/google/uuid"
)

type LobbyHandler interface{
	StartNewGame(broadcast chan []byte, ownerID uuid.UUID, game *models.Game, action actions.Action) (*models.Game,error)
}

type lobbyHandler struct {
}

func newLobbyHandler() LobbyHandler {
	return &lobbyHandler{}
}

func (l lobbyHandler) StartNewGame(broadcast chan []byte, ownerID uuid.UUID, game *models.Game, action actions.Action) (*models.Game,error) {
	if *action.GetData().ReporterID != ownerID {
		return nil,nil
	}

	if game != nil {
		return nil,nil
	}

	newGame := models.NewGame()

	return &newGame, nil
}