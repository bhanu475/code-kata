package cmd_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	cmd "github.com/bhanu475/code-kata/cmd/todo"
	"github.com/matryer/is"
	"github.com/spf13/cobra"
)

func execute(t *testing.T, c *cobra.Command, args ...string) (string, error) {
	t.Helper()

	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs(args)

	err := c.Execute()
	return strings.TrimSpace(buf.String()), err
}

func TestRootCmd(t *testing.T) {
	is := is.New(t)

	tt := []struct {
		args []string
		err  error
		out  string
	}{
		{
			args: []string{"-n", "0"},
			err:  errors.New("number of todos should be greater than 0"),
		},
		{
			args: nil,
			err:  errors.New("number of todos should be greater than 0"),
		},
		// {
		// 	args: []string{"-e", "https://jsonplaceholder.typicode.com/todos", "-n", "1", "-f", "all"},
		// 	err:  nil,
		// 	out:  "ID:1, Title: delectus aut autem, Completed: false",
		// },
		// {
		// 	args: []string{"--toggle"},
		// 	err:  nil,
		// 	out:  "ok",
		// },
	}

	root := &cobra.Command{Use: "root", RunE: cmd.RootCmdRunE}
	cmd.RootCmdFlags(root)

	for _, tc := range tt {
		out, err := execute(t, root, tc.args...)

		is.Equal(tc.err, err)

		if tc.err == nil {
			is.Equal(tc.out, out)
		}
	}
}
