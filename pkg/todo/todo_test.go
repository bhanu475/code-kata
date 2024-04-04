package todo

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

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
	todoChan := make(chan *Todo)

	// Call the FetchTodo function
	wg.Add(1)
	go FetchTodo(ctx, server.URL, 1, todoChan, &wg)

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
	todoChan := make(chan *Todo)

	// Call the FetchTodo function
	wg.Add(1)
	go FetchTodo(ctx, server.URL, 1, todoChan, &wg)

	// Wait for the todo to be fetched
	todo := <-todoChan

	// Verify the fetched todo is nil
	if todo != nil {
		t.Errorf("Expected a nil todo, got %+v", todo)
	}

	// Wait for the FetchTodo function to complete
	wg.Wait()
}
