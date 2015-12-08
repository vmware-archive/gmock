package gmock_test

import (
	. "github.com/cfmobile/gmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	kOriginalValue = "original value"
	kMockValue = "mock value"
	kSecondMockValue = "second mock value"
)

var _ = Describe("GMock", func() {
	var subject *GMock // Test subject

	var constructSubject func()
	var panicked bool

	var someVar string // Target variable to mock in our tests
	var mockValue string // Variable containing the mock value to be set to the target

	BeforeEach(func() { // Reset all base values for each test
		subject = nil
		constructSubject = nil
		panicked = false
		someVar = kOriginalValue
		mockValue = kMockValue
	})

	var panicRecover = func() {
		panicked = recover() != nil
	}

	JustBeforeEach(func() {
		defer panicRecover()
		constructSubject()
	})

	Describe("MockTarget", func() {

		Context("when creating a new GMock with a target", func() {

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

				It("should not panic", func() {
					Expect(panicked).To(BeFalse())
				})

				It("should return a valid GMock object", func() {
					Expect(subject).NotTo(BeNil())
				})

				It("should not have altered the value of the target", func() {
					Expect(someVar).To(Equal(kOriginalValue))
				})

				It("should have backed up the pointer to the original target", func() {
					var originalPtr = subject.GetOriginal().Addr().Interface()
					Expect(originalPtr).To(Equal(&someVar))
				})

				It("should not have a mock value defined by default", func() {
					var mockTargetPtr = subject.GetTarget().Addr().Interface()
					Expect(mockTargetPtr).To(Equal(&someVar))
				})
			})
	    })
	})

	Describe("GMock", func() {

		BeforeEach(func() {
		    constructSubject = func() { // Construct a valid GMock for this set of tests
				subject = CreateMockWithTarget(&someVar)
			}
		})

	    Context("when calling Replace on a GMock object with a valid mock value", func() {

	        JustBeforeEach(func() { // It has to be a JustBeforeEach so that it happens after subject is constructed
	            subject.Replace(mockValue)
	        })

			It("should replace the value in the original var with the mock value", func() {
			    Expect(someVar).To(Equal(kMockValue))
			})

			It("should have retained the original value for restoring later", func() {
				Expect(subject.GetOriginal().Interface()).To(Equal(kOriginalValue))
			})

			It("should have a Target that points to the mock value", func() {
			    Expect(subject.GetTarget().Interface()).To(Equal(mockValue))
			})
	    })

		Context("when calling Replace on a GMock object with an invalid mock value", func() {

			It("should not have panicked when creating the mock", func() {
			    Expect(panicked).To(BeFalse())
			})

			Context("- mock value with a different type", func() {
				invalidMockValue := 21

				JustBeforeEach(func() {
					defer panicRecover()
					subject.Replace(invalidMockValue)
				})

				It("should have panicked", func() {
					Expect(panicked).To(BeTrue())
				})
			})

			Context("- mock value is a pointer to the same type as the target", func() {

				JustBeforeEach(func() {
				    defer panicRecover()
					subject.Replace(&mockValue)
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

		Context("before calling Restore on a GMock with a mock value set", func() {

			JustBeforeEach(func() {
			    subject.Replace(mockValue)
			})

			It("the target should equal to the mock value", func() {
			    Expect(someVar).To(Equal(kMockValue))
			})

			Context("when Restore is called", func() {
			    JustBeforeEach(func() {
			        subject.Restore()
			    })

				It("should have restored the target to the original value", func() {
				    Expect(someVar).To(Equal(kOriginalValue))
				})
			})
		})
	})
	
	Describe("MockTargetWithValue", func() {

		It("the target should be unaltered before the constructor is called", func() {
		    Expect(someVar).To(Equal(kOriginalValue))
		})

		Context("when constructor with pre-assigned mock value is called", func() {
			BeforeEach(func() {
				constructSubject = func() {
					subject = MockTargetWithValue(&someVar, mockValue)
				}
			})

			It("should have mocked the variable right away", func() {
			    Expect(someVar).To(Equal(kMockValue))
			})

			Context("when Replace is called a second time", func() {
				secondMockValue := kSecondMockValue

				JustBeforeEach(func() {
				    subject.Replace(secondMockValue)
				})

				It("should have mocked the variable again, with the second mock value", func() {
				    Expect(someVar).To(Equal(kSecondMockValue))
				})

				It("should contain a copy of the original value in Original, and not of the first mocked value (which is lost)", func() {
				    Expect(subject.GetOriginal().Interface()).To(Equal(kOriginalValue))
				})

				Context("and when Restore is finally called", func() {
					JustBeforeEach(func() {
					    subject.Restore()
					})

					It("should restore the initial value of the variable", func() {
					    Expect(someVar).To(Equal(kOriginalValue))
					})
				})
			})
		})
	})

})
