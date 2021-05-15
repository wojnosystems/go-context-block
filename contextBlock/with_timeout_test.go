package contextBlock_test

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/wojnosystems/go-context-block/contextBlock"
	"time"
)

var _ = Describe("WithTimeout", func() {
	var (
		bg context.Context
	)
	BeforeSuite(func() {
		bg = context.Background()
	})
	It("has a deadline", func() {
		contextBlock.WithTimeout(bg, 10*time.Second, func(ctx context.Context) {
			_, hasDeadline := ctx.Deadline()
			Expect(hasDeadline).Should(BeTrue())
		})
	})
	Context("after block completes", func() {
		It("cancels the context", func() {
			var escapedContext context.Context
			contextBlock.WithTimeout(bg, 10*time.Second, func(ctx context.Context) {
				Expect(ctx.Err()).ShouldNot(HaveOccurred())
				escapedContext = ctx
			})
			Expect(escapedContext.Err()).Should(HaveOccurred())
		})
	})
})
