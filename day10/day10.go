package day10

import (
	"advent/rkutil"
	"fmt"
	"sort"
	"strings"
)

type Bracket int

const (
	BracketParen = Bracket(iota)
	BracketSquare
	BracketCurly
	BracketPointy
)

func score(b Bracket) int {
	switch b {
	case BracketParen:
		return 3
	case BracketSquare:
		return 57
	case BracketCurly:
		return 1197
	case BracketPointy:
		return 25137
	}
	rkutil.UnexpectedCodePath()
	return 0
}

// [({(<(())[]>[[{[]{<()<>>

func bracketOf(c rune) Bracket {
	switch c {
	case '(', ')':
		return BracketParen
	case '[', ']':
		return BracketSquare
	case '{', '}':
		return BracketCurly
	case '<', '>':
		return BracketPointy
	}
	rkutil.UnexpectedCodePath()
	return 0
}

func scoreLine(line string) int {
	var bracketStack []Bracket
	for _, ch := range line {
		b := bracketOf(ch)
		if ch == '(' || ch == '[' || ch == '{' || ch == '<' {
			bracketStack = append(bracketStack, b)
		} else if ch == ')' || ch == ']' || ch == '}' || ch == '>' {
			if bracketStack[len(bracketStack)-1] != b {
				return score(b)
			}
			bracketStack = bracketStack[0 : len(bracketStack)-1]
		}
	}
	return 0
}

func solveEasier(filename string) int {
	lines := rkutil.MustLines(filename)
	result := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		s := scoreLine(line)
		if s > 0 {
			fmt.Println(line)
			result += s
		}
	}
	return result
}

func bracketCompleteScore(b Bracket) int {
	switch b {
	case BracketParen:
		return 1
	case BracketSquare:
		return 2
	case BracketCurly:
		return 3
	case BracketPointy:
		return 4
	}
	rkutil.UnexpectedCodePath()
	return 0
}

func SolveEasierSample() int {
	return solveEasier("sample_input.txt")
}

func SolveEasier() int {
	return solveEasier("input.txt")
}

func completeLine(line string) int {
	var bracketStack []Bracket
	for _, ch := range line {
		b := bracketOf(ch)
		if ch == '(' || ch == '[' || ch == '{' || ch == '<' {
			bracketStack = append(bracketStack, b)
		} else if ch == ')' || ch == ']' || ch == '}' || ch == '>' {
			if bracketStack[len(bracketStack)-1] != b {
				rkutil.UnexpectedCodePath()
				return 0
			}
			bracketStack = bracketStack[0 : len(bracketStack)-1]
		}
	}
	lineScore := 0
	for i := len(bracketStack) - 1; i >= 0; i-- {
		lineScore = (lineScore * 5) + bracketCompleteScore(bracketStack[i])
	}
	return lineScore
}

func solveHarder(filename string) int {
	lines := rkutil.MustLines(filename)
	var results []int
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if scoreLine(line) > 0 {
			continue
		}
		results = append(results, completeLine(line))
	}
	sort.Ints(results)
	return results[len(results)/2]
}

func SolveHarderSample() int {
	return solveHarder("sample_input.txt")
}

func SolveHarder() int {
	return solveHarder("input.txt")
}
