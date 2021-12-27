package day16

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToday(t *testing.T) {
	assert.Equal(t, 16, SolveEasierString("8A004A801A8002F478"))
	assert.Equal(t, 12, SolveEasierString("620080001611562C8802118E34"))
	assert.Equal(t, 23, SolveEasierString("C0015000016115A2E0802F182340"))
	assert.Equal(t, 31, SolveEasierString("A0016C880162017C3686B18A3D4780"))
	assert.Equal(t, 960, SolveEasier())

	assert.Equal(t, 3, SolveHarderString("C200B40A82"))
	assert.Equal(t, 54, SolveHarderString("04005AC33890"))
	assert.Equal(t, 7, SolveHarderString("880086C3E88112"))
	assert.Equal(t, 9, SolveHarderString("CE00C43D881120"))
	assert.Equal(t, 1, SolveHarderString("D8005AC2A8F0"))
	assert.Equal(t, 0, SolveHarderString("F600BC2D8F"))
	assert.Equal(t, 0, SolveHarderString("9C005AC2F8F0"))
	assert.Equal(t, 1, SolveHarderString("9C0141080250320F1802104A08"))
	assert.Equal(t, 12301926782560, SolveHarder())
}
