package day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay12(t *testing.T) {
	assert.Equal(t, 10, SolveEasierSample())
	assert.Equal(t, 4707, SolveEasier())
	assert.Equal(t, 36, SolveHarderSample())
	assert.Equal(t, 130493, SolveHarder())
}
