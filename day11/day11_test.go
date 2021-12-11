package day11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay11(t *testing.T) {
	assert.Equal(t, 1656, SolveEasierSample())
	assert.Equal(t, 1661, SolveEasier())
	assert.Equal(t, 195, SolveHarderSample())
	assert.Equal(t, 334, SolveHarder())
}
