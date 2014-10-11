package httpclient

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type HttpClient interface {
	SendRequest(req *http.Request) (*HttpResponse, error)
}

type DefaultHttpClient struct {
	http.Client
}

type HttpResponse struct {
	Status int
	Body   []byte
}

func CreateRequest(method, url, body string) (*http.Request, error) {
	request, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (client *DefaultHttpClient) SendRequest(req *http.Request) (*HttpResponse, error) {
	var (
		err          error
		response     *http.Response
		responseBody []byte
	)
	response, err = client.Do(req)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	responseBody, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &HttpResponse{Status: response.StatusCode, Body: responseBody}, nil
}
