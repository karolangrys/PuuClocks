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

var _ = Describe("ActionExecutor", func() {
	var (
		databasesCtrl *gomock.Controller
		databasesMock *repository.MockDatabases

		reportTimeCtrl *gomock.Controller

		reportTime *actions.MockAction

		actionExecutor ActionExecutor
	)

	BeforeEach(func() {
		databasesCtrl = gomock.NewController(GinkgoT())
		databasesMock = repository.NewMockDatabases(databasesCtrl)

		databasesMock.EXPECT().RedisDB().Return(nil).AnyTimes()

		reportTimeCtrl = gomock.NewController(GinkgoT())

		reportTime = actions.NewMockAction(reportTimeCtrl)

		reportTime.EXPECT().GetType().Return(actions.ActionTypeReportTime).AnyTimes()

		actionExecutor = newActionExecuter()
	})

	Context("Action report time", func() {
		It("should throw error if player not found", func() {
			// given
			playerUUID := uuid.New()

			players := []*models.Player{
				{ConnectionID: uuid.Max},
			}

			game := models.Game{
				Players: players,
			}

			// when
			res := actionExecutor.Execute(&game, playerUUID, reportTime)

			// then
			Expect(res.Error()).To(Equal(fmt.Sprintf("couldn't find player who reported time with %d connection ID in %d game", playerUUID, game.ID)))
		})

		It("should throw error if player have no card at hand", func() {
			// given
			playerUUID := uuid.New()

			players := []*models.Player{
				{ConnectionID: playerUUID, PlayingHand: []models.Card{}},
			}

			game := models.Game{
				Players: players,
			}

			playerCalledTime := 12.0
			playerData := actions.ActionData{
				ReportedTime: &playerCalledTime,
			}

			reportTime.EXPECT().GetData().Return(playerData)

			// when
			res := actionExecutor.Execute(&game, playerUUID, reportTime)

			// then
			Expect(res.Error()).To(Equal(fmt.Sprintf("player with %d connection, dont have any cards", playerUUID)))
		})

		It("should apply default rule", func() {
			// given
			playerUUID := uuid.New()

			players := []*models.Player{
				{
					ConnectionID: playerUUID,
					PlayingHand: []models.Card{
						{ID: 1, Hour: 10.0, ClockID: 3},
						{ID: 2, Hour: 8.0, ClockID: 4},
					},
				},
				{
					ConnectionID: uuid.Max,
					PlayingHand: []models.Card{
						{ID: 1, Hour: 12.0, ClockID: 0},
					},
				},
				{
					ConnectionID: uuid.Max,
					PlayingHand: []models.Card{
						{ID: 1, Hour: 12.0, ClockID: 0},
					},
				},
			}

			game := models.Game{
				Players:      players,
				Turn:         0,
				TurnDirection: models.GameDirectionClockWise,
				TimeDirection: models.GameDirectionClockWise,
				Rules: models.DefaultRules(),
				ExpectedTime: 1.00,
			}

			playerCalledTime := 12.0
			playerData := actions.ActionData{
				ReportedTime: &playerCalledTime,
			}

			reportTime.EXPECT().GetData().Return(playerData)

			// when
			res := actionExecutor.Execute(&game, playerUUID, reportTime)

			// then
			Expect(res).To(BeNil())
			Expect(game.Turn).To(Equal(1))
			Expect(game.ExpectedTime).To(Equal(2.0))
			Expect(game.TurnDirection).To(Equal(models.GameDirectionClockWise))
			Expect(game.TimeDirection).To(Equal(models.GameDirectionClockWise))
			Expect(*game.LastCalledTime).To(Equal(playerCalledTime))
			Expect(game.ExpectedSynchronization).To(BeFalse())
		})

		It("should apply custom rule", func() {
			// given
			playerUUID := uuid.New()

			players := []*models.Player{
				{
					ConnectionID: playerUUID,
					PlayingHand: []models.Card{
						{ID: 1, Hour: 12.0, ClockID: 1},
						{ID: 2, Hour: 8.0, ClockID: 4},
					},
				},
				{
					ConnectionID: uuid.Max,
					PlayingHand: []models.Card{
						{ID: 1, Hour: 12.0, ClockID: 3},
					},
				},
				{
					ConnectionID: uuid.Max,
					PlayingHand: []models.Card{
						{ID: 1, Hour: 12.0, ClockID: 0},
					},
				},
			}

			game := models.Game{
				Players:      players,
				Turn:         0,
				TurnDirection: models.GameDirectionClockWise,
				TimeDirection: models.GameDirectionClockWise,
				Rules: models.DefaultRules(),
				ExpectedTime: 1.00,
			}

			playerCalledTime := 10.0
			playerData := actions.ActionData{
				ReportedTime: &playerCalledTime,
			}

			reportTime.EXPECT().GetData().Return(playerData)

			// when
			res := actionExecutor.Execute(&game, playerUUID, reportTime)

			// then
			Expect(res).To(BeNil())
			Expect(game.Turn).To(Equal(1))
			Expect(game.ExpectedTime).To(Equal(0.0))
			Expect(game.TurnDirection).To(Equal(models.GameDirectionClockWise))
			Expect(game.TimeDirection).To(Equal(models.GameDirectionCounterClockWise))
			Expect(*game.LastCalledTime).To(Equal(playerCalledTime))
			Expect(game.ExpectedSynchronization).To(BeFalse())
		})

		It("should apply overload - expected synchornization", func() {
			// given
			playerUUID := uuid.New()

			players := []*models.Player{
				{
					ConnectionID: playerUUID,
					PlayingHand: []models.Card{
						{ID: 1, Hour: 12.0, ClockID: 1},
						{ID: 2, Hour: 8.0, ClockID: 4},
					},
				},
				{
					ConnectionID: uuid.Max,
					PlayingHand: []models.Card{
						{ID: 1, Hour: 12.0, ClockID: 3},
					},
				},
				{
					ConnectionID: uuid.Max,
					PlayingHand: []models.Card{
						{ID: 1, Hour: 12.0, ClockID: 0},
					},
				},
			}
			
			lastCalledTime := float64(12.00)

			game := models.Game{
				Players:      players,
				Turn:         0,
				TurnDirection: models.GameDirectionClockWise,
				TimeDirection: models.GameDirectionClockWise,
				Rules: models.DefaultRules(),
				LastCalledTime: &lastCalledTime,
				ExpectedTime: 12.00,
			}

			playerCalledTime := 12.0
			playerData := actions.ActionData{
				ReportedTime: &playerCalledTime,
			}

			reportTime.EXPECT().GetData().Return(playerData)

			// when
			res := actionExecutor.Execute(&game, playerUUID, reportTime)

			// then
			Expect(res).To(BeNil())
			Expect(game.Turn).To(Equal(1))
			Expect(game.ExpectedTime).To(Equal(1.0))
			Expect(game.TurnDirection).To(Equal(models.GameDirectionClockWise))
			Expect(game.TimeDirection).To(Equal(models.GameDirectionClockWise))
			Expect(*game.LastCalledTime).To(Equal(playerCalledTime))
			Expect(game.ExpectedSynchronization).To(BeFalse())
		})
	})
})
