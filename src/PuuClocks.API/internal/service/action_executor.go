package service

import (
	"fmt"
	"puuclocks/internal/log"
	"puuclocks/internal/models"
	"puuclocks/internal/models/actions"
	"puuclocks/internal/repository"

	"github.com/google/uuid"
)

type ActionExecutor interface {
	Execute(game *models.Game, socketID uuid.UUID, action actions.Action) error
}

type actionExecutor struct {
	redis repository.Redis
}

func newActionExecuter(redis repository.Redis) ActionExecutor {
	return &actionExecutor{
		redis: redis,
	}
}

func (a actionExecutor) Execute(game *models.Game, socketID uuid.UUID, action actions.Action) error {
	switch action.GetType() {
	case actions.ActionTypeReportTime:
		var functionToApply func(*models.Game)
		var overload bool

		p := a.findPlayerBySocketID(game, socketID)
		if p == nil {
			return fmt.Errorf("couldn't find player who reported time with %d connection ID in %d game", socketID, game.ID)
		}

		game.LastCalledTime = action.GetData().ReportedTime

		drawedCard, err := a.drawCard(game, p)
		if err != nil {
			return err
		}

		functionToApply, overload, err = a.getRuleToApply(game, drawedCard)
		if err != nil {
			log.Log.Warn(err)
		}

		if overload {
			a.changeTurn(game, 1)
			a.changeTime(game, 1)
		} else if functionToApply != nil {
			functionToApply(game)
			if game.TurnDirection == models.GameDirectionClockWise {
				a.changeTurn(game, 1)
			} else {
				a.changeTurn(game, -1)
			}
		} else {
			if game.TurnDirection == models.GameDirectionClockWise {
				a.changeTurn(game, 1)
			} else {
				a.changeTurn(game, -1)
			}

			if game.TimeDirection == models.GameDirectionClockWise {
				a.changeTime(game, 1)
			} else {
				a.changeTime(game, -1)
			}
		}

		game.DiscardedCards = append(game.DiscardedCards, drawedCard)
		game.State = models.GameStateAction

	case actions.ActionTypeSynchronizationRule:

	}

	return nil
}

func (a actionExecutor) drawCard(game *models.Game, reporter *models.Player) (models.Card, error) {
	var card *models.Card
	var player *models.Player

	for _, p := range game.Players {
		if p == reporter {
			player = p
			break
		}
	}

	if len(player.PlayingHand) == 0 {
		return models.Card{}, fmt.Errorf("player with %d connection, dont have any cards", reporter.ConnectionID)
	}

	card, player.PlayingHand = &player.PlayingHand[0], player.PlayingHand[1:]

	return *card, nil
}

func (a actionExecutor) findPlayerBySocketID(game *models.Game, socketID uuid.UUID) *models.Player {
	for _, p := range game.Players {
		if p != nil && p.ConnectionID == socketID {
			return p
		}
	}

	return nil
}

func (a actionExecutor) getRuleToApply(game *models.Game, card models.Card) (func(*models.Game), bool, error) {
	var f func(*models.Game)
	occured := 0
	for _, rule := range game.Rules {
		doesOccure, err := rule.Occure(game, &card)
		if err != nil {
			return nil, false, err
		}

		if doesOccure {
			occured++
		}

		if occured == 1 {
			f, err = rule.RetrieveThen()
			if err != nil {
				return nil, false, err
			}
		}

		if occured > 1 {
			return nil, true, nil
		}
	}

	return f, false, nil
}

func (a actionExecutor) changeTurn(game *models.Game, howMany int) {
	var turn int

	turn += howMany
	if howMany < 0 {
		turn = len(game.Players) - 1
	}
	turn %= len(game.Players)

	game.Turn = turn
}

func (a actionExecutor) changeTime(game *models.Game, howMuch float64) {
	var exp float64

	exp = game.ExpectedTime + howMuch
	if exp > 12 {
		exp -= 12
	}
	if exp < 0 {
		exp += 12
	}

	game.ExpectedTime = exp
}
