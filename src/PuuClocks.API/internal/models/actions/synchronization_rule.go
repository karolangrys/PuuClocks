package actions

type SynchronizationRule struct {
	action
}

func (s SynchronizationRule) Validate(data string) *action {
	if data != string(ActionTypeSynchronizationRule) {
		return nil
	}

	return &action{
		Type: ActionTypeSynchronizationRule,
	}
}
