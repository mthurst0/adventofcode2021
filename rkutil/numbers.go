package rkutil

import (
	"fmt"
	"strconv"
)

func MustParseInt(s string, base int, bitSize int) int64 {
	result, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		Unexpected(fmt.Errorf("expected integer conversion failure: %w", err))
	}
	return result
}
