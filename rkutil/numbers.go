package rkutil

import (
	"fmt"
	"strconv"
	"strings"
)

func MustParseInt(s string, base int, bitSize int) int64 {
	result, err := strconv.ParseInt(s, base, bitSize)
	if err != nil {
		Unexpected(fmt.Errorf("expected integer conversion failure: %w", err))
	}
	return result
}

func MustCommaDelimitedNumbers(s string) []int {
	var result []int
	for _, s1 := range strings.Split(s, ",") {
		v, err := strconv.Atoi(s1)
		if err != nil {
			Unexpected(fmt.Errorf("expected integer conversion failure: %w", err))
		}
		result = append(result, v)
	}
	return result
}

func MustWhitespaceDelimitedNumbers(s string) []int {
	fs := strings.Fields(s)
	var result []int
	for _, f := range fs {
		v, err := strconv.Atoi(f)
		if err != nil {
			Unexpected(fmt.Errorf("expected integer conversion failure: %w", err))
		}
		result = append(result, v)
	}
	return result
}

func MaxInt(lhs, rhs int) int {
	if lhs > rhs {
		return lhs
	}
	return rhs
}

func MinInt(lhs, rhs int) int {
	if lhs < rhs {
		return lhs
	}
	return rhs
}

func MaxUint64(lhs, rhs uint64) uint64 {
	if lhs > rhs {
		return lhs
	}
	return rhs
}

func MinUint64(lhs, rhs uint64) uint64 {
	if lhs < rhs {
		return lhs
	}
	return rhs
}
