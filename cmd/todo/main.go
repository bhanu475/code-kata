package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/spf13/cobra" // Using cobra for better CLI experience

	"github.com/bhanu475/code-kata/internal/client" // Optional, if using a custom http client
	"github.com/bhanu475/code-kata/pkg/todo"
)

const defaultEndpoint = "https://jsonplaceholder.typicode.com/todos"

var (
	endpoint string
	numTodos int
	filter   string // Use a string for filter options (even, odd, all)

	// Flags for filtering and configuration (consider adding more as needed)
	filterEven bool
	filterOdd  bool
	fetchAll   bool
	completed  bool // Flag to filter by completed status
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Fetch and print TODOs",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		switch filter {
		case "even":
			filterEven = true
		case "odd":
			filterOdd = true
		case "all":
			fetchAll = true
		default:
			err = fmt.Errorf("invalid filter: %s (valid options: even, odd, all)", filter)
		}
		if err != nil {
			return err
		}

		if (filterEven && filterOdd) || (filterEven && fetchAll) || (filterOdd && fetchAll) {
			return fmt.Errorf("conflicting flags: cannot use even, odd, and all together")
		}

		return FetchAndPrintTodos(context.Background())
	},
}

func init() {
	rootCmd.Flags().StringVarP(&endpoint, "endpoint", "e", defaultEndpoint, "API endpoint for Todos")
	rootCmd.Flags().IntVarP(&numTodos, "number", "n", 20, "Number of TODOs to fetch")
	rootCmd.Flags().StringVarP(&filter, "filter", "f", "", "Filter for fetching TODOs (even, odd, all)")
	rootCmd.MarkFlagRequired("filter") // Make the filter flag required

	// Define flag handlers for filtering options (consider adding more)
	rootCmd.Flags().BoolVarP(&filterEven, "even", "", false, "Filter for even-numbered TODOs")
	rootCmd.Flags().BoolVarP(&filterOdd, "odd", "", false, "Filter for odd-numbered TODOs")
	rootCmd.Flags().BoolVarP(&fetchAll, "all", "", false, "Fetch all TODOs (ignores even/odd filters)")
	rootCmd.Flags().BoolVarP(&completed, "completed", "c", false, "Filter for completed TODOs (only shows completed)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func FetchAndPrintTodos(ctx context.Context,) error {
	client := client.NewHTTPClient()

	var wg sync.WaitGroup
	todoChan := make(chan *todo.Todo, numTodos)

	go func() {
		defer close(todoChan)
		for todo := range todoChan {
			if completed && !todo.Completed {
				continue
			}
			fmt.Printf("Title: %s, Completed: %t\n", todo.Title, todo.Completed)
			wg.Done()
		}
	}()
	j := 1
	for i := 1; i <= numTodos; i++ {
		var shouldFetch bool

		switch {
		case filterEven && i%2 == 0:
			j = i * 2
			shouldFetch = true
		case filterOdd && i%2 != 0:
			j = i*2 + 1
			shouldFetch = true
		case fetchAll:
			j = i
			shouldFetch = true
		default:
			shouldFetch = true
		}

		if !shouldFetch {
			continue
		}
		wg.Add(1)
		go todo.FetchTodo(ctx, client, endpoint, j, todoChan, &wg)
	}

	go func() {
		wg.Wait()
		close(todoChan)
	}()

	for todo := range todoChan {
		fmt.Printf("Title: %s, Completed: %t\n", todo.Title, todo.Completed)
	}

	return nil
}
