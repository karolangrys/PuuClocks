package actions

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
