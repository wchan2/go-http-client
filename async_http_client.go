package httpclient

import "net/http"

type AsyncHttpClient interface {
	Send(req *http.Request) AsyncHttpClient
	Receive() (HttpResponse, error)
}

func NewAsyncHttpClient() AsyncHttpClient {
	return &asyncHttpClient{
		responseChannel: make(chan HttpResponse),
		errorChannel:    make(chan error),
	}
}

type asyncHttpClient struct {
	responseChannel chan HttpResponse
	errorChannel    chan error
	simpleHttpClient
}

func (client *asyncHttpClient) Send(request *http.Request) AsyncHttpClient {
	go func() {
		response, err := client.simpleHttpClient.Send(request)
		if err != nil {
			client.errorChannel <- err
		} else {
			client.responseChannel <- response
		}

		close(client.responseChannel)
		close(client.errorChannel)
	}()

	return client
}

func (client *asyncHttpClient) Receive() (HttpResponse, error) {
	select {
	case err := <-client.errorChannel:
		return nil, err
	case response := <-client.responseChannel:
		return response, nil
	}
}
