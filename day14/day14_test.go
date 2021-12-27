package day14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToday(t *testing.T) {
	assert.Equal(t, 1588, SolveEasierSample())
	assert.Equal(t, 2768, SolveEasier())
	assert.Equal(t, 2188189693529, SolveHarderSample())
	assert.Equal(t, 2914365137499, SolveHarder())
}
