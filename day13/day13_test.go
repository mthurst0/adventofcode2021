package day13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToday(t *testing.T) {
	assert.Equal(t, 17, SolveEasierSample())
	assert.Equal(t, 638, SolveEasier())
	assert.Equal(t, 16, SolveHarderSample())
	assert.Equal(t, 98, SolveHarder())
}
