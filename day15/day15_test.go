package day15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToday(t *testing.T) {
	assert.Equal(t, 40, SolveEasierSample())
	assert.Equal(t, 487, SolveEasier())
	assert.Equal(t, 315, SolveHarderSample())
	assert.Equal(t, 0, SolveHarder())
}
