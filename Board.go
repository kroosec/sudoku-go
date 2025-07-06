// Package sudoku provides functionality to solve Sudoku puzzles. It uses a
// backtracking search algorithm combined with constraint propagation.
package sudoku

import (
	"fmt"
	"strings"
)

const (
	EmptySquare = 0

	numDigits  = 9
	numRows    = 9
	numColumns = 9
	numSquares = numRows * numColumns
)

var (
	ErrInvalidBoardString = fmt.Errorf("invalid board string")
	ErrInvalidPosition    = fmt.Errorf("invalid position")
	ErrInvalidValue       = fmt.Errorf("invalid value")
	ErrDuplicateValue     = fmt.Errorf("value already exists in unit")

	allValues = "123456789"

	// peers is a map where the key is a square's coordinates (row, column),
	// and the value is a slice of coordinates of its peers.
	peers map[[2]int][][2]int
)

func init() {
	peers = make(map[[2]int][][2]int)
	for r := range numRows {
		for c := range numColumns {
			key := [2]int{r, c}
			peerSet := make(map[[2]int]struct{})

			// Row peers
			for i := range numColumns {
				peerSet[[2]int{r, i}] = struct{}{}
			}

			// Column peers
			for i := range numRows {
				peerSet[[2]int{i, c}] = struct{}{}
			}

			// Box peers
			boxRowStart := (r / 3) * 3
			boxColStart := (c / 3) * 3
			for i := boxRowStart; i < boxRowStart+3; i++ {
				for j := boxColStart; j < boxColStart+3; j++ {
					peerSet[[2]int{i, j}] = struct{}{}
				}
			}

			delete(peerSet, key)

			peerList := make([][2]int, 0, len(peerSet))
			for p := range peerSet {
				peerList = append(peerList, p)
			}
			peers[key] = peerList
		}
	}
}

type Board [numRows][numColumns]string

func newEmptyBoard() *Board {
	board := &Board{}

	for i := range numRows {
		for j := range numColumns {
			board[i][j] = allValues
		}
	}

	return board
}

func (b *Board) insertValues(str string) error {
	if str == "" {
		return nil
	}

	// Sanitize input string. Keep relevant characters only
	var cleanStr strings.Builder
	for _, c := range str {
		if (c >= '1' && c <= '9') || c == '.' || c == '0' {
			cleanStr.WriteRune(c)
		}
	}

	if cleanStr.Len() != numSquares {
		return ErrInvalidBoardString
	}

	for i, c := range cleanStr.String() {
		if isEmptyChar(c) {
			continue
		}

		value := int(c - '0')
		row := i / numColumns
		column := i % numColumns

		if err := b.assign(row, column, value); err != nil {
			return err
		}
	}
	return nil
}

func (b *Board) eliminate(row, column int) bool {
	switch len(b[row][column]) {
	case 0:
		// Contradiction: removed last value.
		return false
	case 1:
		// If a square is reduced to one value, then eliminate it from its peers.
		value := b[row][column][0]
		for _, p := range peers[[2]int{row, column}] {
			if !b.eliminateSquare(p[0], p[1], value) {
				return false
			}
		}
	}
	return true
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
	for i := range numRows {
		for j := range numColumns {
			newBoard[i][j] = b[i][j]
		}
	}
	return newBoard
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

	for i := range numRows {
		for j := range numColumns {
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
	return value >= 1 && value <= numDigits
}
