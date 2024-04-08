package todo_test

import (
	"context"
	"sync"
	"testing"

	"github.com/bhanu475/code-kata/pkg/todo"
	"github.com/matryer/is"
)

func TestFetchAndPrintTodos(t *testing.T) {
	is := is.New(t)

	ctx := context.Background()
	endpoint := "https://jsonplaceholder.typicode.com/todos"
	numTodos := 10
	filter := "even"
	completed := true

	todos, err := todo.FetchAndPrintTodos(ctx, endpoint, numTodos, filter, completed)
	is.NoErr(err)
	//is.Equal(len(todos), numTodos)
	for _, todo := range todos {
		is.True(todo.ID%2 == 0)
	}

	filter = "odd"
	todos, err = todo.FetchAndPrintTodos(ctx, endpoint, numTodos, filter, completed)
	is.NoErr(err)
	//is.Equal(len(todos), numTodos)
	for _, todo := range todos {
		is.True(todo.ID%2 != 0)
	}
	filter = "all"
	todos, err = todo.FetchAndPrintTodos(ctx, endpoint, numTodos, filter, completed)
	is.NoErr(err)
	is.Equal(len(todos), numTodos)

	filter = "invalid"
	todos, err = todo.FetchAndPrintTodos(ctx, endpoint, numTodos, filter, completed)
	is.NoErr(err)
	is.Equal(len(todos), numTodos)
}

func TestFetchTodo(t *testing.T) {
	is := is.New(t)

	ctx := context.Background()
	endpoint := "https://jsonplaceholder.typicode.com/todos"
	numTodos := 10

	var wg sync.WaitGroup
	todoChan := make(chan *todo.Todo, numTodos)

	for i := 1; i <= numTodos; i++ {
		wg.Add(1)
		go todo.FetchTodo(ctx, endpoint, i, todoChan, &wg)
	}

	wg.Wait()
	close(todoChan)

	var todos []todo.Todo
	for todo := range todoChan {
		if todo != nil {
			todos = append(todos, *todo)
		}
	}

	is.Equal(len(todos), numTodos)
}
