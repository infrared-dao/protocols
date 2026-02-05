package fetchers

import (
	"context"
	"net/http"
)

type Option func(req *http.Request) *http.Request

// WithHeader returns a function that adds a header to the request.
func WithHeader(key, value string) func(req *http.Request) *http.Request {
	return func(req *http.Request) *http.Request {
		req.Header.Set(key, value)
		return req
	}
}

// HttpClient is HTTP client.
type HttpClient interface {
	DoGraphQL(ctx context.Context, url, query string, variables, data any, opts ...Option) error
	DoJSON(ctx context.Context, method, url string, reqPayload, resPayload any, opts ...Option) error
	DoBytes(ctx context.Context, method, url string, reqBodyBytes []byte, opts ...Option) ([]byte, error)
}
