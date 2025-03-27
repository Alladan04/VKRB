package mivar

import (
	"errors"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	timeout    time.Duration
}

// ClientConfig holds configuration for the Client
type ClientConfig struct {
	BaseURL    string
	Timeout    time.Duration
	HTTPClient *http.Client // Optional, will create default if nil
}

// New creates a new Client instance
func New(config ClientConfig) (*Client, error) {
	if config.BaseURL == "" {
		return nil, errors.New("base URL cannot be empty")
	}

	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second // Default timeout
	}

	httpClient := config.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: config.Timeout,
		}
	}

	return &Client{
		baseURL:    config.BaseURL,
		httpClient: httpClient,
		timeout:    config.Timeout,
	}, nil
}
