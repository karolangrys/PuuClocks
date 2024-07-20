package game

import (
	"fmt"
	"puuclocks/internal/models"
	"puuclocks/internal/models/actions"

	"github.com/google/uuid"
)

type FoulChecker interface {
	CheckForFaul(game *models.Game, socketID uuid.UUID, action actions.Action) error
}

type foulChecker struct{}

func newFoulChecker() FoulChecker {
	return &foulChecker{}
}

func (c foulChecker) CheckForFaul(game *models.Game, socketID uuid.UUID, action actions.Action) error {
	if action.GetType() == actions.ActionTypeReportError {
		return nil
	}

	var player *models.Player
	for _, p := range game.Players {
		if p != nil && p.ConnectionID == socketID {
			player = p
		}
	}

	if player == nil {
		return fmt.Errorf("couldn't obtain player with %d connection to determine who did action", socketID)
	}

	game.LastActionCaller = player

	if game.AreRulesBroken {
		return nil
	}

	game.LastActionCaller = player

	switch action.GetType() {
	case actions.ActionTypeSynchronization:
		if game.ExpectedSynchronization {
			return nil
		}
	case actions.ActionTypeReportTime:
		if game.Players[game.Turn] == player && *action.GetData().ReportedTime == game.ExpectedTime {
			return nil
		}
	}

	game.AreRulesBroken = true

	return nil
}
