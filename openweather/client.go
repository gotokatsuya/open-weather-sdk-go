package openweather

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

// API endpoint base constants
const (
	APIEndpoint = "https://api.openweathermap.org"
)

// Client type
type Client struct {
	httpClient *http.Client
	endpoint   *url.URL

	apiKey string
}

// NewClient returns a new client instance.
func NewClient(apiKey string, httpClient *http.Client) (*Client, error) {
	c := &Client{
		httpClient: httpClient,
		apiKey:     apiKey,
	}
	u, err := url.Parse(APIEndpoint)
	if err != nil {
		return nil, err
	}
	c.endpoint = u
	return c, nil
}

// mergeQuery method
func (c *Client) mergeQuery(path string, q interface{}) (string, error) {
	v := reflect.ValueOf(q)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return path, nil
	}

	u, err := url.Parse(path)
	if err != nil {
		return path, err
	}

	qs, err := query.Values(q)
	if err != nil {
		return path, err
	}
	qs.Set("appid", c.apiKey)

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// NewRequest method
func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {

	if body != nil {
		merged, err := c.mergeQuery(path, body)
		if err != nil {
			return nil, err
		}
		path = merged
	}

	u, err := c.endpoint.Parse(path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// Do method
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}
	return resp, err
}

type ErrorResponse struct {
	ErrorCode    string `json:"cod,omitempty"`
	ErrorMessage string `json:"message,omitempty"`
}
