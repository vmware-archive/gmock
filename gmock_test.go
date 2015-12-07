package gmock_test

import (
	. "github.com/cfmobile/gmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"reflect"
	"net/http"
)

var _ = Describe("GMock", func() {
	var subject *GMock

	Describe("MockTarget", func() {
	    Context("when a new GMock is created with a set target", func() {
			var someVar = http.DefaultClient

	        BeforeEach(func() {
	            subject = MockTarget(someVar)
	        })

			It("should return a valid GMock", func() {
			    Expect(subject).NotTo(BeNil())
			})

			It("should set the Original value of the mock object as the value of the target", func() {
				Expect(subject.GetOriginal()).To(Equal(reflect.ValueOf(someVar)))
			})

			It("should initially have the target unmocked", func() {
			    Expect(subject.GetTarget()).To(Equal(reflect.ValueOf(someVar)))
			})
	    })
	})

})
