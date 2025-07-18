package sudoku_test

import (
	"testing"

	"github.com/kroosec/sudoku-go"
	"github.com/stretchr/testify/assert"
)

// Easy problems.
// Source: Project Euler https://projecteuler.net/project/resources/p096_sudoku.txt
var easyProblems = []string{
	"003020600 900305001 001806400 008102900 700000008 006708200 002609500 800203009 005010300",
	"200080300 060070084 030500209 000105408 000000000 402706000 301007040 720040060 004010003",
	"000000907 000420180 000705026 100904000 050000040 000507009 920108000 034059000 507000000",
	"030050040 008010500 460000012 070502080 000603000 040109030 250000098 001020600 080060020",
	"020810740 700003100 090002805 009040087 400208003 160030200 302700060 005600008 076051090",
	"100920000 524010000 000000070 050008102 000000000 402700090 060000000 000030945 000071006",
	"043080250 600000000 000001094 900004070 000608000 010200003 820500000 000000005 034090710",
	"480006902 002008001 900370060 840010200 003704100 001060049 020085007 700900600 609200018",
	"000900002 050123400 030000160 908000000 070000090 000000205 091000050 007439020 400007000",
	"001900003 900700160 030005007 050000009 004302600 200000070 600100030 042007006 500006800",
	"000125400 008400000 420800000 030000095 060902010 510000060 000003049 000007200 001298000",
	"062340750 100005600 570000040 000094800 400000006 005830000 030000091 006400007 059083260",
	"300000000 005009000 200504000 020000700 160000058 704310600 000890100 000067080 000005437",
	"630000000 000500008 005674000 000020000 003401020 000000345 000007004 080300902 947100080",
	"000020040 008035000 000070602 031046970 200000000 000501203 049000730 000000010 800004000",
	"361025900 080960010 400000057 008000471 000603000 259000800 740000005 020018060 005470329",
	"050807020 600010090 702540006 070020301 504000908 103080070 900076205 060090003 080103040",
	"080005000 000003457 000070809 060400903 007010500 408007020 901020000 842300000 000100080",
	"003502900 000040000 106000305 900251008 070408030 800763001 308000104 000020000 005104800",
	"000000000 009805100 051907420 290401065 000000000 140508093 026709580 005103600 000000000",
	"020030090 000907000 900208005 004806500 607000208 003102900 800605007 000309000 030020050",
	"005000006 070009020 000500107 804150000 000803000 000092805 907006000 030400010 200000600",
	"040000050 001943600 009000300 600050002 103000506 800020007 005000200 002436700 030000040",
	"004000000 000030002 390700080 400009001 209801307 600200008 010008053 900040000 000000800",
	"360020089 000361000 000000000 803000602 400603007 607000108 000000000 000418000 970030014",
	"500400060 009000800 640020000 000001008 208000501 700500000 000090084 003000600 060003002",
	"007256400 400000005 010030060 000508000 008060200 000107000 030070090 200000004 006312700",
	"000000000 079050180 800000007 007306800 450708096 003502700 700000005 016030420 000000000",
	"030000080 009000500 007509200 700105008 020090030 900402001 004207100 002000800 070000090",
	"200170603 050000100 000006079 000040700 000801000 009050000 310400000 005000060 906037002",
	"000000080 800701040 040020030 374000900 000030000 005000321 010060050 050802006 080000000",
	"000000085 000210009 960080100 500800016 000000000 890006007 009070052 300054000 480000000",
	"608070502 050608070 002000300 500090006 040302050 800050003 005000200 010704090 409060701",
	"050010040 107000602 000905000 208030501 040070020 901080406 000401000 304000709 020060010",
	"053000790 009753400 100000002 090080010 000907000 080030070 500000003 007641200 061000940",
	"006080300 049070250 000405000 600317004 007000800 100826009 000702000 075040190 003090600",
	"005080700 700204005 320000084 060105040 008000500 070803010 450000091 600508007 003010600",
	"000900800 128006400 070800060 800430007 500000009 600079008 090004010 003600284 001007000",
	"000080000 270000054 095000810 009806400 020403060 006905100 017000620 460000038 000090000",
	"000602000 400050001 085010620 038206710 000000000 019407350 026040530 900020007 000809000",
	"380000000 000400785 009020300 060090000 800302009 000040070 001070500 495006000 000000092",
	"000158000 002060800 030000040 027030510 000000000 046080790 050000080 004070100 000325000",
	"010500200 900001000 002008030 500030007 008000500 600080004 040100700 000700006 003004050",
	"080000040 000469000 400000007 005904600 070608030 008502100 900000005 000781000 060000010",
	"904200007 010000000 000706500 000800090 020904060 040002000 001607000 000000030 300005702",
	"000700800 006000031 040002000 024070000 010030080 000060290 000800070 860000500 002006000",
	"001007090 590080001 030000080 000005800 050060020 004100000 080000030 100020079 020700400",
	"000003017 015009008 060000000 100007000 009000200 000500004 000000020 500600340 340200000",
	"300200000 000107000 706030500 070009080 900020004 010800050 009040301 000702000 000008006",
}

// Hard problems.
// Source: magictour.free.fr/top95
var hardProblems = []string{
	"4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......",
	"52...6.........7.13...........4..8..6......5...........418.........3..2...87.....",
	"6.....8.3.4.7.................5.4.7.3..2.....1.6.......2.....5.....8.6......1....",
	"48.3............71.2.......7.5....6....2..8.............1.76...3.....4......5....",
	"....14....3....2...7..........9...3.6.1.............8.2.....1.4....5.6.....7.8...",
	"......52..8.4......3...9...5.1...6..2..7........3.....6...1..........7.4.......3.",
	"6.2.5.........3.4..........43...8....1....2........7..5..27...........81...6.....",
	".524.........7.1..............8.2...3.....6...9.5.....1.6.3...........897........",
	".923.........8.1...........1.7.4...........658.........6.5.2...4.....7.....9.....",
	"6..3.2....5.....1..........7.26............543.........8.15........4.2........7..",
	".6.5.1.9.1...9..539....7....4.8...7.......5.8.817.5.3.....5.2............76..8...",
	"..5...987.4..5...1..7......2...48....9.1.....6..2.....3..6..2.......9.7.......5..",
	"3.6.7...........518.........1.4.5...7.....6.....2......2.....4.....8.3.....5.....",
	"1.....3.8.7.4..............2.3.1...........958.........5.6...7.....8.2...4.......",
	"6..3.2....4.....1..........7.26............543.........8.15........4.2........7..",
	"....3..9....2....1.5.9..............1.2.8.4.6.8.5...2..75......4.1..6..3.....4.6.",
	"45.....3....8.1....9...........5..9.2..7.....8.........1..4..........7.2...6..8..",
	".237....68...6.59.9.....7......4.97.3.7.96..2.........5..47.........2....8.......",
	"..84...3....3.....9....157479...8........7..514.....2...9.6...2.5....4......9..56",
	".98.1....2......6.............3.2.5..84.........6.........4.8.93..5...........1..",
	"..247..58..............1.4.....2...9528.9.4....9...1.........3.3....75..685..2...",
	"4.....8.5.3..........7......2.....6.....5.4......1.......6.3.7.5..2.....1.9......",
	".2.3......63.....58.......15....9.3....7........1....8.879..26......6.7...6..7..4",
	"1.....7.9.4...72..8.........7..1..6.3.......5.6..4..2.........8..53...7.7.2....46",
	"4.....3.....8.2......7........1...8734.......6........5...6........1.4...82......",
	".......71.2.8........4.3...7...6..5....2..3..9........6...7.....8....4......5....",
	"6..3.2....4.....8..........7.26............543.........8.15........8.2........7..",
	".47.8...1............6..7..6....357......5....1..6....28..4.....9.1...4.....2.69.",
	"......8.17..2........5.6......7...5..1....3...8.......5......2..4..8....6...3....",
	"38.6.......9.......2..3.51......5....3..1..6....4......17.5..8.......9.......7.32",
	"...5...........5.697.....2...48.2...25.1...3..8..3.........4.7..13.5..9..2...31..",
	".2.......3.5.62..9.68...3...5..........64.8.2..47..9....3.....1.....6...17.43....",
	".8..4....3......1........2...5...4.69..1..8..2...........3.9....6....5.....2.....",
	"..8.9.1...6.5...2......6....3.1.7.5.........9..4...3...5....2...7...3.8.2..7....4",
	"4.....5.8.3..........7......2.....6.....5.8......1.......6.3.7.5..2.....1.8......",
	"1.....3.8.6.4..............2.3.1...........958.........5.6...7.....8.2...4.......",
	"1....6.8..64..........4...7....9.6...7.4..5..5...7.1...5....32.3....8...4........",
	"249.6...3.3....2..8.......5.....6......2......1..4.82..9.5..7....4.....1.7...3...",
	"...8....9.873...4.6..7.......85..97...........43..75.......3....3...145.4....2..1",
	"...5.1....9....8...6.......4.1..........7..9........3.8.....1.5...2..4.....36....",
	"......8.16..2........7.5......6...2..1....3...8.......2......7..3..8....5...4....",
	".476...5.8.3.....2.....9......8.5..6...1.....6.24......78...51...6....4..9...4..7",
	".....7.95.....1...86..2.....2..73..85......6...3..49..3.5...41724................",
	".4.5.....8...9..3..76.2.....146..........9..7.....36....1..4.5..6......3..71..2..",
	".834.........7..5...........4.1.8..........27...3.....2.6.5....5.....8........1..",
	"..9.....3.....9...7.....5.6..65..4.....3......28......3..75.6..6...........12.3.8",
	".26.39......6....19.....7.......4..9.5....2....85.....3..2..9..4....762.........4",
	"2.3.8....8..7...........1...6.5.7...4......3....1............82.5....6...1.......",
	"6..3.2....1.....5..........7.26............843.........8.15........8.2........7..",
	"1.....9...64..1.7..7..4.......3.....3.89..5....7....2.....6.7.9.....4.1....129.3.",
	".........9......84.623...5....6...453...1...6...9...7....1.....4.5..2....3.8....9",
	".2....5938..5..46.94..6...8..2.3.....6..8.73.7..2.........4.38..7....6..........5",
	"9.4..5...25.6..1..31......8.7...9...4..26......147....7.......2...3..8.6.4.....9.",
	"...52.....9...3..4......7...1.....4..8..453..6...1...87.2........8....32.4..8..1.",
	"53..2.9...24.3..5...9..........1.827...7.........981.............64....91.2.5.43.",
	"1....786...7..8.1.8..2....9........24...1......9..5...6.8..........5.9.......93.4",
	"....5...11......7..6.....8......4.....9.1.3.....596.2..8..62..7..7......3.5.7.2..",
	".47.2....8....1....3....9.2.....5...6..81..5.....4.....7....3.4...9...1.4..27.8..",
	"......94.....9...53....5.7..8.4..1..463...........7.8.8..7.....7......28.5.26....",
	".2......6....41.....78....1......7....37.....6..412....1..74..5..8.5..7......39..",
	"1.....3.8.6.4..............2.3.1...........758.........7.5...6.....8.2...4.......",
	"2....1.9..1..3.7..9..8...2.......85..6.4.........7...3.2.3...6....5.....1.9...2.5",
	"..7..8.....6.2.3...3......9.1..5..6.....1.....7.9....2........4.83..4...26....51.",
	"...36....85.......9.4..8........68.........17..9..45...1.5...6.4....9..2.....3...",
	"34.6.......7.......2..8.57......5....7..1..2....4......36.2..1.......9.......7.82",
	"......4.18..2........6.7......8...6..4....3...1.......6......2..5..1....7...3....",
	".4..5..67...1...4....2.....1..8..3........2...6...........4..5.3.....8..2........",
	".......4...2..4..1.7..5..9...3..7....4..6....6..1..8...2....1..85.9...6.....8...3",
	"8..7....4.5....6............3.97...8....43..5....2.9....6......2...6...7.71..83.2",
	".8...4.5....7..3............1..85...6.....2......4....3.26............417........",
	"....7..8...6...5...2...3.61.1...7..2..8..534.2..9.......2......58...6.3.4...1....",
	"......8.16..2........7.5......6...2..1....3...8.......2......7..4..8....5...3....",
	".2..........6....3.74.8.........3..2.8..4..1.6..5.........1.78.5....9..........4.",
	".52..68.......7.2.......6....48..9..2..41......1.....8..61..38.....9...63..6..1.9",
	"....1.78.5....9..........4..2..........6....3.74.8.........3..2.8..4..1.6..5.....",
	"1.......3.6.3..7...7...5..121.7...9...7........8.1..2....8.64....9.2..6....4.....",
	"4...7.1....19.46.5.....1......7....2..2.3....847..6....14...8.6.2....3..6...9....",
	"......8.17..2........5.6......7...5..1....3...8.......5......2..3..8....6...4....",
	"963......1....8......2.5....4.8......1....7......3..257......3...9.2.4.7......9..",
	"15.3......7..4.2....4.72.....8.........9..1.8.1..8.79......38...........6....7423",
	"..........5724...98....947...9..3...5..9..12...3.1.9...6....25....56.....7......6",
	"....75....1..2.....4...3...5.....3.2...8...1.......6.....1..48.2........7........",
	"6.....7.3.4.8.................5.4.8.7..2.....1.3.......2.....5.....7.9......1....",
	"....6...4..6.3....1..4..5.77.....8.5...8.....6.8....9...2.9....4....32....97..1..",
	".32.....58..3.....9.428...1...4...39...6...5.....1.....2...67.8.....4....95....6.",
	"...5.3.......6.7..5.8....1636..2.......4.1.......3...567....2.8..4.7.......2..5..",
	".5.3.7.4.1.........3.......5.8.3.61....8..5.9.6..1........4...6...6927....2...9..",
	"..5..8..18......9.......78....4.....64....9......53..2.6.........138..5....9.714.",
	"..........72.6.1....51...82.8...13..4.........37.9..1.....238..5.4..9.........79.",
	"...658.....4......12............96.7...3..5....2.8...3..19..8..3.6.....4....473..",
	".2.3.......6..8.9.83.5........2...8.7.9..5........6..4.......1...1...4.22..7..8.9",
	".5..9....1.....6.....3.8.....8.4...9514.......3....2..........4.8...6..77..15..6.",
	".....2.......7...17..3...9.8..7......2.89.6...13..6....9..5.824.....891..........",
	"3...8.......7....51..............36...2..4....7...........6.13..452...........8..",
}

func TestSolver(t *testing.T) {
	t.Run("solve easy problems", func(t *testing.T) {
		for _, boardString := range easyProblems {
			assertProblem(t, boardString)
		}
	})

	t.Run("solve hard problems", func(t *testing.T) {
		for _, boardString := range hardProblems {
			assertProblem(t, boardString)
		}
	})
}

func assertProblem(t *testing.T, boardString string) {
	t.Helper()

	t.Run(boardString, func(t *testing.T) {
		board, err := sudoku.NewBoard(boardString)
		assert.NoError(t, err)

		solved := sudoku.Solver(board)
		assert.NotNil(t, solved)

		for i := range 9 {
			for j := range 9 {
				value, err := solved.GetValue(i, j)
				assert.NoError(t, err)
				assert.NotEqual(t, sudoku.EmptySquare, value)
			}
		}

		// Can import solved board, ie. check for erroneous solutions.
		_, err = sudoku.NewBoard(solved.String())
		assert.NoError(t, err)
	})
}
