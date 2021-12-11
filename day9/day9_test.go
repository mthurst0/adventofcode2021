package day9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay9(t *testing.T) {
	assert.Equal(t, 15, SolveEasierSample())
	assert.Equal(t, 535, SolveEasier())
	assert.Equal(t, 1134, SolveHarderSample())
	assert.Equal(t, 1122700, SolveHarder())
}
