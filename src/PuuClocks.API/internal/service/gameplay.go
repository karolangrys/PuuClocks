package service

import (
	"puuclocks/internal/models"
	"puuclocks/internal/models/actions"

	"github.com/google/uuid"
)

type Gameplay interface {
	ProcessAction(game *models.Game, socketID uuid.UUID, action actions.Action, broadcast chan (string)) (bool, error)
}

type gameplay struct {
	validator        Validator
	foulChecker     FoulChecker
	actionExecutor   ActionExecutor
	outcomeEvaluator OutcomeEvaluator
}

type gamePlayServices struct {
	validator        Validator
	foulChecker     FoulChecker
	actionExecutor   ActionExecutor
	outcomeEvaluator OutcomeEvaluator
}

func newGameplay(services gamePlayServices) Gameplay {
	return &gameplay{
		validator:        services.validator,
		actionExecutor:   services.actionExecutor,
		foulChecker:     services.foulChecker,
		outcomeEvaluator: services.outcomeEvaluator,
	}
}

func (g gameplay) ProcessAction(game *models.Game, socketID uuid.UUID, action actions.Action, broadcast chan (string)) (bool, error) {
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

func (g gameplay) shouldCloseGame(game *models.Game, socketID uuid.UUID, action actions.Action) (bool, error) {
	return false, nil
}
