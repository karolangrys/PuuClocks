package models

var ThenRules = map[int]func(*Game){
	1: SynchronizationThenRule,
	2: ReverseDirectionThenRule,
}

func SynchronizationThenRule(g *Game) {
	g.ExpectedSynchronization = true
}

func ReverseDirectionThenRule(g *Game) {
	g.Direction = !g.Direction
}
