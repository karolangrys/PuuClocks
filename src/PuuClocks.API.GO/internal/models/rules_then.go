package models

import "puuclocks/internal/consts"

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
	if exp > consts.MaxHour {
		exp -= consts.MaxHour
	}
	if exp < 0 {
		exp += consts.MaxHour
	}

	game.ExpectedTime = exp
}
