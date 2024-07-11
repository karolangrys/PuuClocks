package models

var ThenRules = map[int]func(*Game){
	1: SynchronizationThenRule,
	2: ReverseTimeDirectionThenRule,
}

func SynchronizationThenRule(g *Game) {
	g.ExpectedSynchronization = true
	changeTime(g, 1)
}

func ReverseTimeDirectionThenRule(g *Game) {
	g.TimeDirection = !g.TimeDirection
	changeTime(g, -1)
}

func changeTime(game *Game, howMuch float64) {
	var exp float64

	exp = game.ExpectedTime + howMuch
	if exp > 12 {
		exp -= 12
	}
	if exp < 0 {
		exp += 12
	}

	game.ExpectedTime = exp
}