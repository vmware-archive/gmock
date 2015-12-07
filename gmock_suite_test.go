package gmock_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGmock(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gmock Suite")
}
