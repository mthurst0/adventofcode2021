package day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay12(t *testing.T) {
	assert.Equal(t, 0, SolveEasierSample())
	assert.Equal(t, 0, SolveEasier())
	assert.Equal(t, 0, SolveHarderSample())
	assert.Equal(t, 0, SolveHarder())
}
