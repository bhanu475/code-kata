package todo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func FetchTodo(ctx context.Context, endpoint string, id int, todoChan chan<- *Todo, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(fmt.Sprintf("%s/%d", endpoint, id))
	if err != nil {
		fmt.Println("Error:", err)
		todoChan <- nil
		return
	}
	var todo Todo
	json.NewDecoder(resp.Body).Decode(&todo)
	todoChan <- &todo
}

func FetchAndPrintTodos(ctx context.Context, endpoint string, numTodos int, filter string, completed bool) error {
	var wg sync.WaitGroup

	todoChan := make(chan *Todo, numTodos)
	j := 1
	for i := 1; i <= numTodos; i++ {
		var shouldFetch bool
		switch {
		case filter == "even":
			j = i * 2
			shouldFetch = true
		case filter == "odd":
			j = i*2 + 1
			shouldFetch = true
		case filter == "all":
			j = i
			shouldFetch = true
		default:
			j = i
			shouldFetch = true
		}

		if shouldFetch {
			wg.Add(1)
			go FetchTodo(ctx, endpoint, j, todoChan, &wg)
		}

	}

	go func() {
		wg.Wait()
		close(todoChan)
	}()
	for todo := range todoChan {
		if todo == nil {
			fmt.Printf("failed to fetch TODOs")
		}
		fmt.Printf("ID:%d, Title: %s, Completed: %t\n", todo.ID, todo.Title, todo.Completed)
	}
	return nil
}
