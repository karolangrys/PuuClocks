package concluder

import (
	"puuclocks/actions"
	"puuclocks/internal/models"
	"puuclocks/internal/service"

	"github.com/google/uuid"
)

type Concluder interface {
	ProcessAction(game *models.Game, socketID uuid.UUID, action actions.Action, broadcast chan []byte) (bool, error)
}

type concluder struct {
	validator        service.Validator
	actionExecutor   service.ActionExecutor
	foulChecker      service.FoulChecker
	outcomeEvaluator service.OutcomeEvaluator
}

func NewConcluder(s service.Service) Concluder {
	return &concluder{
		validator:        s.Validator(),
		actionExecutor:   s.ActionExecutor(),
		foulChecker:      s.FoulChecker(),
		outcomeEvaluator: s.OutcomeEvaluator(),
	}
}

func (c concluder) ProcessAction(game *models.Game, socketID uuid.UUID, action actions.Action, broadcast chan []byte) (bool, error) {
	canBePerformed, err := c.validator.ValidateAction(game, action)
	if err != nil {
		return true, err
	}

	if !canBePerformed {
		return false, nil
	}

	err = c.foulChecker.CheckForFaul(game, socketID, action)
	if err != nil {
		return true, err
	}

	err = c.actionExecutor.Execute(game, socketID, action)
	if err != nil {
		return true, err
	}

	c.outcomeEvaluator.ShouldPunishOrAward(game, socketID, action, broadcast)

	return c.shouldCloseGame(game, socketID, action)
}

func (c concluder) shouldCloseGame(_ *models.Game, _ uuid.UUID, _ actions.Action) (bool, error) {
	return false, nil
}
