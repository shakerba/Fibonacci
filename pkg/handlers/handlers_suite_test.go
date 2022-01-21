package handlers_test

import (
	. "github.ibm.com/genctl/test"
	"github.com/gin-gonic/gin"
	"testing"
	"net/http/httptest"

"bytes"

"net/http"


"encoding/json"

)

func TestHandlers(t *testing.T) {
	TestPackage(t, "Handlers Suite")
}


// NewTestContext builds a context for testing handlers
func NewTestContext(params map[string]string, jsonBody interface{}, headers map[string]string) (*gin.Context, *httptest.ResponseRecorder) {

	// Set request path and recorder
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	// Set a test request ID
	ctx.Set("reqID", "TestRequestId")

	// Set any required parameters
	if params != nil {
		ctx.Params = gin.Params{}
		for k, v := range params {
			ctx.Params = append(ctx.Params, gin.Param{Key: k, Value: v})
		}
	}

	// Set JSON body if provided. Defaulting requests with a body to POST, and all other requests to GET. The
	// handler is called directly, so the method is only used for mux.
	if jsonBody != nil {
		jsonString, _ := json.Marshal(jsonBody)
		ctx.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString(string(jsonString)))
		ctx.Request.Header.Add("Content-Type", "application/json")
	} else {
		ctx.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString(""))
	}

	for header, value := range headers {
		ctx.Request.Header.Add(header, value)
	}

	return ctx, w
}
