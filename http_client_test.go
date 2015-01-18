package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "gopkg.in/check.v1"
)

var _ = Suite(&simpleHttpClientSuite{})

type simpleHttpClientSuite struct {
	simpleHttpClient HttpClient
}

func (suite *simpleHttpClientSuite) SetUpTest(c *C) {
	suite.simpleHttpClient = &simpleHttpClient{}
}

func (suite *simpleHttpClientSuite) TestNewSimpleHttpClientReturnsNewClient(c *C) {
	c.Assert(NewAsyncHttpClient(), NotNil)
}

func (suite *simpleHttpClientSuite) TestSimpleHttpClientSendRequestReceivesResponse(c *C) {
	var (
		sampleResponseBody []byte = []byte(`{"hello": "world"}`)
		err                error
		req                *http.Request
		resp               HttpResponse
	)
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(sampleResponseBody)
	}))

	req, err = http.NewRequest("GET", testServer.URL, nil)
	c.Assert(err, IsNil)

	resp, err = suite.simpleHttpClient.SendRequest(req)
	c.Assert(err, IsNil)
	c.Assert(resp, NotNil)
}

func (suite *simpleHttpClientSuite) TestSimpleHttpClientSendRequestReceivesRequestBody(c *C) {
	var (
		sampleRequestBody []byte = []byte(`{"good": "bye"}`)
		err               error
		resp              HttpResponse
		clientRequest     *http.Request
	)
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, serverRequest *http.Request) {
		receivedRequestBody, requestReadError := ioutil.ReadAll(serverRequest.Body)
		c.Assert(requestReadError, IsNil)
		c.Assert(receivedRequestBody, DeepEquals, sampleRequestBody)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"hello": "world"}`))
	}))

	clientRequest, err = http.NewRequest("GET", testServer.URL, bytes.NewReader(sampleRequestBody))
	c.Assert(err, IsNil)

	resp, err = suite.simpleHttpClient.SendRequest(clientRequest)
	c.Assert(err, IsNil)
	c.Assert(resp, NotNil)
}

func (suite *simpleHttpClientSuite) TestSimpleHttpClientSendRequestReturnsResponseStatus(c *C) {
	var (
		err  error
		req  *http.Request
		resp HttpResponse
	)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req, err = http.NewRequest("GET", testServer.URL, nil)
	c.Assert(err, IsNil)

	resp, err = suite.simpleHttpClient.SendRequest(req)
	c.Assert(err, IsNil)

	c.Assert(resp, NotNil)
	c.Assert(resp.Status(), Equals, http.StatusOK)
}

func (suite *simpleHttpClientSuite) TestSimpleHttpClientSendRequestReturnsResponseHeaders(c *C) {
	var (
		err  error
		req  *http.Request
		resp HttpResponse
	)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("test-header", "test-header-value")
	}))

	req, err = http.NewRequest("GET", testServer.URL, nil)
	c.Assert(err, IsNil)

	resp, err = suite.simpleHttpClient.SendRequest(req)
	c.Assert(err, IsNil)

	c.Assert(resp, NotNil)
	c.Assert(resp.Header().Get("Test-Header"), DeepEquals, "test-header-value")
}

func (suite *simpleHttpClientSuite) TestSimpleHttpClientSendRequestReturnsResponseBody(c *C) {
	var (
		err         error
		req         *http.Request
		resp        HttpResponse
		requestBody = []byte(`{"foo": "bar"}`)
	)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write(requestBody)
	}))

	req, err = http.NewRequest("GET", testServer.URL, nil)
	c.Assert(err, IsNil)

	resp, err = suite.simpleHttpClient.SendRequest(req)
	c.Assert(err, IsNil)

	c.Assert(resp, NotNil)
	c.Assert(resp.Body(), DeepEquals, requestBody)
}
