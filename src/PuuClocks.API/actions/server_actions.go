package actions

import (
	"encoding/json"
	"puuclocks/internal/models"
)

type ServerSocketEvent string

var (
	ServerSocketEventGameStarting         ServerSocketEvent = "game-starting"
	ServerSocketEventLobbyOwner           ServerSocketEvent = "lobby-owner"
	ServerSocketEventNewPlayer            ServerSocketEvent = "new-player"
	ServerSocketEventPlayerConnected      ServerSocketEvent = "player-connected"
	ServerSocketEventPlayerDisconnected   ServerSocketEvent = "player-disconnected"
	ServerSocketEventCurrentPlayers       ServerSocketEvent = "current-players"
	ServerSocketEventUserMadeAction       ServerSocketEvent = "user-action"
	ServerSocketEventAvailableUserActions ServerSocketEvent = "available-user-actions"
)

type ServerSocketEventData struct {
	ConnectedPlayerNickname    *string                              `json:",omitempty"`
	DisconnectedPlayerNickname *string                              `json:",omitempty"`
	CurrentPlayers             []string                             `json:",omitempty"`
	ActionMade                 *ServerSocketEventDataUserMadeAction `json:",omitempty"`
	AvailableActions           []ActionType                         `json:",omitempty"`
}

type ServerSocketEventDataUserMadeAction struct {
	Nickname   string
	ActionType ActionType
	Data       ActionData
}

type ServerSocketEventMessage struct {
	Event ServerSocketEvent     `json:"Event"`
	Data  ServerSocketEventData `json:"Data"`
}

func ServerSocketEventMessageStartGame() []byte {
	message := ServerSocketEventMessage{
		Event: ServerSocketEventGameStarting,
	}

	r, err := json.Marshal(message)
	if err != nil {
		return []byte("")
	}

	return r
}

func ServerSocketEventMessageLobbyOwner(nickname string) []byte {
	message := ServerSocketEventMessage{
		Event: ServerSocketEventLobbyOwner,
		Data: ServerSocketEventData{
			ConnectedPlayerNickname: &nickname,
		},
	}

	r, err := json.Marshal(message)
	if err != nil {
		return []byte("")
	}

	return r
}

func ServerSocketEventMessagePlayerConnected(nickname string) []byte {
	message := ServerSocketEventMessage{
		Event: ServerSocketEventPlayerConnected,
		Data: ServerSocketEventData{
			ConnectedPlayerNickname: &nickname,
		},
	}

	r, err := json.Marshal(message)
	if err != nil {
		return []byte("")
	}

	return r
}

func ServerSocketEventMessagePlayerDisconnected(nickname string) []byte {
	message := ServerSocketEventMessage{
		Event: ServerSocketEventPlayerDisconnected,
		Data: ServerSocketEventData{
			DisconnectedPlayerNickname: &nickname,
		},
	}

	r, err := json.Marshal(message)
	if err != nil {
		return []byte("")
	}

	return r
}

func ServerSocketEventMessageCurrentPlayers(currentPlayerNicknames []string) []byte {
	message := ServerSocketEventMessage{
		Event: ServerSocketEventCurrentPlayers,
		Data: ServerSocketEventData{
			CurrentPlayers: currentPlayerNicknames,
		},
	}

	r, err := json.Marshal(message)
	if err != nil {
		return []byte("")
	}

	return r
}

func ServerSocketEventMessageUserMadeAction(action Action, userNickname string) []byte {
	message := ServerSocketEventMessage{
		Event: ServerSocketEventUserMadeAction,
		Data: ServerSocketEventData{
			ActionMade: &ServerSocketEventDataUserMadeAction{
				Nickname:   userNickname,
				ActionType: action.GetType(),
				Data:       action.GetData(),
			},
		},
	}

	r, err := json.Marshal(message)
	if err != nil {
		return []byte("")
	}

	return r
}

func ServerSocketEventMessageAvailableUserActions(gameState models.GameState) []byte {
	var availableActions []ActionType

	switch gameState {
	case models.GameStateReportTime:
		availableActions = []ActionType{ActionTypeReportTime, ActionTypeReportError}
	case models.GameStateSynchronization, models.GameStateAction:
		availableActions = []ActionType{ActionTypeReportError, ActionTypeSynchronization}
	}

	message := ServerSocketEventMessage{
		Event: ServerSocketEventAvailableUserActions,
		Data: ServerSocketEventData{
			AvailableActions: availableActions,
		},
	}

	r, err := json.Marshal(message)
	if err != nil {
		return []byte("")
	}

	return r
}

/*
	Later move this to other place
	It shouldnt be player action
	Create separate channel for it in lobby
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
