package day1

import (
	"advent/rkutil"
	"fmt"
	"strconv"
	"strings"
)

func SolveEasier() string {
	lines := rkutil.MustLines("day1/input.txt")
	last := 0
	result := 0
	first := true
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			rkutil.Unexpected(err)
		}
		if num > last {
			if !first {
				result++
			}
		}
		first = false
		last = num
	}
	return fmt.Sprintf("%d", result)
}

func SolveHarder() string {
	v := rkutil.MustLinesOfNumbers("day1/input.txt")
	if len(v) <= 3 {
		rkutil.Unexpected(fmt.Errorf("expected at least 3 values: %d", len(v)))
	}
	cur0 := v[0]
	cur1 := v[1]
	cur2 := v[2]
	result := 0
	for i := 3; i < len(v); i++ {
		window := cur1 + cur2 + v[i]
		if window > (cur0 + cur1 + cur2) {
			result++
		}
		cur0 = cur1
		cur1 = cur2
		cur2 = v[i]
	}
	return fmt.Sprintf("%d", result)
}
