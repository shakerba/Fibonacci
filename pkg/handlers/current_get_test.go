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

var _ = Describe("CurrentGet", func() {

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
		config.B = 1
		expected := &models.CurrentResponse{
	    Number: config.B,
			Err: nil,
		}

		expectedJSON, err := json.Marshal(expected)
		Expect(err).NotTo(HaveOccurred())

		// Perform request and validate output
		ctx, r := NewTestContext(ctxParams, nil, nil)
		handler.GetCurrent(ctx)
		Expect(r.Code).To(Equal(http.StatusOK))
		Expect(r.Body.String()).To(Equal(string(expectedJSON)))
	})
})
