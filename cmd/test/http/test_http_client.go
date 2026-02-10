package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/infrared-dao/protocols/fetchers"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

const defaultTimeout = 10 * time.Second

// TestHttpClient implements protocols.HttpClient for testing purposes.
type TestHttpClient struct {
	client  *http.Client
	timeout time.Duration
}

func NewTestHttpClient() *TestHttpClient {
	return &TestHttpClient{
		client: &http.Client{
			Timeout: defaultTimeout,
		},
		timeout: defaultTimeout,
	}
}

func (c *TestHttpClient) DoGraphQL(ctx context.Context, url, query string, variables, data any, opts ...fetchers.Option) error {
	type graphQLRequest struct {
		Query     string `json:"query"`
		Variables any    `json:"variables,omitempty"`
	}

	type gqlErr struct {
		Message string `json:"message"`
	}

	type gqlEnvelope struct {
		Data   any      `json:"data"`
		Errors []gqlErr `json:"errors"`
	}

	reqBody := graphQLRequest{
		Query:     query,
		Variables: variables,
	}

	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal GraphQL request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(reqJSON))
	if err != nil {
		return fmt.Errorf("failed to create GraphQL request: %w", err)
	}

	for _, opt := range opts {
		req = opt(req)
	}

	response, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("GraphQL request failed: %w", err)
	}
	defer func() {
		if closeErr := response.Body.Close(); closeErr != nil {
			log.Error().Err(closeErr).Msg("failed to close response body")
		}
	}()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read GraphQL response: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		log.Error().
			Int("statusCode", response.StatusCode).
			Str("url", url).
			Str("response", string(responseBody)).
			Msg("GraphQL request returned non-OK status")
		return fmt.Errorf("GraphQL request returned status %d", response.StatusCode)
	}

	// Unwrap GraphQL envelope
	var envelope gqlEnvelope
	envelope.Data = data

	if err := json.Unmarshal(responseBody, &envelope); err != nil {
		return fmt.Errorf("failed to unmarshal GraphQL response: %w", err)
	}

	// Check for GraphQL errors
	if len(envelope.Errors) > 0 {
		var msgs []string
		for _, e := range envelope.Errors {
			msgs = append(msgs, e.Message)
		}
		return fmt.Errorf("graphql errors: %s", fmt.Sprintf("%v", msgs))
	}

	return nil
}

func (c *TestHttpClient) DoJSON(ctx context.Context, method, url string, reqPayload, resPayload any, opts ...fetchers.Option) error {
	var reqBody io.Reader
	if reqPayload != nil {
		reqJSON, err := json.Marshal(reqPayload)
		if err != nil {
			return fmt.Errorf("failed to marshal request payload: %w", err)
		}
		reqBody = bytes.NewBuffer(reqJSON)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return fmt.Errorf("failed to create %s request: %w", method, err)
	}

	for _, opt := range opts {
		req = opt(req)
	}

	response, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("%s request failed: %w", method, err)
	}
	defer func() {
		if closeErr := response.Body.Close(); closeErr != nil {
			log.Error().Err(closeErr).Msg("failed to close response body")
		}
	}()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		log.Error().
			Int("statusCode", response.StatusCode).
			Str("method", method).
			Str("url", url).
			Str("response", string(responseBody)).
			Msg("Request returned non-OK status")
		return fmt.Errorf("request returned status %d", response.StatusCode)
	}

	if resPayload != nil {
		if err := json.Unmarshal(responseBody, resPayload); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}

func (c *TestHttpClient) DoBytes(ctx context.Context, method, url string, reqBodyBytes []byte, opts ...fetchers.Option) ([]byte, error) {
	var reqBody io.Reader
	if reqBodyBytes != nil {
		reqBody = bytes.NewBuffer(reqBodyBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create %s request: %w", method, err)
	}

	for _, opt := range opts {
		req = opt(req)
	}

	response, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s request failed: %w", method, err)
	}
	defer func() {
		if closeErr := response.Body.Close(); closeErr != nil {
			log.Error().Err(closeErr).Msg("failed to close response body")
		}
	}()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		log.Error().
			Int("statusCode", response.StatusCode).
			Str("method", method).
			Str("url", url).
			Str("response", string(responseBody)).
			Msg("Request returned non-OK status")
		return nil, fmt.Errorf("request returned status %d", response.StatusCode)
	}

	return responseBody, nil
}
