package httpclient

import "net/http"

type HttpResponse interface {
	Status() int
	Header() http.Header
	Body() []byte
}
type httpResponse struct {
	status  int
	headers http.Header
	body    []byte
}

func (response *httpResponse) Status() int {
	return response.status
}

func (response *httpResponse) Header() http.Header {
	return response.headers
}

func (response *httpResponse) Body() []byte {
	return response.body
}
