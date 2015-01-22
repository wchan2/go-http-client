package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "gopkg.in/check.v1"
)

var _ = Suite(&asyncHttpClientSuite{})

type asyncHttpClientSuite struct {
	client AsyncHttpClient
}

func (suite *asyncHttpClientSuite) SetUpTest(c *C) {
	suite.client = NewAsyncHttpClient()
}

func (suite *asyncHttpClientSuite) TestNewAsyncHttpClientReturnsNewClient(c *C) {
	c.Assert(NewAsyncHttpClient(), NotNil)
}

func (suite *asyncHttpClientSuite) TestSendsRequestWithBody(c *C) {
	sampleBody := []byte(`{"foo": "bar"}`)
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, serverRequest *http.Request) {
		receivedRequestBody, readRequestError := ioutil.ReadAll(serverRequest.Body)
		c.Assert(readRequestError, IsNil)
		c.Assert(receivedRequestBody, DeepEquals, sampleBody)
	}))

	clientRequest, err := http.NewRequest("GET", testServer.URL, bytes.NewReader(sampleBody))
	c.Assert(err, IsNil)
	suite.client.Send(clientRequest)
}

func (suite *asyncHttpClientSuite) TestReceivesResponseWithStatus(c *C) {
	var (
		err      error
		request  *http.Request
		response HttpResponse
	)
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, serverRequest *http.Request) {
		w.WriteHeader(http.StatusConflict)
	}))

	request, err = http.NewRequest("GET", testServer.URL, nil)
	c.Assert(err, IsNil)

	response, err = suite.client.Send(request).Receive()
	c.Assert(err, IsNil)
	c.Assert(response.Status(), Equals, http.StatusConflict)
}

func (suite *asyncHttpClientSuite) TestReceivesResponseWithHeaders(c *C) {
	var (
		err      error
		request  *http.Request
		response HttpResponse
	)
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, serverRequest *http.Request) {
		w.Header().Add("test-header", "test-header-value")
	}))

	request, err = http.NewRequest("GET", testServer.URL, nil)
	c.Assert(err, IsNil)

	response, err = suite.client.Send(request).Receive()
	c.Assert(err, IsNil)
	c.Assert(response.Header().Get("Test-Header"), Equals, "test-header-value")
}

func (suite *asyncHttpClientSuite) TestReceivesResponseWithBody(c *C) {
	var (
		clientRequest *http.Request
		response      HttpResponse
		err           error
	)
	sampleBody := []byte(`{"foo": "bar"}`)
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, serverRequest *http.Request) {
		w.Write(sampleBody)
	}))

	clientRequest, err = http.NewRequest("GET", testServer.URL, nil)
	c.Assert(err, IsNil)

	response, err = suite.client.Send(clientRequest).Receive()
	c.Assert(err, IsNil)
	c.Assert(response.Body(), NotNil)
}
