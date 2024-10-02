package service

import (
	"fmt"
	"puuclocks/actions"
	"puuclocks/internal/models"

	"slices"
)

type Validator interface {
	ValidateAction(game *models.Game, action actions.Action) (bool, error)
}

type validator struct{}

func newValidator() Validator {
	return &validator{}
}

func (v validator) ValidateAction(game *models.Game, action actions.Action) (bool, error) {
	switch game.State {
	case models.GameStateReportTime:
		if action.GetType() != actions.ActionTypeReportTime {
			return false, nil
		}

	case models.GameStateAction, models.GameStateSynchronization:
		allowedActions := []actions.ActionType{actions.ActionTypeReportError, actions.ActionTypeSynchronization}
		if !slices.Contains(allowedActions, action.GetType()) {
			return false, nil
		}
	default:
		return false, fmt.Errorf("unknown game state %d", game.State)
	}

	return true, nil
}
