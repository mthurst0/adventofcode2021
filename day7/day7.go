package day7

import (
	"advent/rkutil"
	"fmt"
	"math"
	"strconv"
)

var sampleInput = "16,1,2,0,4,2,7,1,2,14"

func calcFuelCheap(nums []int, target int) int {
	fuel := 0
	for _, n := range nums {
		diff := target - n
		if diff < 0 {
			diff = -diff
		}
		fuel += diff
	}
	return fuel
}

func s(diff int) int {
	if diff%2 == 0 {
		return (diff + 1) * (diff / 2)
	} else {
		diff--
		return (diff+1)*(diff/2) + (diff + 1)
	}
}

func calcFuelExpensive(nums []int, target int) int {
	fuel := 0
	for _, n := range nums {
		diff := target - n
		if diff < 0 {
			diff = -diff
		}
		fuel += s(diff)
	}
	return fuel
}

func solve(nums []int, calcFuel func([]int, int) int) int {
	sum := 0
	max := 0
	for _, n := range nums {
		sum += n
		if n > max {
			max = n
		}
	}
	//fmt.Printf("%d %d %d", sum/len(nums), len(nums), max)
	point := int(math.Round((float64(sum) / float64(len(nums))) / 2))

	min := calcFuel(nums, point)
	fmt.Printf("%d %d\n", point, min)
	for i := point; i > 0; i-- {
		f := calcFuel(nums, i)
		if f < min {
			fmt.Printf("%d %d\n", i, f)
			min = f
		} else {
			fmt.Printf("%d %d (stop)\n", i, f)
			break
		}
	}
	for i := point; i < max; i++ {
		f := calcFuel(nums, i)
		if f <= min {
			fmt.Printf("%d %d\n", i, f)
			min = f
		} else {
			fmt.Printf("%d %d (stop)\n", i, f)
			break
		}
	}
	return min
}

func SolveEasierSample() string {
	return strconv.Itoa(solve(rkutil.MustCommaDelimitedNumbers(sampleInput), calcFuelCheap))
}

func SolveEasier() string {
	lines := rkutil.MustLines("input.txt")
	return strconv.Itoa(solve(rkutil.MustCommaDelimitedNumbers(lines[0]), calcFuelCheap))
}

func SolveHarderSample() string {
	return strconv.Itoa(solve(rkutil.MustCommaDelimitedNumbers(sampleInput), calcFuelExpensive))
}

func SolveHarder() string {
	lines := rkutil.MustLines("input.txt")
	return strconv.Itoa(solve(rkutil.MustCommaDelimitedNumbers(lines[0]), calcFuelExpensive))
}
