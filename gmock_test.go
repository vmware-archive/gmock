package gmock_test

import (
	. "github.com/cfmobile/gmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GMock", func() {
	var subject *GMock

	BeforeEach(func() {
		subject = nil
	})

	Describe("MockTarget", func() {

	    Context("when creating a new GMock with a target", func() {
			var someVar = "some variable to mock"

			Context("and the target is not passed as a pointer", func() {
				var callConstructorPassingValue = func() {
					subject = CreateMockWithTarget(someVar)
				}

				It("should panic", func() {
					Expect(callConstructorPassingValue).To(Panic())
				})
			})

			Context("and the target is passed as a pointer", func() {
				var callConstructorPassingPointer = func() {
					subject = CreateMockWithTarget(&someVar)
				}

				It("should not panic", func() {
					Expect(callConstructorPassingPointer).NotTo(Panic())
				})

				It("should return a valid GMock object", func() {
					callConstructorPassingPointer()
					Expect(subject).NotTo(BeNil())
				})
			})
	    })
	})

})
