package httpclient

import (
	"io/ioutil"
	"net/http"
)

type HttpClient interface {
	SendRequest(req *http.Request) (HttpResponse, error)
}

func NewSimpleHttpClient() HttpClient {
	return &simpleHttpClient{}
}

type simpleHttpClient struct {
	http.Client
}

func (client *simpleHttpClient) SendRequest(req *http.Request) (HttpResponse, error) {
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

	return &httpResponse{status: response.StatusCode, headers: response.Header, body: responseBody}, nil
}
