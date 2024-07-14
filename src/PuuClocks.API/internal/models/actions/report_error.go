package actions

type ReportError struct {
	action
}

func (r ReportError) Validate() *action {
	return &action{
		Type: ActionTypeReportError,
	}
}