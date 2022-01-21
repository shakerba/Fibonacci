package handlers_test

import (
	"encoding/json"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pex/fibonacci/pkg/config"
	"github.com/pex/fibonacci/pkg/handlers"
	"github.com/pex/fibonacci/pkg/models"
)

var _ = Describe("NextGet", func() {

	var (
		handler   *handlers.RouteHandler
		ctxParams = map[string]string{}
	)

	BeforeEach(func() {
		// Create the handler
		handler = &handlers.RouteHandler{
			ServiceName: "test-service-name",
			Version:     "test-version",
		}
	})

	It("returns 200 with valid status response body", func() {
		config.C = 1
		expected := &models.NextResponse{
	    Number: config.C,
			Err: nil,
		}

		expectedJSON, err := json.Marshal(expected)
		Expect(err).NotTo(HaveOccurred())

		// Perform request and validate output
		ctx, r := NewTestContext(ctxParams, nil, nil)
		handler.GetNext(ctx)
		Expect(r.Code).To(Equal(http.StatusOK))
		Expect(r.Body.String()).To(Equal(string(expectedJSON)))
	})
})
