package sudoku

func nextEmptySquare(b *Board) (int, int) {
	for i := 0; i < numRows; i++ {
		for j := 0; j < numColumns; j++ {
			if b[i][j] == EmptySquare {
				return i, j
			}
		}
	}
	return -1, -1
}

func Solver(b *Board) bool {
	row, column := nextEmptySquare(b)
	if row == -1 || column == -1 {
		return true
	}
	// try each value for this empty square
	for value := 1; value <= 9; value++ {
		err := b.SetValue(row, column, value)
		if err != nil {
			// value present in unit
			continue
		}
		// try solving the board with this value
		if Solver(b) {
			return true
		}
	}
	// It is important to reset square, as the board might be solved with different values in the caller functions.
	b.SetValue(row, column, EmptySquare)
	return false
}
