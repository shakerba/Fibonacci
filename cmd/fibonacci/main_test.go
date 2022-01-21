package main_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("main", func() {
	var (
		command    *exec.Cmd
		session    *gexec.Session
		client     *http.Client
		listenPort int
	)

	BeforeEach(func() {

		listenPort = getPort()
		command = newServerCommand(listenPort)
		session = nil
		client = newHTTPClient()

	})

	JustBeforeEach(func() {
		var err error

		if command != nil {
			session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(func() error {
				_, err := client.Get(fmt.Sprintf("http://127.0.0.1:%d", listenPort))
				return err
			}).Should(Succeed())

		}
	})

	AfterEach(func() {

		if session != nil {
			session.Interrupt()
			Eventually(session).Should(gexec.Exit())
		}
	})

	It("responds to the default health checks", func() {
		resp, err := client.Get(fmt.Sprintf("http://127.0.0.1:%d/healthz", listenPort))
		Expect(err).NotTo(HaveOccurred())

		body, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.Body.Close()).To(Succeed())
		Expect(string(body)).To(ContainSubstring("OK"))
	})

	It("responds to /swagger-api handler with valid HTML", func() {
		resp, err := client.Get(fmt.Sprintf("http://127.0.0.1:%d/swagger-api/", listenPort))
		Expect(err).NotTo(HaveOccurred())

		body, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.Body.Close()).To(Succeed())
		Expect(string(body)).To(ContainSubstring("<html>"))
		Expect(string(body)).To(ContainSubstring("</html>"))
	})

	It("responds to /swagger handler with successful redirect", func() {
		resp, err := client.Get(fmt.Sprintf("http://127.0.0.1:%d/swagger/", listenPort))
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
	})

	It("responds to /swagger-api-json handler with valid JSON content", func() {
		resp, err := client.Get(fmt.Sprintf("http://127.0.0.1:%d/swagger-api-json/", listenPort))
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))

		body, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(body)).To(ContainSubstring("paths"))
	})

	It("does not execute gin in debug mode", func() {
		_, err := client.Get(fmt.Sprintf("http://127.0.0.1:%d/", listenPort))
		Expect(err).NotTo(HaveOccurred())

		Expect(session.Out).NotTo(gbytes.Say("GIN-debug"))
		Expect(session.Err).NotTo(gbytes.Say("GIN-debug"))
	})
})
