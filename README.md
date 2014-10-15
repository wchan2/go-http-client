go-httpclient
=============

A wrapper around the standard library Go http client making it easier to use

### Usage

```go
// declaring the error
var (
	err      error
	req      *http.Request
	response *httpclient.HttpResponse
)

// create the request using a simpler string formats

req, err = httpclient.CreateRequest("GET", "http://google.com", `{"test": "test"}`)

// making sure there are no errors upon creating the request
if err != nil {
	// ...
}

// creating a new SimpleHttpClient directly
client := &httpclient.SimpleHttpClient{}

// retrieving a response and possibly and error
response, err = client.SendRequest(req)

// check for response errors before proceeding
if err != nil {
	// ...
}

// print the status code and the response body
fmt.Println(response.Status)
fmt.Println(string(response.Body))
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
