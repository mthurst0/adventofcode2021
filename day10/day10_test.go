package day10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay10(t *testing.T) {
	assert.Equal(t, 26397, SolveEasierSample())
	assert.Equal(t, 341823, SolveEasier())
	assert.Equal(t, 288957, SolveHarderSample())
	assert.Equal(t, 2801302861, SolveHarder())
}
