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

				It("should not have altered the value of the target", func() {
					Expect(someVar).To(Equal("some variable to mock"))
				})

				validMockTests(&someVar)
			})
	    })
	})

	Describe("GMock", func() {
		someVar := "some variable to mock"

		BeforeEach(func() {
		    constructSubject = func() {
				subject = CreateMockWithTarget(&someVar)
			}
		})

	    Context("when calling Replace on a GMock object with a valid mock value", func() {
			mockVar := "this is a fake value"

	        JustBeforeEach(func() { // It has to be a JustBeforeEach so that it happens after subject is constructed
	            subject.Replace(mockVar)
	        })

			It("should replace the value in the original var with the mock value", func() {
			    Expect(someVar).To(Equal(mockVar))
			})

			It("should have a Target that points to the mock value", func() {
			    Expect(subject.GetTarget().Interface()).To(Equal(mockVar))
			})

			It("should have an unaltered pointer to the original target", func() {
			    originalPtr := subject.GetOriginal().Addr().Interface()
				Expect(originalPtr).To(Equal(&someVar))
			})
	    })

		Context("when calling Replace on a GMock object with an invalid mock value", func() {
			Context("- mock value with a different type", func() {
				mockVar := 21

				JustBeforeEach(func() {
					defer panicRecover()
					subject.Replace(mockVar)
				})

				It("should have panicked", func() {
					Expect(panicked).To(BeTrue())
				})
			})

			Context("- mock value is nil", func() {
			    JustBeforeEach(func() {
					defer panicRecover()
			        subject.Replace(nil)
			    })

				It("should not have panicked", func() {
				    Expect(panicked).To(BeFalse())
				})

				It("should have assigned a default value of that type to the mock", func() {
				    Expect(someVar).To(Equal(""))
				})
			})
		})
	})
	
})
