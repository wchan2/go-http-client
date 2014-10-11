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