package day7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay7(t *testing.T) {
	assert.Equal(t, "37", SolveEasierSample())
	assert.Equal(t, "335330", SolveEasier())
	assert.Equal(t, "168", SolveHarderSample())
	assert.Equal(t, "168", SolveHarder())
}
