package service

import (
	"puuclocks/internal/repository"
)

type Service interface {
	Gameplay() Gameplay
	Validator() Validator
	ActionExecutor() ActionExecutor
	FoulChecker() FoulChecker
	OutcomeEvaluator() OutcomeEvaluator

	Health() error
}

type service struct {
	gameplay         Gameplay
	validator        Validator
	actionExecutor   ActionExecutor
	foulChecker      FoulChecker
	outcomeEvaluator OutcomeEvaluator
}

func NewService(databases repository.Databases) Service {
	validator := newValidate()
	outcomeEvaluator := newOutcomeEvaluator()
	foulChecker := newFoulChecker(databases.RedisDB())
	actionExecutor := newActionExecuter(databases.RedisDB())

	return &service{
		actionExecutor:   actionExecutor,
		foulChecker:      foulChecker,
		validator:        validator,
		outcomeEvaluator: outcomeEvaluator,
		gameplay: newGameplay(gamePlayServices{
			validator:        validator,
			actionExecutor:   actionExecutor,
			foulChecker:      foulChecker,
			outcomeEvaluator: outcomeEvaluator,
		}),
	}
}

func (s *service) Gameplay() Gameplay {
	return s.gameplay
}

func (s *service) Validator() Validator {
	return s.validator
}

func (s *service) ActionExecutor() ActionExecutor {
	return s.actionExecutor
}

func (s *service) FoulChecker() FoulChecker {
	return s.foulChecker
}

func (s *service) OutcomeEvaluator() OutcomeEvaluator {
	return s.outcomeEvaluator
}

func (s *service) Health() error {
	return nil
}
