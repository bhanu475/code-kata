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

func FetchTodo(ctx context.Context, endpoint string, id int, tchan chan<- *Todo, wg *sync.WaitGroup) {
	var todo Todo
	resp, err := http.Get(fmt.Sprintf("%s/%d", endpoint, id))
	if err != nil {
		fmt.Println("Error:", err)
		tchan <- nil
		wg.Done()
		return
	}
	json.NewDecoder(resp.Body).Decode(&todo)
	tchan <- &todo
	wg.Done()
}
