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

	allValues = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
)

type Board [numRows][numColumns][]int

func newEmptyBoard() *Board {
	board := &Board{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = make([]int, len(allValues))
			copy(board[i][j], allValues)
		}
	}
	return board
}

// NewBoard creates a sudoku grid from a string, filling empty squares that have only one possible value.
func NewBoard(str string) (*Board, error) {
	if str != "" && len(str) < numSquares {
		return nil, ErrInvalidBoardString
	}

	board := newEmptyBoard()
	if str == "" {
		return board, nil
	}

	numInserted := 0
	for _, c := range str {
		value := int(c - '0')
		if !isValidValue(value) && !isEmptyChar(c) {
			continue
		}
		if numInserted == numSquares {
			return nil, ErrInvalidBoardString
		}

		row := (numInserted / numColumns)
		column := (numInserted % numColumns)

		if !isEmptyChar(c) {
			err := board.assign(row, column, value)
			if err != nil {
				return nil, err
			}
		}
		numInserted++
	}

	if numInserted != numSquares {
		return nil, ErrInvalidBoardString
	}
	return board, nil
}

func DuplicateBoard(old *Board) *Board {
	board := &Board{}
	for i := 0; i < numRows; i++ {
		for j := 0; j < numColumns; j++ {
			board[i][j] = make([]int, len(old[i][j]))
			copy(board[i][j], old[i][j])
		}
	}
	return board
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

	b[row][column] = []int{value}
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

func (b *Board) eliminateInBox(row, column, value int) bool {
	boxRow := (row / 3) * 3
	boxColumn := (column / 3) * 3

	for i := boxRow; i < boxRow+3; i++ {
		for j := boxColumn; j < boxColumn+3; j++ {
			if i == row && j == column {
				continue
			}
			if len(b[i][j]) == 1 && b[i][j][0] == value {
				if b[i][j][0] == value {
					return false
				}
				continue
			}
			for k, val := range b[i][j] {
				if val == value {
					b[i][j][k] = b[i][j][len(b[i][j])-1]
					b[i][j] = b[i][j][:len(b[i][j])-1]
					if !b.eliminate(i, j) {
						return false
					}
					break
				}
			}
		}
	}
	return true
}

func (b *Board) eliminateInColumn(row, column, value int) bool {
	for i := 0; i < numRows; i++ {
		if i == row {
			continue
		}
		if len(b[i][column]) == 1 {
			if b[i][column][0] == value {
				return false
			}
			continue
		}
		for j, val := range b[i][column] {
			if val == value {
				b[i][column][j] = b[i][column][len(b[i][column])-1]
				b[i][column] = b[i][column][:len(b[i][column])-1]
				if !b.eliminate(i, column) {
					return false
				}
				break
			}
		}
	}
	return true
}

func (b *Board) eliminateInRow(row, column, value int) bool {
	for i := 0; i < numColumns; i++ {
		if i == column {
			continue
		}
		if len(b[row][i]) == 1 {
			if b[row][i][0] == value {
				return false
			}
			continue
		}
		for j, val := range b[row][i] {
			if val == value {
				b[row][i][j] = b[row][i][len(b[row][i])-1]
				b[row][i] = b[row][i][:len(b[row][i])-1]
				if !b.eliminate(row, i) {
					return false
				}
				break
			}
		}
	}
	return true
}

func (b *Board) valuePossible(row, column int, value int) bool {
	for _, val := range b[row][column] {
		if val == value {
			return true
		}
	}
	return false
}

func (b *Board) GetValue(row, column int) (int, error) {
	if !isValidPosition(row, column) {
		return -1, ErrInvalidPosition
	}
	if len(b[row][column]) > 1 {
		return EmptySquare, nil
	}
	return b[row][column][0], nil
}

func (b *Board) String() string {
	var str strings.Builder
	for i := 0; i < numRows; i++ {
		for j := 0; j < numColumns; j++ {
			if len(b[i][j]) > 1 {
				str.WriteByte('.')
			} else {
				str.WriteByte(byte(b[i][j][0] + '0'))
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
