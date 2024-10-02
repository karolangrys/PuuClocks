package service_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	"puuclocks/actions"
	"puuclocks/internal/models"
	"puuclocks/internal/repository"
	"puuclocks/internal/service"
)

var _ = Describe("Validator", Ordered, func() {
	var (
		databasesCtrl *gomock.Controller
		databasesMock *repository.MockDatabases

		reportErrorCtrl         *gomock.Controller
		reportTimeCtrl          *gomock.Controller
		startGameCtrl           *gomock.Controller
		synchronizationRuleCtrl *gomock.Controller

		reportError         *actions.MockAction
		reportTime          *actions.MockAction
		startGame           *actions.MockAction
		synchronizationRule *actions.MockAction

		validator service.Validator
	)

	BeforeEach(func() {
		databasesCtrl = gomock.NewController(GinkgoT())
		databasesMock = repository.NewMockDatabases(databasesCtrl)

		databasesMock.EXPECT().RedisDB().Return(nil).AnyTimes()

		reportErrorCtrl = gomock.NewController(GinkgoT())
		reportTimeCtrl = gomock.NewController(GinkgoT())
		startGameCtrl = gomock.NewController(GinkgoT())
		synchronizationRuleCtrl = gomock.NewController(GinkgoT())

		reportError = actions.NewMockAction(reportErrorCtrl)
		reportTime = actions.NewMockAction(reportTimeCtrl)
		startGame = actions.NewMockAction(startGameCtrl)
		synchronizationRule = actions.NewMockAction(synchronizationRuleCtrl)

		reportError.EXPECT().GetType().Return(actions.ActionTypeReportError).AnyTimes()
		reportTime.EXPECT().GetType().Return(actions.ActionTypeReportTime).AnyTimes()
		startGame.EXPECT().GetType().Return(actions.ActionTypeStartGame).AnyTimes()
		synchronizationRule.EXPECT().GetType().Return(actions.ActionTypeSynchronization).AnyTimes()

		validator = service.NewService().Validator()
	})

	Context("Game state equal to report time", func() {
		var (
			game models.Game
		)

		BeforeEach(func() {
			game = models.Game{
				State: models.GameStateReportTime,
			}
		})

		It("should return true for report time action", func() {
			// given && when
			allowed, err := validator.ValidateAction(&game, reportTime)
			// then
			Expect(allowed).To(Equal(true))
			Expect(err).To(BeNil())
		})

		It("should return false for not report time action", func() {
			actions := []actions.Action{
				reportError,
				startGame,
				synchronizationRule,
			}

			for _, a := range actions {
				// given && when
				allowed, err := validator.ValidateAction(&game, a)
				// then
				Expect(allowed).To(Equal(false))
				Expect(err).To(BeNil())
			}
		})
	})

	Context("Game state equal to synchronization or action", func() {
		Describe("synchronization game", func() {
			var (
				synchronizationGame models.Game
			)

			BeforeEach(func() {
				synchronizationGame = models.Game{State: models.GameStateSynchronization}
			})

			It("should return false", func() {
				testedActions := []actions.Action{
					startGame,
					reportTime,
				}

				for _, a := range testedActions {
					// when
					allowed, err := validator.ValidateAction(&synchronizationGame, a)
					// then
					Expect(allowed).To(Equal(false))
					Expect(err).To(BeNil())
				}
			})

			It("should return true", func() {
				// given
				actions := []actions.Action{
					synchronizationRule,
					reportError,
				}

				for _, a := range actions {
					// when
					allowed, err := validator.ValidateAction(&synchronizationGame, a)
					// then
					Expect(allowed).To(Equal(true))
					Expect(err).To(BeNil())
				}
			})
		})

		Describe("action game", func() {
			var (
				actionGame models.Game
			)

			BeforeEach(func() {
				actionGame = models.Game{State: models.GameStateAction}
			})

			It("should return false", func() {
				testedActions := []actions.Action{
					startGame,
					reportTime,
				}

				for _, a := range testedActions {
					// when
					allowed, err := validator.ValidateAction(&actionGame, a)
					// then
					Expect(allowed).To(Equal(false))
					Expect(err).To(BeNil())
				}
			})

			It("should return true", func() {
				// given
				actions := []actions.Action{
					synchronizationRule,
					reportError,
				}

				for _, a := range actions {
					// when
					allowed, err := validator.ValidateAction(&actionGame, a)
					// then
					Expect(allowed).To(Equal(true))
					Expect(err).To(BeNil())
				}
			})
		})
	})
})
