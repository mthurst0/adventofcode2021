package day3

import (
	"advent/rkutil"
	"fmt"
	"strconv"
	"strings"
)

func bitCounts(lines []string) []int {
	lenOfFirstLine := len(lines[0])
	result := make([]int, lenOfFirstLine)
	for _, line := range lines {
		if len(line) != lenOfFirstLine {
			rkutil.Unexpected(fmt.Errorf(
				"expected all lines to be the same length: %d vs %d", lenOfFirstLine, len(line)))
		}
		for i := 0; i < len(line); i++ {
			if line[i] == '1' {
				result[i]++
			}
		}
	}
	return result
}

func bitCountInRow(v map[string]bool, idx int) int {
	result := 0
	for line, _ := range v {
		if line[idx] == '1' {
			result++
		}
	}
	return result
}

func SolveEasier() string {
	lines := rkutil.MustTrimmedLines("day3/input.txt")
	if len(lines) < 1 {
		rkutil.Unexpected(fmt.Errorf("expected at least one line: %d", len(lines)))
	}
	bc := bitCounts(lines)
	midpoint := len(lines) / 2
	gamma := strings.Builder{}
	epsilon := strings.Builder{}
	for i := 0; i < len(lines[0]); i++ {
		if bc[i] >= midpoint {
			gamma.WriteString("1")
			epsilon.WriteString("0")
		} else {
			gamma.WriteString("0")
			epsilon.WriteString("1")
		}
	}
	g, err := strconv.ParseInt(gamma.String(), 2, 64)
	if err != nil {
		rkutil.Unexpected(fmt.Errorf("expected integer conversion failure: %w", err))
	}
	e, err := strconv.ParseInt(epsilon.String(), 2, 64)
	if err != nil {
		rkutil.Unexpected(fmt.Errorf("expected integer conversion failure: %w", err))
	}
	return fmt.Sprintf("%d", g*e)
}

func include(bc int, midpoint int, lhs, rhs uint8) bool {
	if bc >= midpoint {
		if lhs == rhs {
			return true
		}
	} else {
		if lhs != rhs {
			return true
		}
	}
	return false
}

func filter(bcs []int, lines []string, r uint8) string {
	midpoint := len(lines) / 2
	fMap := make(map[string]bool)
	for _, line := range lines {
		if include(bcs[0], midpoint, line[0], r) {
			fMap[line] = true
		}
	}
	for chIndex := 1; chIndex < len(lines[0]); chIndex++ {
		bc := bitCountInRow(fMap, chIndex)
		midpoint = (len(fMap) + 1) / 2
		for k, _ := range fMap {
			if !include(bc, midpoint, k[chIndex], r) {
				delete(fMap, k)
				if len(fMap) == 1 {
					for kk := range fMap {
						return kk
					}
				}
			}
		}
	}
	rkutil.Unexpected(fmt.Errorf("filter walk ended unexpectedly"))
	return ""
}

func SolveHarder() string {
	lines := rkutil.MustTrimmedLines("day3/input.txt")
	if len(lines) < 1 {
		rkutil.Unexpected(fmt.Errorf("expected at least one line: %d", len(lines)))
	}
	bc := bitCounts(lines)
	o2 := rkutil.MustParseInt(filter(bc, lines, '1'), 2, 64)
	co2 := rkutil.MustParseInt(filter(bc, lines, '0'), 2, 64)
	return fmt.Sprintf("%d", o2*co2)
}

//00100
//11110
//10110
//10111
//10101
//01111
//00111
//11100
//10000
//11001
//00010
//01010

//-5
//11110
//10110
//10111
//10101
//11100
//10000
//11001

//-3
//10110
//10111
//10101
//10000

//-1
//10110
//10111
//10101

//-2
//10101
