package service

import (
	"puuclocks/actions"
	"puuclocks/internal/common"
	"puuclocks/internal/models"

	"github.com/google/uuid"
)

type OutcomeEvaluator interface {
	ShouldPunishOrAward(game *models.Game, socketID uuid.UUID, action actions.Action, broadcast chan []byte)
}

type outcomeEvaluator struct{}

func newOutcomeEvaluator() OutcomeEvaluator {
	return &outcomeEvaluator{}
}

func (o outcomeEvaluator) ShouldPunishOrAward(game *models.Game, socketID uuid.UUID, action actions.Action, _ chan []byte) {
	if action.GetType() == actions.ActionTypeReportError {
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
}

func (o outcomeEvaluator) punishPlayers(game *models.Game, players []*models.Player) {
	chunkedCards := common.Chunk(game.DiscardedCards, len(players))
	for i, c := range chunkedCards {
		players[i].PlayingHand = append(players[i].PlayingHand, c...)
	}
	game.DiscardedCards = []models.Card{}
}
