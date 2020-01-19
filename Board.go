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
)

type Board [numRows][numColumns]int

type Square struct {
	row    int
	column int
}

func NewBoard(str string) (*Board, error) {
	if str != "" && len(str) < numSquares {
		return nil, ErrInvalidBoardString
	}
	board := &Board{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = EmptySquare
		}
	}
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

		if isEmptyChar(c) {
			value = EmptySquare
		}
		err := board.SetValue(row, column, value)
		if err != nil {
			return nil, err
		}
		numInserted++
	}
	if numInserted != numSquares {
		return nil, ErrInvalidBoardString
	}
	return board, nil
}

func (b *Board) SetValue(row, column, value int) error {
	if !isValidPosition(row, column) {
		return ErrInvalidPosition
	}
	if !isValidValue(value) && value != EmptySquare {
		return ErrInvalidValue
	}
	if b.valueExists(row, column, value) {
		return ErrDuplicateValue
	}
	b[row][column] = value
	return nil
}

func getPeers(row, column int) []Square {
	var peers []Square
	for i := 0; i < numColumns; i++ {
		if i != column {
			peers = append(peers, Square{row, i})
		}
		if i != row {
			peers = append(peers, Square{i, column})
		}
	}
	startRow := ((row) / 3) * 3
	startColumn := ((column) / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startColumn; j < startColumn+3; j++ {
			if i != row && j != column {
				peers = append(peers, Square{i, j})
			}
		}
	}
	return peers
}

func (b *Board) valueExists(row, column int, value int) bool {
	if value == EmptySquare {
		return false
	}
	peers := getPeers(row, column)
	for _, square := range peers {
		if b[square.row][square.column] == value {
			return true
		}
	}
	return false
}

func (b *Board) GetValue(row, column int) (int, error) {
	if !isValidPosition(row, column) {
		return -1, ErrInvalidPosition
	}
	return b[row][column], nil
}

func (b *Board) String() string {
	var str strings.Builder
	for i := 0; i < numRows; i++ {
		for j := 0; j < numColumns; j++ {
			if b[i][j] == EmptySquare {
				str.WriteByte('.')
			} else {
				str.WriteByte(byte(b[i][j] + '0'))
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
