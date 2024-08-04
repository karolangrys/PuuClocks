package service

type Service interface {
	ActionExecutor() ActionExecutor
	LobbyHandler() LobbyHandler
	Validator() Validator
	FoulChecker() FoulChecker
	OutcomeEvaluator() OutcomeEvaluator

	Health() error
}

type service struct {
	lobbyHandler     LobbyHandler
	actionExecutor   ActionExecutor
	foulChecker      FoulChecker
	outcomeEvaluator OutcomeEvaluator
	validator        Validator
}

func NewService() Service {
	return &service{
		lobbyHandler:     newLobbyHandler(),
		actionExecutor:   newActionExecuter(),
		foulChecker:      newFoulChecker(),
		outcomeEvaluator: newOutcomeEvaluator(),
		validator:        newValidator(),
	}
}

func (s *service) LobbyHandler() LobbyHandler {
	return s.lobbyHandler
}

func (s *service) ActionExecutor() ActionExecutor {
	return s.actionExecutor
}

func (s *service) Validator() Validator {
	return s.validator
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
