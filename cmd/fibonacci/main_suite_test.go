package main_test

import (

	"encoding/json"
	"fmt"

	"net/http"

	"os/exec"

	"testing"
	"crypto/tls"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"net/http/httptest"
)

var (
	serverPath string
	mux           *http.ServeMux
	fakeServer *httptest.Server
)

func TestApiRestServer(t *testing.T) {
	TestPackage(t, "Server Suite")
}

var _ = SynchronizedBeforeSuite(func() []byte {
	server, err := gexec.Build("github.com/pex/fibonacci/cmd/fibonacci")
	Expect(err).NotTo(HaveOccurred())

	createFakeServer()

	payload, err := json.Marshal(map[string]string{
		"server": server,
	})
	return payload
}, func(payload []byte) {
	context := map[string]string{}

	err := json.Unmarshal(payload, &context)
	Expect(err).NotTo(HaveOccurred())

	serverPath = context["server"]

})

var _ = SynchronizedAfterSuite(func() {
}, func() {
	gexec.CleanupBuildArtifacts()
	fakeServer.Close()
})

func newServerCommand(port int, extraFlags ...string) *exec.Cmd {

	ret := exec.Command(
		serverPath,
		append([]string{
			"--httpAddr", fmt.Sprintf("127.0.0.1:%d", port),
		}, extraFlags...)...,
	)

	return ret
}

const IntegrationBasePort = 9000

func getPort() int {
	return IntegrationBasePort + GinkgoParallelNode()
}

func newHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}

func createFakeServer() {
	// Create a fake endpoint
	mux = http.NewServeMux()
	fakeServer = httptest.NewServer(mux)
	mux.HandleFunc("/current",
		func(w http.ResponseWriter, r *http.Request) {
			// w.Header().Set("Transaction-Id", "testtransactionid")
			var respbody string
			status := http.StatusOK
			respbody = fmt.Sprintf(`{"current":"1"}`)
			w.WriteHeader(status)
			fmt.Fprint(w, respbody)
		})
	mux.HandleFunc("/next",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			respbody := fmt.Sprintf(`{"next":"1"}`)
			fmt.Fprint(w, respbody)
			return
		})
		mux.HandleFunc("/previous",
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				respbody := fmt.Sprintf(`{"previous":"1"}`)
				fmt.Fprint(w, respbody)
				return
		})
}
