package jenkins

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// Request struct
type Request struct {
	Host    string
	Params  url.Values
	Headers http.Header
	auth    BasicAuth
}

// BasicAuth auth
type BasicAuth struct {
	Username string
	Password string
}

// NewRequest new request struct
func NewRequest(host string) *Request {
	request := &Request{host, url.Values{}, http.Header{}, BasicAuth{}}
	return request
}

// SetHeader set request header
func (request *Request) SetHeader(key string, value string) *Request {
	request.Headers.Set(key, value)
	return request
}

// SetParam set request param
func (request *Request) SetParam(key string, value string) *Request {
	request.Params.Set(key, value)
	return request
}

// SetBasicAuth set auth
func (request *Request) SetBasicAuth(username, password string) *Request {
	request.auth = BasicAuth{
		Username: username,
		Password: password,
	}
	return request
}

// Post start post request
func (request *Request) Post(path string, body io.Reader, responseStruct interface{}) (*http.Response, error) {
	requestURL := request.buildURL(path)
	req, err := http.NewRequest("POST", requestURL, body)
	if err != nil {
		return nil, err
	}
	if len(request.auth.Username) > 0 && len(request.auth.Password) > 0 {
		req.SetBasicAuth(request.auth.Username, request.auth.Password)
	}

	for k := range request.Headers {
		req.Header.Add(k, request.Headers.Get(k))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(responseStruct)
	return resp, nil
}

func (request *Request) buildURL(path string) (requestURL string) {
	requestURL = request.Host + path
	if request.Params != nil {
		queryString := request.Params.Encode()
		if queryString != "" {
			requestURL = requestURL + "?" + queryString
		}
	}
	return
}
