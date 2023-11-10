package testing

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestkickoffApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "kickoffApi Suite")
}

var _ = BeforeSuite(func() {
	err := StartServer()
	Expect(err).To(BeNil())
})

var _ = AfterSuite(func() {
	err := ShutdownServer()
	Expect(err).To(BeNil())
})
