# Go Sudoku Solver

This is a Sudoku solver written in Go, inspired by Peter Norvig's essay on [Solving Every Sudoku Puzzle](https://norvig.com/sudoku.html).

It uses a backtracking search algorithm combined with constraint propagation to solve Sudoku puzzles efficiently.

## Features

- Solves any valid Sudoku puzzle.
- Handles various input formats for puzzles.
- Includes a comprehensive test suite with easy and hard puzzles.

## Usage

The package can be used to create and solve Sudoku boards.

```go
package main

import (
	"fmt"
	"github.com/kroosec/sudoku-go"
)

func main() {
	boardString := "4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......"
	board, err := sudoku.NewBoard(boardString)
	if err != nil {
		panic(err)
	}

	fmt.Println("Unsolved board:")
	fmt.Println(board)

	solvedBoard := sudoku.Solver(board)
	if solvedBoard != nil {
		fmt.Println("Solved board:")
		fmt.Println(solvedBoard)
	} else {
		fmt.Println("Could not solve the board.")
	}
}
```

## Building and Running

To build the solver, run:

```bash
go build ./cmd/sudoku
```

This will create an executable named `sudoku` in the project root. You can run it with a puzzle string as an argument:

```bash
./sudoku "4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......"
```

If no puzzle is provided, it will solve a default one.

You can also run it directly without building:

```bash
go run ./cmd/sudoku
```

## Running Tests

To run the tests for this project, use the standard `go test` command:

```bash
go test -v ./...
```

For more information on the algorithm, read [Peter Norvig's article](https://norvig.com/sudoku.html).
