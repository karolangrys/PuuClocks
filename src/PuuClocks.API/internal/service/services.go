package service

import "puuclocks/internal/service/game"

type Service interface {
	GameLoop() game.GameLoop
	LobbyHandler() LobbyHandler

	Health() error
}

type service struct {
	gameLoop game.GameLoop
	lobbyHandler LobbyHandler
}

func NewService() Service {
	game := game.NewGame()
	
	return &service{
		gameLoop: game.GameLoop(),
		lobbyHandler: newLobbyHandler(),
	}
}

func (s *service) GameLoop() game.GameLoop {
	return s.gameLoop
}

func (s *service) LobbyHandler() LobbyHandler {
	return s.lobbyHandler
}

func (s *service) Health() error {
	return nil
}