package client

import "net/http"

func NewHTTPClient() *http.Client {
	// Implement custom HTTP client if needed
	return http.DefaultClient
}
