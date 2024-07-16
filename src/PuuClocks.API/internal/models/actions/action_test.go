package actions_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"puuclocks/internal/models/actions"
)

var _ = Describe("Action", Ordered, func() {
	Describe("ValidateUserProvidedAction", func() {
		Describe("Type StartGame", func() {
			It("should return nil due to containing data", func() {
				// given
				b := []byte(fmt.Sprintf(`{"type":%q,data:{}}`, actions.ActionTypeStartGame))

				// when
				action := actions.ValidateUserProvidedAction(b)

				// then
				Expect(action).To(BeNil())
			})

			It("should return action", func() {
				// given
				b := []byte(fmt.Sprintf(`{"type":%q}`, actions.ActionTypeStartGame))

				// when
				action := actions.ValidateUserProvidedAction(b)

				// then
				Expect(action).ToNot(BeNil())
				Expect(action.Type).To(Equal(actions.ActionTypeStartGame))
				Expect(action.Data).To(BeNil())
			})
		})

		Describe("Type ReportError", func() {
			It("should return nil due to containing data", func() {
				// given
				b := []byte(fmt.Sprintf(`{"type":%q,data:{}}`, actions.ActionTypeReportError))

				// when
				action := actions.ValidateUserProvidedAction(b)

				// then
				Expect(action).To(BeNil())
			})

			It("should return action", func() {
				// given
				b := []byte(fmt.Sprintf(`{"type":%q}`, actions.ActionTypeReportError))

				// when
				action := actions.ValidateUserProvidedAction(b)

				// then
				Expect(action).ToNot(BeNil())
				Expect(action.Type).To(Equal(actions.ActionTypeReportError))
				Expect(action.Data).To(BeNil())
			})
		})

		Describe("Type Synchronization", func() {
			It("should return nil due to containing data", func() {
				// given
				b := []byte(fmt.Sprintf(`{"type":%q,data:{}}`, actions.ActionTypeSynchronization))

				// when
				action := actions.ValidateUserProvidedAction(b)

				// then
				Expect(action).To(BeNil())
			})

			It("should return action", func() {
				// given
				b := []byte(fmt.Sprintf(`{"type":%q}`, actions.ActionTypeSynchronization))

				// when
				action := actions.ValidateUserProvidedAction(b)

				// then
				Expect(action).ToNot(BeNil())
				Expect(action.Type).To(Equal(actions.ActionTypeSynchronization))
				Expect(action.Data).To(BeNil())
			})
		})
		
		Describe("Type ReportTime", func() {
			It("should return nil due to missing data", func() {
				// given
				b := []byte(fmt.Sprintf(`{"type":%q}`, actions.ActionTypeReportTime))

				// when
				action := actions.ValidateUserProvidedAction(b)

				// then
				Expect(action).To(BeNil())
			})

			It("should return nil due to wrong data format", func() {
				// given
				tests := [][]byte{
					[]byte(fmt.Sprintf(`{"type":%q,"data":{"reportedTime":-1}}`, actions.ActionTypeReportTime)),
					[]byte(fmt.Sprintf(`{"type":%q,"data":{"reportedTime":15}}`, actions.ActionTypeReportTime)),
					[]byte(fmt.Sprintf(`{"type":%q,"data":{"reportedTime":0}}`, actions.ActionTypeReportTime)),
				}

				for _, t := range tests{
					// when
					action := actions.ValidateUserProvidedAction(t)

					// then
					Expect(action).To(BeNil())
				}
			})

			It("should return action", func() {
				// given
				time := 11.30
				b := []byte(fmt.Sprintf(`{"type":%q,"data":{"reportedTime":%.2f}}`, actions.ActionTypeReportTime, time))

				// when
				action := actions.ValidateUserProvidedAction(b)

				// then
				Expect(action).ToNot(BeNil())
				Expect(action.Type).To(Equal(actions.ActionTypeReportTime))
				Expect(*action.Data.ReportedTime).To(Equal(time))
				Expect(action.Data.ReporterID).To(BeNil())
			})
		})
	})
})
