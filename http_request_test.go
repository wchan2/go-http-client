package httpclient

import . "gopkg.in/check.v1"

type responseRequestTestSuite struct{}

var _ = Suite(&responseRequestTestSuite{})

func (suite *responseRequestTestSuite) TestNewRequestReturnsRequest(c *C) {
	req, err := NewRequest("GET", "http://localhost:8080/test", "")
	c.Assert(err, IsNil)
	c.Assert(req, NotNil)
}
