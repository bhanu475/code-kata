package cmd

import (
	"context"
	"errors"

	"github.com/bhanu475/code-kata/pkg/todo"
	"github.com/bhanu475/code-kata/util"
	"github.com/spf13/cobra" // Using cobra for better CLI experience
)

const defaultEndpoint = "https://jsonplaceholder.typicode.com/todos"

var (
	endpoint  string
	numTodos  int
	filter    string // Use a string for filter options (even, odd, all)
	completed bool
)

var rootCmd = &cobra.Command{
	Use:  "todo",
	RunE: RootCmdRunE,
}

func RootCmdRunE(cmd *cobra.Command, args []string) error {
	e, err := cmd.Flags().GetString("endpoint")
	if err != nil {
		return err
	}
	if util.IsUrl(e) {
		endpoint = e
	} else {
		return errors.New("endpoint can not be empty")
	}

	n, err := cmd.Flags().GetInt("number")
	if err != nil {
		return err
	}
	if n <= 0 {
		return errors.New("number of todos should be greater than 0")
	}
	numTodos = n

	f, err := cmd.Flags().GetString("filter")
	if err != nil {
		return err
	}

	if f == "" {
		return errors.New("filter can not be empty")
	}
	filter = f

	switch filter {
	case "even", "odd", "all":
	default:
		return errors.New("filter should be one of even, odd, all")
	}

	return todo.FetchAndPrintTodos(context.Background(), endpoint, numTodos, filter, completed)
}

func RootCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("endpoint", "e", defaultEndpoint, "API endpoint for Todos")
	cmd.Flags().IntP("number", "n", 1, "Number of TODOs to fetch")
	//cmd.MarkFlagRequired("number")

	cmd.Flags().StringP("filter", "f", "all", "Filter for fetching TODOs (even, odd, all)")
	//cmd.MarkFlagRequired("filter")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	RootCmdFlags(rootCmd)
}
