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
			Direction: models.GameDirectionClockWise,
		}
		// when
		models.ReverseDirectionThenRule(&game)
		// then
		Expect(game.Direction).To(Equal(models.GameDirectionCounterClockWise))
	})
})
