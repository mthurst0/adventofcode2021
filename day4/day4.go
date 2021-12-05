package day4

import (
	"advent/rkutil"
	"fmt"
	"strconv"
	"strings"
)

const dim = 5

type BingoBoard struct {
	Squares    [dim][dim]int
	Called     [dim][dim]bool
	WasAWinner bool
}

func (bb BingoBoard) IsRowTrue(rowIdx int) bool {
	for colIdx := 0; colIdx < dim; colIdx++ {
		if !bb.Called[rowIdx][colIdx] {
			return false
		}
	}
	return true
}

func (bb BingoBoard) IsColumnTrue(colIdx int) bool {
	for rowIdx := 0; rowIdx < dim; rowIdx++ {
		if !bb.Called[rowIdx][colIdx] {
			return false
		}
	}
	return true
}

func (bb *BingoBoard) IsWinner() bool {
	if bb.WasAWinner {
		return true
	}
	for i := 0; i < dim; i++ {
		if bb.IsRowTrue(i) {
			bb.WasAWinner = true
			return true
		}
		if bb.IsColumnTrue(i) {
			bb.WasAWinner = true
			return true
		}
	}
	return false
}

func (bb BingoBoard) SumUncalled() int {
	result := 0
	for rowIdx := 0; rowIdx < dim; rowIdx++ {
		for colIdx := 0; colIdx < dim; colIdx++ {
			if !bb.Called[rowIdx][colIdx] {
				result += bb.Squares[rowIdx][colIdx]
			}
		}
	}
	return result
}

type BingoGame struct {
	Calls        []int
	Boards       []*BingoBoard
	WinningCount int
	LastWinner   *BingoBoard
}

func (game *BingoGame) Call(n int) {
	for _, board := range game.Boards {
		for rowIdx := 0; rowIdx < dim; rowIdx++ {
			for colIdx := 0; colIdx < dim; colIdx++ {
				if board.Squares[rowIdx][colIdx] == n {
					board.Called[rowIdx][colIdx] = true
				}
			}
		}
	}
}

func (game *BingoGame) Winner() *BingoBoard {
	for _, board := range game.Boards {
		if board.WasAWinner {
			continue
		}
		if board.IsWinner() {
			game.LastWinner = board
			game.WinningCount++
			return board
		}
	}
	return nil
}

func (game *BingoGame) CalcWinners() {
	for {
		if game.Winner() == nil {
			break
		}
	}
}

func rowValues(s string) [dim]int {
	var result [dim]int
	fs := strings.Fields(s)
	for i, f := range fs {
		v, err := strconv.Atoi(f)
		if err != nil {
			rkutil.Unexpected(fmt.Errorf("expected integer conversion failure: %w", err))
		}
		if i >= dim {
			rkutil.Unexpected(fmt.Errorf("unexpectedly large number of column: %d", i))
		}
		result[i] = v
	}
	return result
}

func parseBingo(filename string) *BingoGame {
	lines := rkutil.MustLines(filename)
	game := BingoGame{}
	game.Calls = rkutil.MustCommaDelimitedNumbers(lines[0])
	curBoard := &BingoBoard{}
	rowIdx := 0
	for _, line := range lines[2:] {
		if line == "" {
			game.Boards = append(game.Boards, curBoard)
			curBoard = &BingoBoard{}
			rowIdx = 0
		} else {
			curBoard.Squares[rowIdx] = rowValues(line)
			rowIdx++
		}
	}
	game.Boards = append(game.Boards, curBoard)
	return &game
}

func SolveEasier() string {
	game := parseBingo("day4/input.txt")
	for _, v := range game.Calls {
		game.Call(v)
		winner := game.Winner()
		if winner != nil {
			uncalled := winner.SumUncalled()
			return fmt.Sprintf("%d", v*uncalled)
		}
	}
	rkutil.Unexpected(fmt.Errorf("did not expect to arrive here"))
	return ""
}

func SolveHarder() string {
	game := parseBingo("day4/input.txt")
	for _, v := range game.Calls {
		game.Call(v)
		game.CalcWinners()
		if game.WinningCount == len(game.Boards) {
			uncalled := game.LastWinner.SumUncalled()
			return fmt.Sprintf("%d", v*uncalled)
		}
	}
	rkutil.Unexpected(fmt.Errorf("did not expect to arrive here"))
	return ""
}
