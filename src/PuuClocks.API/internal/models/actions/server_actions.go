package actions

import "encoding/json"

type ServerSocketEvent string

var (
	ServerSocketEventGameStarting ServerSocketEvent = "game-starting"
)

type ServerSocketEventData struct {
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
