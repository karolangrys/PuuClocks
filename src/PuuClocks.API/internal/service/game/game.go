package game

type Game interface {
	GameLoop() GameLoop
}

type game struct {
	gameLoop GameLoop
}

func NewGame() Game {
	return &game{
		gameLoop: newGameLoop(),
	}
}

func (s *game) GameLoop() GameLoop {
	return s.gameLoop
}
