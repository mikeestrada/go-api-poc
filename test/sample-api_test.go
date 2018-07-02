package test_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"net/http/httptest"
	"net/http"
	"github.homedepot.com/go-api-poc/api/sample"
)

var _ = Describe("SampleApi", func() {
	var server *ghttp.Server
	var client *sample.SampleApi

	BeforeEach(func() {
		server = ghttp.NewServer()
		client = &sample.SampleApi{DBConnection: "thing"}
	})

	AfterEach(func() {
		server.Close()
	})

	// using gomega
	Describe("Given a sample endpoint", func() {
		Context("When a GET request is made", func() {
			It("should return a JSON containing 'something' ", func() {
				// Check the status code is what we expect.
				server.AppendHandlers(ghttp.VerifyRequest("GET", "/sample/configured"))
				response, e := http.Get(server.URL() + "/sample/configured")

				Ω(e).ShouldNot(HaveOccurred())
				Ω(response.StatusCode).Should(Equal(http.StatusOK))

				client.SampleHandler()
				Ω(server.ReceivedRequests()).Should(HaveLen(1))
			})
		})
	})

	Describe("Given a sample endpoint", func() {
		Context("When a GET request is made", func() {
			It("should return a JSON containing 'hello world' ", func() {

				server.AppendHandlers(ghttp.VerifyRequest("GET", "/sample"))
				// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
				// pass 'nil' as the third parameter.
				req, err := http.NewRequest("GET", "/sample", nil)
				if err != nil {
					Fail(err.Error())
				}

				// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(client.SampleApiHandler)

				// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
				// directly and pass in our Request and ResponseRecorder.
				handler.ServeHTTP(rr, req)

				Ω(rr.Code).Should(Equal(http.StatusOK))
				// Check the status code is what we expect.

				// Check the status code is what we expect.
				if status := rr.Code; status != http.StatusOK {
					Fail("Handler returned wrong status")
				}

				// Check the response body is what we expect.
				expected := `{"alive": true}`
				Ω(rr.Body.String()).Should(Equal(expected))
			})
		})
	})
})
