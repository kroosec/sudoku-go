package sudoku_test

import (
	"testing"

	"github.com/kroosec/sudoku-go"
)

func TestBoard(t *testing.T) {
	t.Run("get a square's value and possible values count", func(t *testing.T) {
		board, err := sudoku.NewBoard("")
		assertError(t, err, nil)

		value, err := board.GetValue(0, 0)
		assertError(t, err, nil)
		assertSquareValue(t, value, sudoku.EmptySquare, board)

		value, err = board.CountPossible(0, 0)
		assertError(t, err, nil)
		assertSquareValue(t, value, 9, board)

		_, err = board.GetValue(0, -1)
		assertError(t, err, sudoku.ErrInvalidPosition)

		_, err = board.CountPossible(0, -1)
		assertError(t, err, sudoku.ErrInvalidPosition)
	})

	t.Run("set value and verify that it was updated", func(t *testing.T) {
		type testCase struct {
			name   string
			row    int
			column int
			value  int
			err    error
		}

		cases := []testCase{
			{"Valid #1", 0, 2, 5, nil},
			{"Valid #2", 7, 0, 1, nil},
			{"Invalid value #1", 7, 0, 10, sudoku.ErrInvalidValue},
			{"Invalid value #2", 7, 0, -1, sudoku.ErrInvalidValue},
			{"Invalid row #1", 9, 0, 4, sudoku.ErrInvalidPosition},
			{"Invalid row #2", -1, 0, 2, sudoku.ErrInvalidPosition},
			{"Invalid column #1", 2, 10, 3, sudoku.ErrInvalidPosition},
			{"Invalid column #2", 2, -2, 5, sudoku.ErrInvalidPosition},
		}

		board, err := sudoku.NewBoard("")
		assertError(t, err, nil)

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {

				err := board.SetValue(c.row, c.column, c.value)
				assertError(t, err, c.err)
				if c.err == nil {
					value, err := board.GetValue(c.row, c.column)
					assertError(t, err, nil)
					assertSquareValue(t, value, c.value, board)
				}
			})
		}
	})

	t.Run("export board to string, duplicate produces the same", func(t *testing.T) {
		want := "4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......"

		board, err := sudoku.NewBoard(want)
		assertError(t, err, nil)
		assertBoardString(t, board.String(), want)

		board = board.Duplicate()
		assertBoardString(t, board.String(), want)

	})

	t.Run("New board from strings", func(t *testing.T) {
		type testSquare struct {
			row    int
			column int
			value  int
		}

		type testCase struct {
			name        string
			boardString string
			squares     []testSquare
			err         error
		}

		cases := []testCase{
			{name: "Valid board #1",
				boardString: "4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......",
				err:         nil,
				squares: []testSquare{
					{0, 0, 4},
					{0, 6, 8},
					{0, 8, 5},
					{1, 1, 3},
					{2, 3, 7},
					{8, 0, 1},
					{8, 2, 4},
				},
			},
			{name: "Valid board #2",
				boardString: `
4 . . |. . . |8 . 5
. 3 . |. . . |. . .
. . . |7 . . |. . .
------+------+------
. 2 . |. . . |. 6 .
. . . |. 8 . |4 . .
. . . |. 1 . |. . .
------+------+------
. . . |6 . 3 |. 7 .
5 . . |2 . . |. . .
1 . 4 |. . . |. . .
`,
				squares: []testSquare{
					{0, 6, 8},
					{1, 8, 0},
					{2, 3, 7},
					{3, 8, 0},
					{4, 8, 0},
					{5, 6, 0},
					{6, 8, 0},
					{7, 3, 2},
					{8, 2, 4},
				},
				err: nil,
			},
			{name: "Invalid board #1: Few values",
				boardString: "1234567",
				err:         sudoku.ErrInvalidBoardString,
				squares:     nil,
			},
			{name: "Invalid board #2: Too many values",
				boardString: "00302060090030500100180640000810290070000000800670820000260950080020300900501030081370",
				err:         sudoku.ErrInvalidBoardString,
				squares:     nil,
			},
			{name: "Invalid board #3: Less than 81 valid values",
				boardString: "ab..3.2.6..9..3.5..1..18.64....81.29..7.......8..67.82....26.95..8..2.3..9..5.1.3",
				err:         sudoku.ErrInvalidBoardString,
				squares:     nil,
			},
			{name: "Invalid board #4: Same values in a column",
				boardString: "3........3.......................................................................",
				err:         sudoku.ErrDuplicateValue,
			},
			{name: "Invalid board #5: Same values in a row",
				boardString: "33...............................................................................",
				err:         sudoku.ErrDuplicateValue,
			},
			{name: "Invalid board #6: Same values in a box",
				boardString: ".3.........3.....................................................................",
				err:         sudoku.ErrDuplicateValue,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				board, err := sudoku.NewBoard(c.boardString)
				assertError(t, err, c.err)

				for _, square := range c.squares {
					value, err := board.GetValue(square.row, square.column)
					assertError(t, err, nil)
					assertSquareValue(t, value, square.value, board)
				}
			})
		}
	})
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Fatalf("expected error %v, got %v", want, got)
	}
}

func assertSquareValue(t *testing.T, got, want int, board *sudoku.Board) {
	t.Helper()

	if got != want {
		t.Fatalf("expected square value %d, got %d: %+v", want, got, board)
	}
}

func assertBoardString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("expected board string '%s', got '%s'", want, got)
	}
}
