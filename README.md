go-httpclient [![Build Status](https://travis-ci.org/wchan2/go-httpclient.png?branch=master)](https://travis-ci.org/wchan2/go-httpclient)
====

A wrapper around the standard library Go http client making it easier to use. The package includes a synchronous http client the (`simpleHttpClient`) and an asynchronous http client (`asyncHttpClient`). 

The asynchronous http client allows the request to be sent in a background goroutine and to have the response to be returned later. It blocks when the `ReceiveResponse` is called on an instance of the `asyncHttpClient` struct until a response is received; if the response is already ready, it will just return the response immediately.

### Usage

```go
// declaring the error
var (
	err      error
	request      *http.Request
	response *httpclient.HttpResponse
)

// create the request using a simpler string formats
request, err = httpclient.NewRequest("GET", "http://google.com", `{"test": "test"}`)

// making sure there are no errors upon creating the request
if err != nil {
	// ...
}

// creating a SimpleHttpClient
client := httpclient.NewSimpleHttpClient{}

// retrieving a response and possibly and error
response, err = client.SendRequest(request)

// check for response errors before proceeding
if err != nil {
	// ...
}

// print the status code, the response headers the response body
fmt.Println(response.Status())
fmt.Println(response.Header())
fmt.Println(string(response.Body()))
```

### Testing

A [MockHttpClient](https://gist.github.com/wchan2/92084704799b087d488f) can be used to mock an http client and injected into a function to isolate dependencies and increase the reliability of unit tests.

See below for an example.

```go
// create a new mock http client
mockClient := &mockHttpClient{}

// assign the error field to force it to be returned to test the error functionality
mockClient.err = errors.New("Some error")

// OR don't assign the error field and assign a *HttpResponse to be returned
// and test the case in which a certain response is returned
mockClient.response = &HttpResponse{Status: http.StatusOK, Body: []byte(`{"test": "test"}`)}
```

## License

go-httpclient is released under the [MIT License](http://www.opensource.org/licenses/MIT).