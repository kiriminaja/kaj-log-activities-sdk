package requester

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/kiriminaja/kaj-log-activities-sdk/sdk/presentation"
)

// handler ...
type handler struct{}

// New ...
func New() Contract {
	return &handler{}
}

// RAW ...
func (request *handler) RAW(method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}

// GET request type get
func (request *handler) GET(urlBase string, header, params map[string]string) (presentation.Response, error) {
	var result presentation.Response
	// Create URL object
	u, err := url.Parse(urlBase)
	if err != nil {
		return result, err
	}

	// Add query parameters to URL
	q := u.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return result, err
	}
	if header != nil {
		for content, value := range header {
			req.Header.Set(content, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return result, err
	}

	if resp.StatusCode != 200 {
		return result, errors.New(string(body))
	}
	result.Cast(body, &result)

	return result, nil
}

// POST request type post
func (request *handler) POST(url string, header map[string]string, payload []byte) (presentation.Response, error) {
	var result presentation.Response
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return result, err
	}
	if header != nil {
		for content, value := range header {
			req.Header.Set(content, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}

	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return result, err
	}
	if resp.StatusCode != 200 {
		return result, errors.New(string(body))
	}
	result.Cast(body, &result)
	return result, nil
}

// WithBasicPOST request type post
func (request *handler) WithBasicPOST(url string, header map[string]string, payload []byte, username, password string) (presentation.Response, error) {
	var result presentation.Response
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)
	if err != nil {
		return result, err
	}
	if len(header) != 0 {
		for content, value := range header {
			req.Header.Set(content, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}

	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return result, err
	}
	if resp.StatusCode != 200 {
		return result, errors.New(string(body))
	}
	result.Cast(body, &result)
	return result, nil
}

// PUT request type post
func (request *handler) PUT(url string, header map[string]string, payload []byte) (presentation.Response, error) {
	var result presentation.Response
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return result, err
	}
	if header != nil {
		for content, value := range header {
			req.Header.Set(content, value)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return result, err
	}
	if (resp.StatusCode != 200) && (resp.StatusCode != 201) {
		return result, errors.New(string(body))
	}
	result.Cast(body, &result)
	return result, nil
}

// DELETE request type get
func (request *handler) DELETE(url string, header map[string]string) (presentation.Response, error) {
	var result presentation.Response
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return result, err
	}
	// set default headers
	req.Header.Add("Content-Type", "application/json")
	// add optional headers
	for content, value := range header {
		req.Header.Set(content, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return result, err
	}
	if resp.StatusCode != 200 {
		return result, errors.New(string(body))
	}
	result.Cast(body, &result)
	return result, nil
}
