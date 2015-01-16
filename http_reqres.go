package httpclient

import (
	"net/http"
	"strings"
)

func NewRequest(method, url, body string) (*http.Request, error) {
	request, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	return request, nil
}

type HttpResponse struct {
	Status  int
	Headers map[string][]string
	Body    []byte
}
