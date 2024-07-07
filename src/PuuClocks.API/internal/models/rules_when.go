package models

var RulesWhen = map[int]func(*Game, *Card) bool{
	1: SameLastCalledTimeWhenRule,
	2: WehicleCardWhenRule,
}

func SameLastCalledTimeWhenRule(g *Game, c *Card) bool {
	return g.LastCalledTime != nil && *g.LastCalledTime == c.Hour
}

func WehicleCardWhenRule(g *Game, c *Card) bool {
	return c.ClockID == 1
}
