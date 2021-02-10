package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// APIHost defines the API host URL for v1.
const APIHost = "https://flexpool.io/api/v1"

// Endpoint identifiers.
const (
	Miner Endpoint = iota
	Worker
	Pool
)

const WeiRatio = 0.000000001

// Endpoint type alias for the sendAPIRequest function.
type Endpoint int

// ResponseError contains the error-related data that could be returned from the API if it's used incorrectly.
type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Response is the primary container used for all responses from any API endpoint, containing the result and the error.
type Response struct {
	Error  ResponseError `json:"error"`
	Result interface{}   `json:"result"`
}

// sendAPIRequest is an internal function that takes an endpoint, and sends a GET request to the given query and method
// with the given set of parameters. Returns the Response container and nil on success, an empty Response and error on failure.
func sendAPIRequest(endpoint Endpoint, query string, method string, params []string) (Response, error) {
	var (
		err               error
		req               *http.Request
		resp              *http.Response
		responseBodyBytes []byte
		responseWrapped   Response
		httpClient        http.Client
	)

	// Build up the URL in format [host] + / + [endpoint] + query/params
	url := APIHost

	switch endpoint {
	case Miner:
		url += "/miner"
	case Worker:
		url += "/worker"
	case Pool:
		url += "/pool"
	default:
		return responseWrapped, errors.New("endpoint not supported")
	}

	// The pool endpoint doesn't use queries, so we'll build it into the URL for the other endpoints only.
	if endpoint != Pool {
		url += "/" + query
	}

	// Method and parameters come last
	url += "/" + method

	for _, param := range params {
		// There's a quirk where the worker endpoint uses '/' for parameter deliminators, where other endpoints use query
		// '?' signs, so we have to account for this.
		if endpoint == Worker {
			url += "/" + param
		} else {
			url += "?" + param
		}
	}

	// Build up the GET request - data is currently not used for the API, just the URL.
	if req, err = http.NewRequest("GET", url, bytes.NewBuffer([]byte{})); err != nil {
		return responseWrapped, err
	}

	// Depending on CORS settings we might need to explicitly define the content-type, so we'll set it just in case.
	req.Header.Set("Content-Type", "application/json")

	// Fire off the request to the API
	if resp, err = httpClient.Do(req); err != nil {
		return responseWrapped, err
	}

	// Parse the response and marshal it into the Response container
	if responseBodyBytes, err = ioutil.ReadAll(resp.Body); err != nil {
		return responseWrapped, err
	}

	if err = resp.Body.Close(); err != nil {
		return responseWrapped, err
	}

	if err = json.Unmarshal(responseBodyBytes, &responseWrapped); err != nil {
		return responseWrapped, err
	}

	return responseWrapped, nil
}
