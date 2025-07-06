package main

import (
	"fmt"
	"os"

	"github.com/kroosec/sudoku-go"
)

func main() {
	boardString := "4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......"
	if len(os.Args) > 1 {
		boardString = os.Args[1]
	}

	board, err := sudoku.NewBoard(boardString)
	if err != nil {
		fmt.Printf("Error creating board: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Unsolved board:")
	fmt.Println(board)

	solvedBoard := sudoku.Solver(board)
	if solvedBoard != nil {
		fmt.Println("\nSolved board:")
		fmt.Println(solvedBoard)
	} else {
		fmt.Println("\nCould not solve the board.")
	}
}
