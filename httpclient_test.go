package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&SimpleHttpClientSuite{})

type SimpleHttpClientSuite struct{}

func (suite *SimpleHttpClientSuite) TestCreateRequestReturnsRequest(c *C) {
	req, err := CreateRequest("GET", "http://localhost:8080/test", "")
	c.Assert(err, IsNil)
	c.Assert(req, NotNil)
}

func (suite *SimpleHttpClientSuite) TestSimpleHttpClientSendRequestReceivesProperResponse(c *C) {
	var (
		sampleResponseBody []byte = []byte(`{"hello": "world"}`)
		err                error
		req                *http.Request
		resp               *HttpResponse
	)
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(sampleResponseBody)
	}))

	req, err = http.NewRequest("GET", testServer.URL, nil)
	c.Assert(err, IsNil)

	client := &SimpleHttpClient{}
	resp, err = client.SendRequest(req)
	c.Assert(err, IsNil)

	c.Assert(resp.Body, DeepEquals, sampleResponseBody)
	c.Assert(resp.Status, Equals, http.StatusOK)
}

func (suite *SimpleHttpClientSuite) TestSimpleHttpClientSendRequestProperlySendsRequestBody(c *C) {
	var (
		sampleRequestBody   []byte = []byte(`{"good": "bye"}`)
		receivedRequestBody []byte
		requestReadError    error
		err                 error
		req                 *http.Request
	)
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		receivedRequestBody, requestReadError = ioutil.ReadAll(req.Body)

		c.Assert(err, IsNil)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"hello": "world"}`))
	}))

	req, err = http.NewRequest("GET", testServer.URL, bytes.NewReader(sampleRequestBody))
	c.Assert(err, IsNil)

	client := &SimpleHttpClient{}
	client.SendRequest(req)

	c.Assert(requestReadError, IsNil)
	c.Assert(receivedRequestBody, DeepEquals, sampleRequestBody)
}
