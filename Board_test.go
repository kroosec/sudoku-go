package sudoku_test

import (
	"testing"

	"sudoku"
)

func TestBoard(t *testing.T) {
	t.Run("get a square's value", func(t *testing.T) {
		board, err := sudoku.NewBoard("")
		assertError(t, err, nil)

		value, err := board.GetValue(0, 0)
		assertError(t, err, nil)
		assertSquareValue(t, value, sudoku.EmptySquare, board)
		_, err = board.GetValue(0, -1)
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

	t.Run("export board to string", func(t *testing.T) {
		want := "..3.2.6..9..3.5..1..18.64....81.29..7.......8..67.82....26.95..8..2.3..9..5.1.3.."
		board, err := sudoku.NewBoard(want)
		assertError(t, err, nil)
		got := board.String()
		if got != want {
			t.Errorf("expected board string '%s', got '%s'", want, got)
		}
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
				boardString: "003020600 900305001 001806400 008102900 700000008 006708200 002609500 800203009 005010300",
				err:         nil,
				squares: []testSquare{
					{0, 0, 0},
					{0, 1, 0},
					{0, 2, 3},
					{0, 3, 0},
					{0, 4, 2},
					{0, 5, 0},
					{0, 6, 6},
					{0, 7, 0},
					{0, 8, 0},

					{1, 0, 9},
					{1, 1, 0},
					{1, 2, 0},
					{1, 3, 3},
					{1, 4, 0},
					{1, 5, 5},
					{1, 6, 0},
					{1, 7, 0},
					{1, 8, 1},

					{2, 0, 0},
					{2, 1, 0},
					{2, 2, 1},
					{2, 3, 8},
					{2, 4, 0},
					{2, 5, 6},
					{2, 6, 4},
					{2, 7, 0},
					{2, 8, 0},

					{3, 0, 0},
					{3, 1, 0},
					{3, 2, 8},
					{3, 3, 1},
					{3, 4, 0},
					{3, 5, 2},
					{3, 6, 9},
					{3, 7, 0},
					{3, 8, 0},

					{4, 0, 7},
					{4, 1, 0},
					{4, 2, 0},
					{4, 3, 0},
					{4, 4, 0},
					{4, 5, 0},
					{4, 6, 0},
					{4, 7, 0},
					{4, 8, 8},

					{5, 0, 0},
					{5, 1, 0},
					{5, 2, 6},
					{5, 3, 7},
					{5, 4, 0},
					{5, 5, 8},
					{5, 6, 2},
					{5, 7, 0},
					{5, 8, 0},

					{6, 0, 0},
					{6, 1, 0},
					{6, 2, 2},
					{6, 3, 6},
					{6, 4, 0},
					{6, 5, 9},
					{6, 6, 5},
					{6, 7, 0},
					{6, 8, 0},

					{7, 0, 8},
					{7, 1, 0},
					{7, 2, 0},
					{7, 3, 2},
					{7, 4, 0},
					{7, 5, 3},
					{7, 6, 0},
					{7, 7, 0},
					{7, 8, 9},

					{8, 0, 0},
					{8, 1, 0},
					{8, 2, 5},
					{8, 3, 0},
					{8, 4, 1},
					{8, 5, 0},
					{8, 6, 3},
					{8, 7, 0},
					{8, 8, 0},
				},
			},
			{name: "Valid board #2",
				boardString: "..3.2.6..9..3.5..1..18.64....81.29..7.......8..67.82....26.95..8..2.3..9..5.1.3..",
				err:         nil,
				squares: []testSquare{
					{0, 6, 6},
					{1, 8, 1},
					{2, 3, 8},
					{3, 8, 0},
					{4, 8, 8},
					{5, 6, 2},
					{6, 8, 0},
					{7, 8, 9},
					{8, 6, 3},
				},
			},
			{name: "Valid board #3",
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
				boardString: "1238577",
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
			{name: "Invalid board #5: Same values in a box",
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
