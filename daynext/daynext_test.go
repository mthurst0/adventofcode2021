package daynext

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToday(t *testing.T) {
	assert.Equal(t, 0, SolveEasierSample())
	assert.Equal(t, 0, SolveEasier())
	assert.Equal(t, 0, SolveHarderSample())
	assert.Equal(t, 0, SolveHarder())
}
