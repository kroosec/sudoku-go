package sudoku

func nextEmptySquare(b *Board) (int, int) {
	minRow, minColumn, minCount := -1, -1, numDigits+1

	// Search for the empty square that has the least amount of possible values.
	for i := range numRows {
		for j := range numColumns {
			possible, _ := b.CountPossible(i, j)
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
		return b
	}
	// Try each possible value for this square.
	for _, c := range b[row][column] {
		value := int(c - '0')

		// Apply modifications to a duplicate board.
		newBoard := b.Duplicate()
		if err := newBoard.assign(row, column, value); err != nil {
			continue
		}

		// Try solving the board with this value.
		if solved := Solver(newBoard); solved != nil {
			return solved
		}
	}
	return nil
}
