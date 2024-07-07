package models

import "fmt"

type Rule struct {
	WhenID int
	ThenID int
}

func (r Rule) Occure(g *Game, c *Card) (bool, error) {
	rule, ok := RulesWhen[r.WhenID]
	if !ok {
		return false, fmt.Errorf("couldn't find when rule with %d id", r.WhenID)
	}

	return rule(g, c), nil
}

func (r Rule) Then(g *Game) error {
	then, ok := ThenRules[r.ThenID]
	if !ok {
		return fmt.Errorf("couldn't find then rule with %d id", r.ThenID)
	}

	then(g)

	return nil
}

func (r Rule) RetrieveThen() (func(*Game), error) {
	then, ok := ThenRules[r.ThenID]
	if !ok {
		return nil, fmt.Errorf("couldn't find then rule with %d id", r.ThenID)
	}

	return then, nil
}

func DefaultRules() []Rule {
	return []Rule{
		{
			WhenID: 1,
			ThenID: 1,
		},
		{
			WhenID: 2,
			ThenID: 2,
		},
	}
}
