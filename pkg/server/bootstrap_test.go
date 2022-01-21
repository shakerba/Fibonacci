package server

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gin engine creating", func() {
	Context("When a new engine is successfully created", func() {

		It("Creates correct number of routes", func() {
			Expect(testEngine.Routes()).To(HaveLen(5))
		})

		It("Contains the expected number of handlers in the middleware chain", func() {
			Expect(testEngine.Handlers).To(HaveLen(3))
		})
	})
})
