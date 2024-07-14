package actions

type ReportTime struct {
	action
}

func (r ReportTime) Validate(a action) *action {
	if a.Type != ActionTypeReportTime {
		return nil
	}

	if a.Data.ReportedTime == nil {
		return nil
	}

	parsedTime := *a.Data.ReportedTime

	if parsedTime > 12 || parsedTime < 0 {
		return nil
	}

	return &action{
		Type: ActionTypeReportTime,
		Data: ActionData{
			ReportedTime: &parsedTime,
		},
	}
}
