package models_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"puuclocks/internal/models"
)

var _ = Describe("RulesThen", func() {
	It("synchronization", func() {
		// given
		game := models.Game{}
		// when
		models.SynchronizationThenRule(&game)
		// then
		Expect(game.ExpectedSynchronization).To(BeTrue())
	})

	It("reverse direction", func() {
		// given
		game := models.Game{
			TimeDirection: models.GameDirectionClockWise,
		}
		// when
		models.ReverseTimeDirectionThenRule(&game)
		// then
		Expect(game.TimeDirection).To(Equal(models.GameDirectionCounterClockWise))
	})
})
