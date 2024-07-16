package actions

/* 
	Later move this to other place 
	It shouldnt be player action
	Create seperate channel for it in lobby
	And make listener for it
*/

/*
	// Server Actions
	ActionTypeEndOfTurn ActionType = "end-of-turn"

	// Server Actions Turn Related
	ActionTypeBeginReportTimeTurn   ActionType = "begin-report-time-turn"
	ActionTypeBegginActionTurn      ActionType = "begin-action-turn"
	ActionTypeBeginSynchronizedTurn ActionType = "begin-synchronization-turn"

	Pass as parameter broadcast to gameloop ^ for it to be sent for client
*/

/*
	if action.GetType() == actions.ActionTypeEndOfTurn{
		switch game.State {
		case models.GameStateReportTime:
			o.punishPlayers(game, []*models.Player{game.Players[game.Turn]})
		case models.GameStateSynchronization:
			var players []*models.Player
			for _, p := range game.Players {
				if _, ok := game.Synchronization[p.ConnectionID]; ok == false {
					players = append(players, p)
				}
			}

			o.punishPlayers(game, players)
		}
	}
*/