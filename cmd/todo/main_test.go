package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchAndPrintTodos(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a successful response with dummy data
		w.Write([]byte(`{"id": 1, "title": "Test Todo", "completed": true}`))
	}))
	defer server.Close()

	// Set up the test data
	ctx := context.Background()
	numTodos := 1
	filter := "all"
	completed := true

	// Call the FetchAndPrintTodos function
	err := FetchAndPrintTodos(ctx, server.URL, numTodos, filter, completed)

	// Verify the output
	assert.NoError(t, err)
}

func TestFetchAndPrintTodos_Error(t *testing.T) {
	// Create a test server that always returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	// Set up the test data
	ctx := context.Background()
	numTodos := 1
	filter := "all"
	completed := true

	// Call the FetchAndPrintTodos function
	err := FetchAndPrintTodos(ctx, server.URL, numTodos, filter, completed)

	// Verify the error
	assert.Error(t, err)
	assert.EqualError(t, err, "failed to fetch TODOs: internal server error")
}

func TestFetchAndPrintTodos_FilterEven(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a successful response with dummy data
		w.Write([]byte(`{"id": 2, "title": "Test Todo 2", "completed": true}`))
	}))
	defer server.Close()

	// Set up the test data
	ctx := context.Background()
	numTodos := 1
	filter := "even"
	completed := true

	// Call the FetchAndPrintTodos function
	err := FetchAndPrintTodos(ctx, server.URL, numTodos, filter, completed)

	// Verify the output
	assert.NoError(t, err)
}

func TestFetchAndPrintTodos_FilterOdd(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a successful response with dummy data
		w.Write([]byte(`{"id": 3, "title": "Test Todo 3", "completed": true}`))
	}))
	defer server.Close()

	// Set up the test data
	ctx := context.Background()
	numTodos := 1
	filter := "odd"
	completed := true

	// Call the FetchAndPrintTodos function
	err := FetchAndPrintTodos(ctx, server.URL, numTodos, filter, completed)

	// Verify the output
	assert.NoError(t, err)
}

func TestFetchAndPrintTodos_FilterAll(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a successful response with dummy data
		w.Write([]byte(`{"id": 1, "title": "Test Todo 1", "completed": true}`))
	}))
	defer server.Close()

	// Set up the test data
	ctx := context.Background()
	numTodos := 1
	filter := "all"
	completed := true

	// Call the FetchAndPrintTodos function
	err := FetchAndPrintTodos(ctx, server.URL, numTodos, filter, completed)

	// Verify the output
	assert.NoError(t, err)
}

func TestFetchAndPrintTodos_FilterInvalid(t *testing.T) {
	// Set up the test data
	ctx := context.Background()
	numTodos := 1
	filter := "invalid"
	completed := true

	// Call the FetchAndPrintTodos function
	err := FetchAndPrintTodos(ctx, "", numTodos, filter, completed)

	// Verify the error
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid filter: invalid (valid options: even, odd, all)")
}

func TestFetchAndPrintTodos_ConflictingFlags(t *testing.T) {
	// Set up the test data
	ctx := context.Background()
	numTodos := 1
	filter := "all"
	completed := true

	// Call the FetchAndPrintTodos function
	err := FetchAndPrintTodos(ctx, "", numTodos, filter, completed)

	// Verify the error
	assert.Error(t, err)
	assert.EqualError(t, err, "conflicting flags: cannot use even, odd, and all together")
}

func TestFetchAndPrintTodos_FetchError(t *testing.T) {
	// Set up the test data
	ctx := context.Background()
	numTodos := 1
	filter := "all"
	completed := true

	// Call the FetchAndPrintTodos function
	err := FetchAndPrintTodos(ctx, "invalid-url", numTodos, filter, completed)

	// Verify the error
	assert.Error(t, err)
	assert.EqualError(t, err, "failed to fetch TODOs: Get \"invalid-url\": unsupported protocol scheme \"\"")
}
