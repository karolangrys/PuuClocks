package game

import (
	"puuclocks/internal/common"
	"puuclocks/internal/models"
	"puuclocks/internal/models/actions"

	"github.com/google/uuid"
)

type OutcomeEvaluator interface {
	ShouldPunishOrAward(game *models.Game, socketID uuid.UUID, action actions.Action, broadcast chan (string))
}

type outcomeEvaluator struct {
}

func newOutcomeEvaluator() OutcomeEvaluator {
	return &outcomeEvaluator{}
}


func (o outcomeEvaluator) ShouldPunishOrAward(game *models.Game, socketID uuid.UUID, action actions.Action, broadcast chan (string)) {
	if action.GetType() ==  actions.ActionTypeReportError{
		player := game.LastActionCaller
		if !game.AreRulesBroken {
			for _, p := range game.Players {
				if p.ConnectionID == socketID {
					player = p
				}
			}
		}
		o.punishPlayers(game, []*models.Player{player})		
	}

	if action.GetType() == actions.ActionTypeEndOfTurn{
		switch game.State {
		case models.GameStateReportTime:
			o.punishPlayers(game, []*models.Player{game.Players[game.Turn]})
		case models.GameStateSynchronization:
			var players []*models.Player
			for _, p := range game.Players {
				if _, ok := game.Synchronization[p.ConnectionID]; ok == false {
					players = append(players, p)
				}
			}

			o.punishPlayers(game, players)
		}
	}
}

func (o outcomeEvaluator) punishPlayers(game *models.Game, players []*models.Player) {
	chunkedCards := common.Chunk(game.DiscardedCards, len(players))
	for i, c := range chunkedCards {
		players[i].PlayingHand = append(players[i].PlayingHand, c...)
	}
	game.DiscardedCards = []models.Card{}
}