package contextBlock_test

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go-context-block/contextBlock"
	"time"
)

var _ = Describe("WithTimeout", func() {
	var (
		bg context.Context
	)
	BeforeSuite(func() {
		bg = context.Background()
	})
	Context("short-job", func() {
		It("succeeds", func() {
			Expect(contextBlock.WithTimeout(bg, 10*time.Second, func(ctx context.Context) {
				return
			})).Should(Succeed())
		})
	})
	Context("long-running job", func() {
		It("fails", func() {
			Expect(contextBlock.WithTimeout(bg, 1*time.Microsecond, func(ctx context.Context) {
				<-time.After(1 * time.Second)
				return
			})).ShouldNot(Succeed())
		})
	})
})
