package service

import "puuclocks/internal/service/game"

type Service interface {
	GameLoop() game.GameLoop

	Health() error
}

type service struct {
	gameLoop game.GameLoop
}

func NewService() Service {
	game := game.NewGame()
	
	return &service{
		gameLoop: game.GameLoop(),
	}
}

func (s *service) GameLoop() game.GameLoop {
	return s.gameLoop
}

func (s *service) Health() error {
	return nil
}