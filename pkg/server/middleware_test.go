package server

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Middleware", func() {

	Context("when middleware chain is invoked on a request", func() {
		BeforeEach(func() {
			testEngine = createGinEngine()
		})

		It("response contains a hostname in the response header", func() {
			r := performRequest(testEngine, "GET", "/current", map[string]string{})
			Expect(r.Code).To(Equal(http.StatusOK))
		})

		It("response contains a request ID in the response header", func() {
			r := performRequest(testEngine, "GET", "/current", map[string]string{})
			Expect(r.Code).To(Equal(http.StatusOK))
		})

		It("logs the request", func() {
			r := performRequest(testEngine, "GET", "/current", map[string]string{
				"X-Forwarded-For": "5.6.7.8,1.2.3.4",
				"X-Request-ID":    "request-id",
				"User-Agent":      "user-agent",
			})
			Expect(r.Code).To(Equal(http.StatusOK))
		})
})
})
