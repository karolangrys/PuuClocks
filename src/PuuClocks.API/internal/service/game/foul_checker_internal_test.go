package game

import (
	"fmt"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	"puuclocks/internal/models"
	"puuclocks/internal/models/actions"
	"puuclocks/internal/repository"
)

var _ = Describe("FoulChecker", func() {
	var (
		databasesCtrl *gomock.Controller
		databasesMock *repository.MockDatabases

		reportErrorCtrl         *gomock.Controller
		reportTimeCtrl          *gomock.Controller
		synchronizationRuleCtrl *gomock.Controller

		reportError         *actions.MockAction
		reportTime          *actions.MockAction
		synchronizationRule *actions.MockAction

		foulChecker FoulChecker
	)

	BeforeEach(func() {
		databasesCtrl = gomock.NewController(GinkgoT())
		databasesMock = repository.NewMockDatabases(databasesCtrl)

		databasesMock.EXPECT().RedisDB().Return(nil).AnyTimes()

		reportErrorCtrl = gomock.NewController(GinkgoT())
		reportTimeCtrl = gomock.NewController(GinkgoT())
		synchronizationRuleCtrl = gomock.NewController(GinkgoT())

		reportError = actions.NewMockAction(reportErrorCtrl)
		reportTime = actions.NewMockAction(reportTimeCtrl)
		synchronizationRule = actions.NewMockAction(synchronizationRuleCtrl)

		reportError.EXPECT().GetType().Return(actions.ActionTypeReportError).AnyTimes()
		reportTime.EXPECT().GetType().Return(actions.ActionTypeReportTime).AnyTimes()
		synchronizationRule.EXPECT().GetType().Return(actions.ActionTypeSynchronizationRule).AnyTimes()

		foulChecker = newFoulChecker()
	})

	Context("shouldn't execute instead return nil at once - ", func() {
		It("report error", func() {
			// given
			playerUUID := uuid.New()

			players := []*models.Player{
				{ConnectionID: uuid.Max},
				{ConnectionID: playerUUID},
			}

			game := models.Game{
				Players: players,
			}
			// when
			res := foulChecker.CheckForFaul(&game, uuid.Max, reportError)
			// then
			Expect(res).To(BeNil())
			Expect(game.LastActionCaller).To(BeNil())
			Expect(game.AreRulesBroken).To(Equal(false))
		})
	})

	Context("game rules already broken - change player who lastly called action", func() {
		It("should throw error if player not found by connection id", func() {
			// given
			players := []*models.Player{
				{ConnectionID: uuid.Max},
				{ConnectionID: uuid.Max},
			}

			game := models.Game{
				Players:        players,
				AreRulesBroken: true,
			}

			playerUUID := uuid.New()
			// when
			res := foulChecker.CheckForFaul(&game, playerUUID, reportTime)
			// then
			Expect(res.Error()).To(Equal(fmt.Sprintf("couldn't obtain player with %d connection to determine who did action", playerUUID)))
			Expect(game.LastActionCaller).To(BeNil())
			Expect(game.AreRulesBroken).To(Equal(true))
		})

		It("should change last calling player", func() {
			// given
			playerUUID := uuid.New()

			players := []*models.Player{
				{ConnectionID: uuid.Max},
				{ConnectionID: playerUUID},
			}

			game := models.Game{
				Players:        players,
				AreRulesBroken: true,
			}

			// when
			res := foulChecker.CheckForFaul(&game, playerUUID, reportTime)

			// then
			Expect(res).To(BeNil())
			Expect(game.LastActionCaller).To(Equal(players[1]))
			Expect(game.AreRulesBroken).To(Equal(true))
		})
	})

	Context("action called - synchornization rule", func() {
		It("player didn't break rules", func() {
			// given
			playerUUID := uuid.New()

			players := []*models.Player{
				{ConnectionID: uuid.Max},
				{ConnectionID: playerUUID},
			}

			game := models.Game{
				Players:                 players,
				AreRulesBroken:          false,
				ExpectedSynchronization: true,
			}

			// when
			res := foulChecker.CheckForFaul(&game, playerUUID, synchronizationRule)

			// then
			Expect(res).To(BeNil())
			Expect(game.LastActionCaller).To(Equal(players[1]))
			Expect(game.AreRulesBroken).To(Equal(false))
		})

		It("player break rules - synchronization not expected", func() {
			// given
			playerUUID := uuid.New()

			players := []*models.Player{
				{ConnectionID: uuid.Max},
				{ConnectionID: playerUUID},
			}

			game := models.Game{
				Players:                 players,
				AreRulesBroken:          false,
				ExpectedSynchronization: false,
			}

			// when
			res := foulChecker.CheckForFaul(&game, playerUUID, synchronizationRule)

			// then
			Expect(res).To(BeNil())
			Expect(game.LastActionCaller).To(Equal(players[1]))
			Expect(game.AreRulesBroken).To(Equal(true))
		})
	})

	Context("action called - report time", func() {
		It("player didn't break rules", func() {
			// given
			playerUUID := uuid.New()

			players := []*models.Player{
				{ConnectionID: uuid.Max},
				{ConnectionID: playerUUID},
			}

			expectedTime := 12.0

			game := models.Game{
				Players:        players,
				AreRulesBroken: false,
				Turn:           1,
				ExpectedTime:   expectedTime,
			}

			actionData := actions.ActionData{
				ReporterID:   &playerUUID,
				ReportedTime: &expectedTime,
			}

			reportTime.EXPECT().GetData().Return(actionData)

			// when
			res := foulChecker.CheckForFaul(&game, playerUUID, reportTime)

			// then
			Expect(res).To(BeNil())
			Expect(game.LastActionCaller).To(Equal(players[1]))
			Expect(game.AreRulesBroken).To(Equal(false))
		})

		It("player break rules - not his turn", func() {
			// given
			playerUUID := uuid.New()

			players := []*models.Player{
				{ConnectionID: uuid.Max},
				{ConnectionID: playerUUID},
			}

			expectedTime := 12.0

			game := models.Game{
				Players:        players,
				AreRulesBroken: false,
				Turn:           0,
				ExpectedTime:   expectedTime,
			}
			
			// when
			res := foulChecker.CheckForFaul(&game, playerUUID, reportTime)

			// then
			Expect(res).To(BeNil())
			Expect(game.LastActionCaller).To(Equal(players[1]))
			Expect(game.AreRulesBroken).To(Equal(true))
		})

		It("player break rules - wrong time called", func() {
			// given
			playerUUID := uuid.New()

			players := []*models.Player{
				{ConnectionID: uuid.Max},
				{ConnectionID: playerUUID},
			}

			expectedTime := 12.0
			reportedTime := 1.0

			game := models.Game{
				Players:        players,
				AreRulesBroken: false,
				Turn:           1,
				ExpectedTime:   expectedTime,
			}

			actionData := actions.ActionData{
				ReporterID:   &playerUUID,
				ReportedTime: &reportedTime,
			}

			reportTime.EXPECT().GetData().Return(actionData)

			// when
			res := foulChecker.CheckForFaul(&game, playerUUID, reportTime)

			// then
			Expect(res).To(BeNil())
			Expect(game.LastActionCaller).To(Equal(players[1]))
			Expect(game.AreRulesBroken).To(Equal(true))
		})
	})
})
