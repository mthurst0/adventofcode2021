package day6

import (
	"advent/rkutil"
	"strconv"
)

var input = "2,4,1,5,1,3,1,1,5,2,2,5,4,2,1,2,5,3,2,4,1,3,5,3,1,3,1,3,5,4,1,1,1,1,5,1,2,5,5,5,2,3,4,1,1,1,2,1,4,1,3,2,1,4,3,1,4,1,5,4,5,1,4,1,2,2,3,1,1,1,2,5,1,1,1,2,1,1,2,2,1,4,3,3,1,1,1,2,1,2,5,4,1,4,3,1,5,5,1,3,1,5,1,5,2,4,5,1,2,1,1,5,4,1,1,4,5,3,1,4,5,1,3,2,2,1,1,1,4,5,2,2,5,1,4,5,2,1,1,5,3,1,1,1,3,1,2,3,3,1,4,3,1,2,3,1,4,2,1,2,5,4,2,5,4,1,1,2,1,2,4,3,3,1,1,5,1,1,1,1,1,3,1,4,1,4,1,2,3,5,1,2,5,4,5,4,1,3,1,4,3,1,2,2,2,1,5,1,1,1,3,2,1,3,5,2,1,1,4,4,3,5,3,5,1,4,3,1,3,5,1,3,4,1,2,5,2,1,5,4,3,4,1,3,3,5,1,1,3,5,3,3,4,3,5,5,1,4,1,1,3,5,5,1,5,4,4,1,3,1,1,1,1,3,2,1,2,3,1,5,1,1,1,4,3,1,1,1,1,1,1,1,1,1,2,1,1,2,5,3"

func solve(days int) string {
	states := rkutil.MustCommaDelimitedNumbers(input)
	//states := rkutil.MustCommaDelimitedNumbers("3,4,3,1,2")
	var digitCounts [9]int
	for _, s := range states {
		digitCounts[s]++
	}
	for i := 0; i < days; i++ {
		newOnes := digitCounts[0]
		digitCounts[0] = digitCounts[1]
		digitCounts[1] = digitCounts[2]
		digitCounts[2] = digitCounts[3]
		digitCounts[3] = digitCounts[4]
		digitCounts[4] = digitCounts[5]
		digitCounts[5] = digitCounts[6]
		digitCounts[6] = digitCounts[7] + newOnes
		digitCounts[7] = digitCounts[8]
		digitCounts[8] = newOnes
	}
	result := 0
	for i := 0; i < 9; i++ {
		result += digitCounts[i]
	}
	return strconv.Itoa(result)
}

func SolveEasier() string {
	return solve(80)
}

func SolveHarder() string {
	return solve(256)
}
