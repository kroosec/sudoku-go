package sudoku

import (
	"fmt"
	"strings"
)

const (
	EmptySquare = 0

	numRows    = 9
	numColumns = 9
	numSquares = numRows * numColumns
)

var (
	ErrInvalidBoardString = fmt.Errorf("Invalid board string")
	ErrInvalidPosition    = fmt.Errorf("Invalid position")
	ErrInvalidValue       = fmt.Errorf("Invalid value")
	ErrDuplicateValue     = fmt.Errorf("Value already exists in unit")

	allValues = "123456789"
)

type Board [numRows][numColumns]string

func newEmptyBoard() *Board {
	board := &Board{}

	for i := 0; i < numRows; i++ {
		for j := 0; j < numColumns; j++ {
			board[i][j] = allValues
		}
	}

	return board
}

func (b *Board) insertValues(str string) error {
	if str == "" {
		return nil
	}

	numInserted := 0
	for _, c := range str {
		if isEmptyChar(c) {
			numInserted++
			continue
		}

		value := int(c - '0')
		if !isValidValue(value) {
			continue
		}
		if numInserted == numSquares {
			return ErrInvalidBoardString
		}

		row := (numInserted / numColumns)
		column := (numInserted % numColumns)

		err := b.assign(row, column, value)
		if err != nil {
			return err
		}
		numInserted++
	}

	if numInserted != numSquares {
		return ErrInvalidBoardString
	}
	return nil
}

// NewBoard creates a sudoku grid from a string, filling empty squares that have only one possible value.
func NewBoard(str string) (*Board, error) {
	board := newEmptyBoard()

	if err := board.insertValues(str); err != nil {
		return nil, err
	}

	return board, nil
}

func (b *Board) Duplicate() *Board {
	newBoard := &Board{}
	for i := 0; i < numRows; i++ {
		for j := 0; j < numColumns; j++ {
			newBoard[i][j] = b[i][j]
		}
	}
	return newBoard
}

func (b *Board) eliminate(row, column int) bool {
	switch len(b[row][column]) {
	case 1:
		value := b[row][column][0]
		return b.eliminateInRow(row, column, value) && b.eliminateInColumn(row, column, value) && b.eliminateInBox(row, column, value)
	case 0:
		return false
	default:
		return true
	}
}

func (b *Board) assign(row, column, value int) error {
	if !b.valuePossible(row, column, value) {
		return ErrDuplicateValue
	}

	b[row][column] = string(byte(value + '0'))
	if !b.eliminate(row, column) {
		return ErrDuplicateValue
	}
	return nil
}

func (b *Board) SetValue(row, column, value int) error {
	if !isValidPosition(row, column) {
		return ErrInvalidPosition
	}
	if !isValidValue(value) {
		return ErrInvalidValue
	}
	return b.assign(row, column, value)
}

func (b *Board) CountPossible(row, column int) (int, error) {
	if !isValidPosition(row, column) {
		return 0, ErrInvalidPosition
	}

	return len(b[row][column]), nil
}

func (b *Board) eliminateInBox(row, column int, value byte) bool {
	boxRow := (row / 3) * 3
	boxColumn := (column / 3) * 3

	for i := boxRow; i < boxRow+3; i++ {
		for j := boxColumn; j < boxColumn+3; j++ {
			if i == row && j == column {
				continue
			}

			if !b.eliminateSquare(i, j, value) {
				return false
			}
		}
	}
	return true
}

func (b *Board) eliminateInColumn(row, column int, value byte) bool {
	for i := 0; i < numRows; i++ {
		if i == row {
			continue
		}
		if !b.eliminateSquare(i, column, value) {
			return false
		}
	}
	return true
}

func (b *Board) eliminateInRow(row, column int, value byte) bool {
	for i := 0; i < numColumns; i++ {
		if i == column {
			continue
		}
		if !b.eliminateSquare(row, i, value) {
			return false
		}
	}
	return true
}

func (b *Board) eliminateSquare(row, column int, value byte) bool {
	if strings.ContainsRune(b[row][column], rune(value)) {
		b[row][column] = strings.ReplaceAll(b[row][column], string(value), "")

		if !b.eliminate(row, column) {
			return false
		}
	}
	return true
}

func (b *Board) valuePossible(row, column int, value int) bool {
	return strings.ContainsRune(b[row][column], rune(value+'0'))
}

func (b *Board) GetValue(row, column int) (int, error) {
	if !isValidPosition(row, column) {
		return -1, ErrInvalidPosition
	}

	if len(b[row][column]) > 1 {
		return EmptySquare, nil
	}
	return int(b[row][column][0] - '0'), nil
}

func (b *Board) String() string {
	var str strings.Builder

	for i := 0; i < numRows; i++ {
		for j := 0; j < numColumns; j++ {
			if len(b[i][j]) > 1 {
				str.WriteByte('.')
			} else {
				str.WriteString(b[i][j])
			}
		}
	}
	return str.String()
}

func isEmptyChar(c rune) bool {
	return c == '.' || c == '0'
}

func isValidPosition(row, column int) bool {
	if row < 0 || row >= numRows || column < 0 || column >= numColumns {
		return false
	}
	return true
}

func isValidValue(value int) bool {
	return value >= 1 && value <= 9
}
