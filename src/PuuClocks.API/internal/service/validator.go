package service

import (
	"fmt"
	"puuclocks/internal/models"
	"puuclocks/internal/models/actions"

	"slices"
)

type Validator interface {
	ValidateAction(game *models.Game, action actions.Action) (bool, error)
}

type validator struct{}

func newValidate() Validator {
	return &validator{}
}

func (v validator) ValidateAction(game *models.Game, action actions.Action) (bool, error) {
	if action.GetType() == actions.ActionTypeEndOfTurn {
		return true, nil
	}

	switch game.State {
	case models.GameStateReportTime:
		if action.GetType() != actions.ActionTypeReportTime {
			return false, nil
		}

	case models.GameStateAction, models.GameStateSynchronization:
		allowedActions := []actions.ActionType{actions.ActionTypeReportError, actions.ActionTypeSynchronizationRule}
		if !slices.Contains(allowedActions, action.GetType()) {
			return false, nil
		}
	default:
		return false, fmt.Errorf("unknown game state %d", game.State)
	}

	return true, nil
}
