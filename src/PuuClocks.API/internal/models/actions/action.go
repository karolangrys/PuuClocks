package actions

//go:generate mockgen -source=action.go -destination=action_mock.go -package actions

import (
	"puuclocks/internal/models"

	"github.com/google/uuid"
)

type ActionType string

var (
	ActionTypeStartGame ActionType = "start-game"

	// Server Actions
	ActionTypeEndOfTurn ActionType = "end-of-turn"

	// Server Actions Turn Related
	ActionTypeBeginReportTimeTurn   ActionType = "begin-report-time-turn"
	ActionTypeBegginActionTurn      ActionType = "begin-action-turn"
	ActionTypeBeginSynchronizedTurn ActionType = "begin-synchronization-turn"

	// Gameplay related
	ActionTypeReportError         ActionType = "report-error"
	ActionTypeReportTime          ActionType = "report-time"
	ActionTypeSynchronizationRule ActionType = "synchronization-rule"
)

type ActionData struct {
	ReportedTime *float64
	State        *models.GameState
	ReporterID *uuid.UUID
}

type Action interface {
	GetType() ActionType
	GetData() ActionData
}

type action struct {
	Type ActionType
	Data ActionData
}

func (a action) GetType() ActionType {
	return a.Type
}

func (a action) GetData() ActionData {
	return a.Data
}

func ValidateIfUserProvidedActionInstance(data string) *action {
	if a := (ReportTime{}).Validate(data); a != nil {
		return a
	}

	if a := (ReportError{}).Validate(data); a != nil {
		return a
	}

	if a := (StartGame{}).Validate(data); a != nil {
		return a
	}

	if a := (SynchronizationRule{}).Validate(data); a != nil {
		return a
	}

	return nil
}
