package main

import (
	"advent/day1"
	"advent/day2"
	"advent/day3"
	"fmt"
)

func main() {
	// 1521
	fmt.Println("day1 easier: " + day1.SolveEasier())
	// 1543
	fmt.Println("day1 harder: " + day1.SolveHarder())
	// 2120749
	fmt.Println("day2 easier: " + day2.SolveEasier())
	// 2138382217
	fmt.Println("day2 easier: " + day2.SolveHarder())
	// 4103154
	fmt.Println("day3 easier: " + day3.SolveEasier())
	// 4245351
	fmt.Println("day3 harder: " + day3.SolveHarder())
}
