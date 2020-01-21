package sudoku

func nextEmptySquare(b *Board) (int, int) {
	minRow, minColumn, minCount := -1, -1, 10
	for i := 0; i < numRows; i++ {
		for j := 0; j < numColumns; j++ {
			possible := len(b[i][j])
			if possible > 1 && possible < minCount {
				minRow, minColumn, minCount = i, j, possible
			}
		}
	}
	return minRow, minColumn
}

func Solver(b *Board) *Board {
	row, column := nextEmptySquare(b)
	if row == -1 || column == -1 {
		return DuplicateBoard(b)
	}
	// try each value for this empty square
	for value := 1; value <= 9; value++ {
		if !b.valuePossible(row, column, value) {
			continue
		}
		// Apply modifications to a duplicate board.
		newBoard := DuplicateBoard(b)
		if err := newBoard.assign(row, column, value); err != nil {
			// value present in unit
			continue
		}
		// try solving the board with this value
		if solved := Solver(newBoard); solved != nil {
			return solved
		}
	}
	return nil
}
