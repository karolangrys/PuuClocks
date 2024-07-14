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
	ActionTypeReportError     ActionType = "report-error"
	ActionTypeReportTime      ActionType = "report-time"
	ActionTypeSynchronization ActionType = "synchronization-rule"
)

type ActionData struct {
	ReportedTime *float64 `json:"reportedTime"`
	ReporterID   *uuid.UUID `json:"ReportedID"`
}

type Action interface {
	GetType() ActionType
	GetData() ActionData
}

type action struct {
	Type ActionType `json:"type"`
	Data *ActionData `json:"data"`
}

/*
	{
		Type: "xyz",
		Data: {
			SocketID: <- Filled by server

			ReportedTime: float
		}
	}
*/
func (a action) GetType() ActionType {
	return a.Type
}

func (a action) GetData() ActionData {
	return *a.Data
}

func ValidateUserProvidedAction(b []byte) *action {
	var a action
	err := json.Unmarshal(b, &a)
	if err != nil {
		return nil
	}

	if a.Data != nil && a.Data.ReporterID != nil {
		return nil
	}

	if action := (ReportError{}).Validate(a); action != nil {
		return action
	}

	if action := (ReportTime{}).Validate(a); action != nil {
		return action
	}

	if action := (StartGame{}).Validate(a); action != nil {
		return action
	}

	if action := (Synchronization{}).Validate(a); action != nil {
		return action
	}

	return nil
}
