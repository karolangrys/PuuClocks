package actions

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
