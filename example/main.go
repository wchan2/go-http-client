package main

import (
	"fmt"
	"net/http"

	"github.com/wchan2/go-httpclient"
)

func asyncHttpClientExample() {
	var (
		request  *http.Request
		response httpclient.HttpResponse
		err      error
	)
	request, err = httpclient.NewRequest("GET", "http://yahoo.com", `{"test": "test"}`)
	if err != nil {
		panic(err.Error())
	}

	client := httpclient.NewAsyncHttpClient().Send(request)
	response, err = client.Receive()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(response.Status())
	fmt.Println(response.Header())
	fmt.Println(string(response.Body()))
}

func simpleHttpClientExample() {
	var (
		request  *http.Request
		response httpclient.HttpResponse
		err      error
	)
	request, err = httpclient.NewRequest("GET", "http://yahoo.com", `{"test": "test"}`)
	if err != nil {
		panic(err.Error())
	}

	response, err = httpclient.NewSimpleHttpClient().Send(request)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(response.Status())
	fmt.Println(response.Header())
	fmt.Println(string(response.Body()))
}

func main() {
	fmt.Println("Running SimpleHttpClient example")
	simpleHttpClientExample()

	fmt.Println("Running AsyncHttpClient example")
	asyncHttpClientExample()
}
