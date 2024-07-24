package actions

import "encoding/json"

type ServerSocketEvent string

var (
	ServerSocketEventGameStarting       ServerSocketEvent = "game-starting"
	ServerSocketEventLobbyOwner         ServerSocketEvent = "lobby-owner"
	ServerSocketEventNewPlayer          ServerSocketEvent = "new-player"
	ServerSocketEventPlayerConnected    ServerSocketEvent = "player-connected"
	ServerSocketEventPlayerDisconnected ServerSocketEvent = "player-disconnected"
	ServerSocketEventCurrentPlayers     ServerSocketEvent = "current-players"
)

type ServerSocketEventData struct {
	ConnectedPlayerNickname    *string  `json:",omitempty"`
	DisconnectedPlayerNickname *string  `json:",omitempty"`
	CurrentPlayers             []string `json:",omitempty"`
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
