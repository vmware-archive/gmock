package gmock_test

import (
	. "github.com/cfmobile/gmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GMock", func() {
	var subject *GMock
	var constructSubject func()
	var panicked bool

	BeforeEach(func() {
		subject = nil
		constructSubject = nil
		panicked = false
	})

	var panicRecover = func() {
		panicked = recover() != nil
	}

	JustBeforeEach(func() {
		defer panicRecover()
		constructSubject()
	})

	var validMockTests = func(target interface{}) {
		It("should not panic", func() {
			Expect(panicked).To(BeFalse())
		})

		It("should return a valid GMock object", func() {
			Expect(subject).NotTo(BeNil())
		})

		It("should have backed up the pointer to the original target", func() {
			var originalPtr = subject.GetOriginal().Addr().Interface()
			Expect(originalPtr).To(Equal(target))
		})

		It("should not have a mock value defined by default", func() {
			var mockTargetPtr = subject.GetTarget().Addr().Interface()
			Expect(mockTargetPtr).To(Equal(target))
		})
	}

	Describe("MockTarget", func() {
	    Context("when creating a new GMock with a target", func() {
			someVar := "some variable to mock"

			Context("and the target is not passed as a pointer", func() {
				BeforeEach(func() {
					constructSubject = func() {
						subject = CreateMockWithTarget(someVar)
					}
				})

				It("should panic", func() {
					Expect(panicked).To(BeTrue())
				})

				It("should not have constructed the mock", func() {
				    Expect(subject).To(BeNil())
				})
			})

			Context("and the target is passed as a pointer", func() {
				BeforeEach(func() {
					constructSubject = func() {
						subject = CreateMockWithTarget(&someVar)
					}
				})

				validMockTests(&someVar)
			})
	    })
	})

})
