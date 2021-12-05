package day2

import (
	"advent/rkutil"
	"fmt"
	"strconv"
	"strings"
)

func SolveEasier() string {
	lines := rkutil.MustTrimmedLines("day2/input.txt")
	depth := 0
	horizontal := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		v := strings.Split(line, " ")
		if len(v) != 2 {
			rkutil.Unexpected(fmt.Errorf("expecting 2 tokens: %d", len(v)))
		}
		value, err := strconv.Atoi(v[1])
		if err != nil {
			rkutil.Unexpected(fmt.Errorf("expected number: %s", v[1]))
		}
		switch v[0] {
		case "forward":
			horizontal += value
		case "up":
			depth -= value
			if depth < 0 {
				depth = 0
			}
		case "down":
			depth += value
		default:
			rkutil.Unexpected(fmt.Errorf("unexpected direction: %s", v[0]))
		}
	}
	return fmt.Sprintf("%d", depth*horizontal)
}

func SolveHarder() string {
	lines := rkutil.MustTrimmedLines("day2/input.txt")
	depth := 0
	horizontal := 0
	aim := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		v := strings.Split(line, " ")
		if len(v) != 2 {
			rkutil.Unexpected(fmt.Errorf("expecting 2 tokens: %d", len(v)))
		}
		value, err := strconv.Atoi(v[1])
		if err != nil {
			rkutil.Unexpected(fmt.Errorf("expected number: %s", v[1]))
		}
		switch v[0] {
		case "forward":
			horizontal += value
			depth += aim * value
		case "up":
			aim -= value
			// TODO: check for <0, or does it just not matter?
		case "down":
			aim += value
		default:
			rkutil.Unexpected(fmt.Errorf("unexpected direction: %s", v[0]))
		}
	}
	return fmt.Sprintf("%d", depth*horizontal)
}
