package day8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay8(t *testing.T) {
	assert.Equal(t, 26, SolveEasierSample())
	assert.Equal(t, 383, SolveEasier())
	assert.Equal(t, 61229, SolveHarderSample())
	assert.Equal(t, 998900, SolveHarder())
}
