package contextBlock_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestContextBlock(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ContextBlock Suite")
}
