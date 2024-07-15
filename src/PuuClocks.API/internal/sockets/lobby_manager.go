package sockets

import (
	"puuclocks/internal/service"

	"github.com/google/uuid"
)

type LobbyManager interface {
	CreateLobby(services service.Service) Lobby
	FindLobby(uuid.UUID) Lobby
}

type lobbyManager struct {
	Lobbies map[uuid.UUID]Lobby
}

func NewLobbyManager() LobbyManager {
	return &lobbyManager{
		Lobbies: make(map[uuid.UUID]Lobby),
	}
}

func (l lobbyManager) CreateLobby(services service.Service) Lobby {
	lobby := NewLobby(services)
	id := lobby.GetID()
	l.Lobbies[id] = lobby

	return lobby
}

func (l lobbyManager) FindLobby(id uuid.UUID) Lobby {
	return l.Lobbies[id]
}
