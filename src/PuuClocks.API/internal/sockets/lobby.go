package sockets

import (
	"fmt"
	"puuclocks/internal/models"
	"puuclocks/internal/models/actions"
	"puuclocks/internal/service"
	"puuclocks/internal/service/game"

	"github.com/google/uuid"
)

type Lobby interface {
	GetID() uuid.UUID
	GetOwnerID() uuid.UUID
	GetPlayersNicknamesWithout(string) []string

	JoinLobby(Client)
	LeaveLobby(Client)

	ForwardMessage(Message)
}

type lobby struct {
	ID uuid.UUID

	Owner Client

	Join      chan Client
	Leave     chan Client
	Forward   chan Message
	Broadcast chan []byte

	Clients map[Client]bool

	Game         *models.Game
	Gameplay     game.GameLoop
	LobbyHandler service.LobbyHandler

	Settings Settings
}

type Settings struct{}

type Message struct {
	SocketID uuid.UUID
	Data     []byte
}

func NewLobby(services service.Service) Lobby {
	id := uuid.New()
	maxMessageAmount := 10

	l := lobby{
		ID: id,

		Forward:   make(chan Message, maxMessageAmount),
		Join:      make(chan Client),
		Leave:     make(chan Client),
		Clients:   make(map[Client]bool),
		Broadcast: make(chan []byte, maxMessageAmount),

		Gameplay:     services.GameLoop(),
		LobbyHandler: services.LobbyHandler(),
	}

	go l.run()

	return &l
}

func (l *lobby) run() {
	for {
		select {
		case client := <-l.Join:
			l.Clients[client] = true
		case client := <-l.Leave:
			delete(l.Clients, client)
			client.Close()
		case msg := <-l.Forward:
			fmt.Println("Action From: ", msg.SocketID, " Data: ", msg.Data)
			action := actions.ValidateUserProvidedAction(msg.Data)
			if action == nil {
				fmt.Println("User action couldn't be validated")
				break
			}

			if action.Data == nil {
				action.Data = &actions.ActionData{}
			}
			action.Data.ReporterID = &msg.SocketID

			actionRelated := actions.ActionRelatedTo(action.Type)
			if actionRelated == nil {
				fmt.Println("Action is not presigned to related type")
				break
			}

			switch *actionRelated {
			case actions.ActionRelatedGameplay:
				_, err := l.Gameplay.ProcessAction(l.Game, msg.SocketID, *action, l.Broadcast)
				if err != nil {
					fmt.Printf("Couldn't process action %v: %v", *action, err)
					break
				}
			case actions.ActionRelatedLobby:
				switch action.Type {
				case actions.ActionTypeStartGame:
					gameParams := service.StartGameParams{
						Broadcast: l.Broadcast,
						OwnerID:   l.GetOwnerID(),
						Game:      l.Game,
						Action:    action,
					}
					g, err := l.LobbyHandler.StartNewGame(gameParams)
					if err != nil {
						fmt.Printf("Couldn't starting game %v: %v", *action, err)
						break
					}

					if g == nil {
						break
					}

					l.Game = g
					l.Broadcast <- actions.ServerSocketEventMessageStartGame()
				}
			}

		case msg := <-l.Broadcast:
			for c := range l.Clients {
				c.SendMessage(msg)
			}
		}
	}
}

func (l *lobby) GetID() uuid.UUID {
	return l.ID
}

func (l *lobby) ForwardMessage(msg Message) {
	l.Forward <- msg
}

func (l *lobby) JoinLobby(c Client) {
	l.Join <- c

	if l.Owner == nil {
		l.Owner = c
		l.Broadcast <- actions.ServerSocketEventMessageLobbyOwner(c.GetNickname())
	} else {
		l.Broadcast <- actions.ServerSocketEventMessagePlayerConnected(c.GetNickname())
	}

	currentPlayers := l.GetPlayersNicknamesWithout(c.GetNickname())
	c.SendMessage(actions.ServerSocketEventMessageCurrentPlayers(currentPlayers))
}

func (l *lobby) LeaveLobby(c Client) {
	l.Leave <- c

	l.Broadcast <- actions.ServerSocketEventMessagePlayerDisconnected(c.GetNickname())
}

func (l *lobby) GetOwnerID() uuid.UUID {
	return l.Owner.GetID()
}

func (l *lobby) GetPlayersNicknamesWithout(nickname string) []string {
	var opponents []string

	for c := range l.Clients {
		n := c.GetNickname()
		if n != nickname {
			opponents = append(opponents, n)
		}
	}

	return opponents
}
