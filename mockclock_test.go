package chronos_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"

	"github.com/mbenford/chronos"
)

var _ = Describe("MockClock", func() {
	Context("Patching", func() {
		It("patches the Now method", func() {
			// Arrange
			called := false
			fake := chronos.NewMock()

			// Act
			fake.Now.Patch(func() time.Time {
				called = true
				return time.Now()
			})

			// Assert
			fake.Clock().Now()
			Expect(called).To(BeTrue())
		})

		It("patches the Sleep method", func() {
			// Arrange
			called := false
			fake := chronos.NewMock()

			// Act
			fake.Sleep.Patch(func(d time.Duration) {
				called = true
			})

			// Assert
			fake.Clock().Sleep(time.Hour)
			Expect(called).To(BeTrue())
		})

		It("patches the After method", func() {
			// Arrange
			called := false
			fake := chronos.NewMock()

			// Act
			fake.After.Patch(func(d time.Duration) <-chan time.Time {
				called = true
				return make(chan time.Time, 1)
			})

			// Assert
			fake.Clock().After(time.Hour)
			Expect(called).To(BeTrue())
		})

		It("patches the AfterFunc method", func() {
			// Arrange
			called := false
			fake := chronos.NewMock()

			// Act
			fake.AfterFunc.Patch(func(d time.Duration, f func()) chronos.Timer {
				called = true
				return nil
			})

			// Assert
			fake.Clock().AfterFunc(time.Hour, func() {})
			Expect(called).To(BeTrue())
		})

		It("patches the NewTimer method", func() {
			// Arrange
			called := false
			fake := chronos.NewMock()

			// Act
			fake.Timer.Patch(func(d time.Duration) chronos.Timer {
				called = true
				return nil
			})

			// Assert
			fake.Clock().NewTimer(time.Hour)
			Expect(called).To(BeTrue())
		})

		It("patches the NewTicker method", func() {
			// Arrange
			called := false
			fake := chronos.NewMock()

			// Act
			fake.Ticker.Patch(func(d time.Duration) chronos.Ticker {
				called = true
				return nil
			})

			// Assert
			fake.Clock().NewTicker(time.Hour)
			Expect(called).To(BeTrue())
		})
	})
})
