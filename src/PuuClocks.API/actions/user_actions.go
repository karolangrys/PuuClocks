package actions

// ----------------------------------- Report Error ----------------------------------- //

type ReportError struct {
	action
}

func (r ReportError) Validate(a action) *action {
	if a.Type != ActionTypeReportError || a.Data != nil {
		return nil
	}

	return &action{
		Type: ActionTypeReportError,
	}
}

// ----------------------------------- Report Time ----------------------------------- //

type ReportTime struct {
	action
}

func (r ReportTime) Validate(a action) *action {
	if a.Type != ActionTypeReportTime {
		return nil
	}

	if a.Data == nil || a.Data.ReportedTime == nil {
		return nil
	}

	parsedTime := *a.Data.ReportedTime

	if parsedTime > 12 || parsedTime <= 0 {
		return nil
	}

	return &action{
		Type: ActionTypeReportTime,
		Data: &ActionData{
			ReportedTime: &parsedTime,
		},
	}
}

// ----------------------------------- Start Game ----------------------------------- //

type StartGame struct {
	action
}

func (s StartGame) Validate(a action) *action {
	if a.Type != ActionTypeStartGame || a.Data != nil {
		return nil
	}

	return &action{
		Type: ActionTypeStartGame,
	}
}

// ----------------------------------- Synchronization ----------------------------------- //

type Synchronization struct {
	action
}

func (s Synchronization) Validate(a action) *action {
	if a.Type != ActionTypeSynchronization || a.Data != nil {
		return nil
	}

	return &action{
		Type: ActionTypeSynchronization,
	}
}
