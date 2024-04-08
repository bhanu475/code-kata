package todo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/bhanu475/code-kata/util"
)

// can move to separate file
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

func FetchAndPrintTodos(ctx context.Context, endpoint string, numTodos int, filter string, completed bool) ([]Todo, error) {
	var wg sync.WaitGroup
	var todos []Todo

	if !util.IsUrl(endpoint) {
		return nil, errors.New("endpoint can not be empty")
	}
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
		todos = append(todos, *todo)
	}
	return todos, nil
}
