package main

import (
	"fmt"
	"net/http"

	"github.com/wchan2/go-httpclient"
)

func main() {
	var (
		req      *http.Request
		response *httpclient.HttpResponse
		err      error
	)
	req, err = httpclient.NewRequest("GET", "http://yahoo.com", `{"test": "test"}`)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("REQUEST", req)
	client := httpclient.NewAsyncHttpClient()
	client.SendRequest(req)

	response, err = client.ReceiveResponse()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("RESPONSE STATUS", response.Status)
	fmt.Println("RESPONSE HEADERS", response.Headers)
	fmt.Println("RESPONSE BODY", string(response.Body))
}
