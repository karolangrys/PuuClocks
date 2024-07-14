package actions

//go:generate mockgen -source=action.go -destination=action_mock.go -package actions

import (
	"encoding/json"

	"github.com/google/uuid"
)

type ActionType string

var (
	ActionTypeStartGame ActionType = "start-game"

	// Gameplay related
	ActionTypeReportError         ActionType = "report-error"
	ActionTypeReportTime          ActionType = "report-time"
	ActionTypeSynchronizationRule ActionType = "synchronization-rule"
)

type ActionData struct {
	ReportedTime *float64
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

func ValidateIfUserProvidedActionInstance(b []byte) *action {
	var a action
    err := json.Unmarshal(b, &a)
    if err != nil {
        return nil
    }

	if action := (ReportTime{}).Validate(a); action != nil {
		return action
	}

	return nil
}