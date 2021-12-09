package day8

import (
	"advent/rkutil"
	"strconv"
	"strings"
)

type Entry struct {
	Patterns []string
	Output   []string
}

func read(filename string) []Entry {
	lines := rkutil.MustLines(filename)
	var result []Entry
	for _, line := range lines {
		s := strings.Split(line, "|")
		rkutil.Ensure(len(s) == 2, "expect a single | in the input line")
		patterns := strings.Fields(strings.TrimSpace(s[0]))
		rkutil.Ensure(len(patterns) == 10, "expect 10 pattern strings")
		output := strings.Fields(strings.TrimSpace(s[1]))
		rkutil.Ensure(len(output) == 4, "expect 4 pattern strings")
		result = append(result, Entry{
			Patterns: patterns,
			Output:   output,
		})
	}
	return result
}

var segmentCounts = []int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}
var lenToDigits map[int][]int

func init() {
	lenToDigits = make(map[int][]int)
	lenToDigits[2] = []int{1}
	lenToDigits[3] = []int{7}
	lenToDigits[4] = []int{4}
	lenToDigits[5] = []int{2, 3, 5}
	lenToDigits[6] = []int{0, 6, 9}
	lenToDigits[7] = []int{8}
}

func solveEasier(filename string) int {
	entries := read(filename)
	result := 0
	for _, entry := range entries {
		for _, output := range entry.Output {
			if len(output) == segmentCounts[1] ||
				len(output) == segmentCounts[4] ||
				len(output) == segmentCounts[7] ||
				len(output) == segmentCounts[8] {
				result++
			}
		}
	}
	return result
}

func subtractStr(lhs, rhs string) string {
	for i := 0; i < len(rhs); i++ {
		lhs = strings.ReplaceAll(lhs, string(rhs[i]), "")
	}
	return lhs
}

func subtractStrs(strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]
	for _, s := range strs[1:] {
		result = subtractStr(result, s)
	}
	return result
}

func unionStr(lhs, rhs string) string {
	result := ""
	for i := 0; i < len(rhs); i++ {
		if strings.Index(lhs, string(rhs[i])) != -1 {
			result += string(rhs[i])
		}
	}
	return result
}

func unionStrs(strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]
	for _, s := range strs[1:] {
		result = unionStr(result, s)
	}
	return result
}

func removeStringsWithout(without string, strs ...string) []string {
	var result []string
	if len(strs) == 0 {
		return nil
	}
	for _, s := range strs {
		if strings.Index(s, without) != -1 {
			result = append(result, s)
		}
	}
	return result
}

func solveHarder(filename string) int {
	entries := read(filename)
	result := 0
	for _, entry := range entries {
		strs := make(map[int][]string)
		for _, p := range entry.Patterns {
			strs[len(p)] = append(strs[len(p)], p)
		}
		// 3 segment "7" minus 2 segment "1" gives the top segment
		//topSegment := subtractStr(strs[3][0], strs[2][0])
		// the three 5-digit numbers share the middle three segments
		middleThreeSegments := unionStrs(strs[5]...)
		// the single 4-digit number has only the middle segment, pick it from the middle
		// three segments
		middleSegment := unionStr(middleThreeSegments, strs[4][0])
		// the bottom segment is the other one left from the middle three segments
		//bottomSegment := subtractStrs(middleThreeSegments, topSegment, middleSegment)
		// subtracting the segments from the "1", and the middle segment from the "4"
		// gives us the top left segment
		topLeftSegment := subtractStrs(strs[4][0], strs[2][0], middleSegment)
		// knowing the middle segment we can pick the "6" and "9" from the 6-digit numbers
		// by filtering out the entry that contains the middle segment
		sixAndNineDigits := removeStringsWithout(middleSegment, strs[6]...)
		// with the "6" and "9" we can union them and subtract the middle three segments, this
		// leaves us with two segments, one being the bottom left and bottom right. The segment
		// that is also in the "1" is the bottom right
		bottomRightSegment := unionStr(subtractStrs(
			unionStr(sixAndNineDigits[0], sixAndNineDigits[1]), middleThreeSegments), strs[2][0])
		// knowing the bottom right, gives us the top right using the "1"
		topRightSegment := subtractStr(strs[2][0], bottomRightSegment)
		// and by process of elimination, we get the bottom left
		bottomLeftSegment := subtractStrs(
			"abcdefg", middleThreeSegments, topLeftSegment, topRightSegment, bottomRightSegment)

		digits := ""
		for _, o := range entry.Output {
			n := lenToDigits[len(o)]
			if len(n) == 1 {
				digits += strconv.Itoa(n[0])
			} else {
				if len(o) == 5 {
					if strings.Index(o, topLeftSegment) != -1 {
						digits += "5"
					} else if strings.Index(o, bottomLeftSegment) != -1 {
						digits += "2"
					} else {
						digits += "3"
					}
				} else if len(o) == 6 {
					if strings.Index(o, middleSegment) == -1 {
						digits += "0"
					} else if unionStr(o, strs[2][0]) == strs[2][0] {
						digits += "9"
					} else {
						digits += "6"
					}
				}
			}
		}
		v, err := strconv.Atoi(digits)
		if err != nil {
			rkutil.Unexpected(err)
		}
		result += v
	}
	return result
}

func SolveEasierSample() int {
	return solveEasier("sample_input.txt")
}

func SolveEasier() int {
	return solveEasier("input.txt")
}

func SolveHarderSample() int {
	return solveHarder("sample_input.txt")
}

func SolveHarder() int {
	return solveHarder("input.txt")
}
