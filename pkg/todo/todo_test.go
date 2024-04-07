package todo_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/bhanu475/code-kata/pkg/todo"
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
	err := todo.FetchAndPrintTodos(ctx, server.URL, numTodos, filter, completed)
	if err != nil {
		t.Errorf("FetchAndPrintTodos() error = %v", err)
	}

	// Verify the output
	assert.NoError(t, err)
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
	err := todo.FetchAndPrintTodos(ctx, server.URL, numTodos, filter, completed)

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
	err := todo.FetchAndPrintTodos(ctx, server.URL, numTodos, filter, completed)

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
	err := todo.FetchAndPrintTodos(ctx, server.URL, numTodos, filter, completed)

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
	err := todo.FetchAndPrintTodos(ctx, "", numTodos, filter, completed)

	// Verify the error
	assert.Error(t, err)
	assert.EqualError(t, err, "endpoint can not be empty")
}

func TestFetchAndPrintTodos_FetchError(t *testing.T) {
	// Set up the test data
	ctx := context.Background()
	numTodos := 1
	filter := "all"
	completed := true

	// Call the FetchAndPrintTodos function
	err := todo.FetchAndPrintTodos(ctx, "invalid-url", numTodos, filter, completed)

	// Verify the error
	assert.Error(t, err)
	assert.EqualError(t, err, "Error: Get \"invalid-url/1\": unsupported protocol scheme \"\"")
}
func TestFetchTodo(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a successful response with dummy data
		w.Write([]byte(`{"id": 1, "title": "Test Todo"}`))
	}))
	defer server.Close()

	// Create a context and wait group
	ctx := context.Background()
	var wg sync.WaitGroup

	// Create a channel to receive the fetched todo
	todoChan := make(chan *todo.Todo)

	// Call the FetchTodo function
	wg.Add(1)
	go todo.FetchTodo(ctx, server.URL, 1, todoChan, &wg)

	// Wait for the todo to be fetched
	todo := <-todoChan

	// Verify the fetched todo
	if todo == nil {
		t.Error("Expected a non-nil todo, got nil")
	} else {
		if todo.ID != 1 {
			t.Errorf("Expected todo ID 1, got %d", todo.ID)
		}
		if todo.Title != "Test Todo" {
			t.Errorf("Expected todo title 'Test Todo', got '%s'", todo.Title)
		}
	}

	// Wait for the FetchTodo function to complete
	wg.Wait()
}

func TestFetchTodo_Error(t *testing.T) {
	// Create a test server that always returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	// Create a context and wait group
	ctx := context.Background()
	var wg sync.WaitGroup

	// Create a channel to receive the fetched todo
	todoChan := make(chan *todo.Todo)

	// Call the FetchTodo function
	wg.Add(1)
	go todo.FetchTodo(ctx, server.URL, 1, todoChan, &wg)

	// Wait for the todo to be fetched
	todo := <-todoChan

	// Verify the fetched todo is nil
	if todo != nil {
		t.Errorf("Expected a nil todo, got %+v", todo)
	}

	// Wait for the FetchTodo function to complete
	wg.Wait()
}
