package game

import (
	"puuclocks/internal/models"
	"puuclocks/internal/models/actions"

	"github.com/google/uuid"
)

type GameLoop interface {
	ProcessAction(game *models.Game, socketID uuid.UUID, action actions.Action, broadcast chan (string)) (bool, error)
}

type gameLoop struct {
	validator        Validator
	foulChecker      FoulChecker
	actionExecutor   ActionExecutor
	outcomeEvaluator OutcomeEvaluator
}

func newGameLoop() GameLoop {
	return &gameLoop{
		validator:        newValidator(),
		actionExecutor:   newActionExecuter(),
		foulChecker:      newFoulChecker(),
		outcomeEvaluator: newOutcomeEvaluator(),
	}
}

func (g gameLoop) ProcessAction(game *models.Game, socketID uuid.UUID, action actions.Action, broadcast chan (string)) (bool, error) {
	canBePerformed, err := g.validator.ValidateAction(game, action)
	if err != nil {
		return true, err
	}

	if !canBePerformed {
		return false, nil
	}

	err = g.foulChecker.CheckForFaul(game, socketID, action)
	if err != nil {
		return true, err
	}

	err = g.actionExecutor.Execute(game, socketID, action)
	if err != nil {
		return true, err
	}

	g.outcomeEvaluator.ShouldPunishOrAward(game, socketID, action, broadcast)

	return g.shouldCloseGame(game, socketID, action)
}

func (g gameLoop) shouldCloseGame(game *models.Game, socketID uuid.UUID, action actions.Action) (bool, error) {
	return false, nil
}
