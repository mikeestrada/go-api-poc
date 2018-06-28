package test_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega/ghttp"
	"net/http/httptest"
	"net/http"
	"fmt"
	"github.com/onsi/gomega"
)

var _ = Describe("SampleApi", func() {
	var server *ghttp.Server

	BeforeEach(func() {
		server = ghttp.NewServer()
		server.AppendHandlers(ghttp.VerifyRequest("GET", "/sample"))
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("Given a sample endpoint", func() {
		Context("When a GET request is made", func() {
			It("should return a JSON containing 'hello world' ", func() {

				// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
				// pass 'nil' as the third parameter.
				req, err := http.NewRequest("GET", "/sample", nil)
				if err != nil {
					Fail(err.Error())
				}

				// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
				rr := httptest.NewRecorder()
				handler := http.HandlerFunc()

				// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
				// directly and pass in our Request and ResponseRecorder.
				handler.ServeHTTP(rr, req)

				// Check the status code is what we expect.
				gomega.Expect(rr.Code).To(gomega.Equal(http.StatusOK))

				// Check the status code is what we expect.
				if status := rr.Code; status != http.StatusOK {
					Fail("Handler returned wrong status")
				}

				// Check the response body is what we expect.
				expected := `{"alive": true}`
				if rr.Body.String() != expected {
					fmt.Printf("handler returned unexpected body: got %v want %v",
						rr.Body.String(), expected)
					Fail("Handler returned unexpected body.")
				}
			})
		})
	})
})
