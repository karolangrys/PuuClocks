package models_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"puuclocks/internal/models"
)

// CalledTimeCard(g *Game, c *Card)
var _ = Describe("RulesWhen", func() {
	var (
		game *models.Game
	)
	Context("test single rule", func() {
		Context("called time card", func() {
			BeforeEach(func() {
				lastCalledTime := 12.0

				game = &models.Game{LastCalledTime: &lastCalledTime}
			})

			It("should return true", func() {
				// given
				card := models.Card{Hour: 12.0}
				// when
				res := models.SameLastCalledTimeWhenRule(game, &card)
				// then
				Expect(res).To(BeTrue())
			})

			It("should return false", func() {
				// given
				card := models.Card{Hour: 10.0}
				// when
				res := models.SameLastCalledTimeWhenRule(game, &card)
				// then
				Expect(res).To(BeFalse())
			})
		})

		Context("wehicle card", func() {
			BeforeEach(func() {
				game = &models.Game{}
			})

			It("should return true", func() {
				// given
				card := models.Card{ClockID: 1}
				// when
				res := models.WehicleCardWhenRule(game, &card)
				// then
				Expect(res).To(BeTrue())
			})

			It("should return false", func() {
				// given
				card := models.Card{ClockID: 2}
				// when
				res := models.WehicleCardWhenRule(game, &card)
				// then
				Expect(res).To(BeFalse())
			})
		})
	})
})
