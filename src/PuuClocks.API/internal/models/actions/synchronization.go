package actions

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
