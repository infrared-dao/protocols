package fetchers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type HTTPParams struct {
	URL            string
	RequestBody    []byte
	RequestTimeout time.Duration
	Headers        map[string]string
}

const (
	DefaultRequestTimeout = 30 * time.Second
)

func timeout(params HTTPParams) time.Duration {
	if params.RequestTimeout <= 0 {
		return DefaultRequestTimeout
	}
	return params.RequestTimeout
}

func HTTPGet(ctx context.Context, params HTTPParams) ([]byte, error) {
	client := &http.Client{
		Timeout: timeout(params),
	}

	req, err := http.NewRequestWithContext(ctx, "GET", params.URL, nil)
	if err != nil {
		err = fmt.Errorf("failed to prepare GET request for URL ('%s'): %w", params.URL, err)
		log.Error().Msg(err.Error())
		return nil, err
	}

	// Set headers
	for key, value := range params.Headers {
		req.Header.Set(key, value)
	}

	response, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("failed to execute GET request for URL ('%s'): %w", params.URL, err)
		log.Error().Msg(err.Error())
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("failed to read response for GET request with URL ('%s'): %w", params.URL, err)
		log.Error().Msg(err.Error())
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("%d status code for GET request with URL ('%s')", response.StatusCode, params.URL)
		log.Error().Msg(err.Error())
		log.Info().Msgf("Response: '%s'", string(responseBody))
		return nil, err
	}

	return responseBody, nil
}

func HTTPPost(ctx context.Context, params HTTPParams) ([]byte, error) {
	client := &http.Client{
		Timeout: timeout(params),
	}

	req, err := http.NewRequestWithContext(ctx, "POST", params.URL, bytes.NewBuffer(params.RequestBody))
	if err != nil {
		err = fmt.Errorf("failed to create POST request for URL ('%s'): %w", params.URL, err)
		log.Error().Msg(err.Error())
		return nil, err
	}

	// Set headers
	for key, value := range params.Headers {
		req.Header.Set(key, value)
	}

	response, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("POST request for URL ('%s') failed: %w", params.URL, err)
		log.Error().Msg(err.Error())
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("failed to read response for POST request with URL ('%s'): %w", params.URL, err)
		log.Error().Msg(err.Error())
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("%d status code for POST request with URL ('%s')", response.StatusCode, params.URL)
		log.Error().Msg(err.Error())
		log.Info().Msgf("Response: '%s'", string(responseBody))
		return nil, err
	}

	return responseBody, nil
}
