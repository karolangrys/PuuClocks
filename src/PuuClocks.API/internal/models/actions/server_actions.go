package actions

import "encoding/json"

type ServerSocketEvent string

var (
	ServerSocketEventGameStarting ServerSocketEvent = "game-starting"
	ServerSocketEventLobbyOwner   ServerSocketEvent = "lobby-owner"
)

type ServerSocketEventData struct {
	OwnerNickname *string
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
			OwnerNickname: &nickname,
		},
	}

	r, err := json.Marshal(message)
	if err != nil {
		return []byte("")
	}

	return r
}
