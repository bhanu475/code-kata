# Go Solution

### Getting Started
``` sh
## build
$ go build -o todo

## help
$ ./todo 
Fetch TODOs from an API and print them

Usage:
  todo [flags]

Flags:
  -e, --endpoint string   API endpoint for Todos (default "https://jsonplaceholder.typicode.com/todos")
  -f, --filter string     Filter for fetching TODOs (even, odd, all) (default "all")
  -h, --help              help for todo
  -n, --number int        Number of TODOs to fetch

## run 

$./todo -n 2
ID:1, Title: delectus aut autem, Completed: false
ID:2, Title: quis ut nam facilis et officia qui, Completed: false


 
  ```

  ### Test

  ```sh
  $ go test ./...
  ?       github.com/bhanu475/code-kata   [no test files]
ok      github.com/bhanu475/code-kata/cmd/todo  (cached)
?       github.com/bhanu475/code-kata/internal/client   [no test files]
ok      github.com/bhanu475/code-kata/pkg/todo  (cached)
ok      github.com/bhanu475/code-kata/util      (cached)
  ```

  ### coverage
  ```sh
  go test -cover ./...
  ?       github.com/bhanu475/code-kata   [no test files]
ok      github.com/bhanu475/code-kata/cmd/todo  0.281s  coverage: 74.2% of statements
?       github.com/bhanu475/code-kata/internal/client   [no test files]
ok      github.com/bhanu475/code-kata/pkg/todo  2.135s  coverage: 88.6% of statements
ok      github.com/bhanu475/code-kata/util      0.004s  coverage: 85.7% of statements
  ```