package httpclient

import (
	"net/http"

	. "gopkg.in/check.v1"
)

type httpResponseSuite struct {
	response *httpResponse
}

var _ = Suite(&httpResponseSuite{})

func (suite *httpResponseSuite) SetUpTest(c *C) {
	suite.response = &httpResponse{}
}

func (suite *httpResponseSuite) TestStatusReturnsStatus(c *C) {
	suite.response.status = http.StatusConflict
	c.Assert(suite.response.Status(), Equals, http.StatusConflict)
}

func (suite *httpResponseSuite) TestHeaderReturnsHeader(c *C) {
	sampleHeaderValue := []string{"test-header-value"}
	suite.response.headers = http.Header{"test-header": sampleHeaderValue}

	c.Assert(suite.response.Header()["test-header"], DeepEquals, sampleHeaderValue)
}

func (suite *httpResponseSuite) TestBodyReturnsBody(c *C) {
	sampleResponseBody := []byte("test-body")
	suite.response.body = []byte("test-body")

	c.Assert(suite.response.Body(), DeepEquals, sampleResponseBody)
}
