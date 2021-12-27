package day14

import (
	"advent/rkutil"
	"fmt"
	"sort"
	"strings"
)

type Pair struct {
	LHS uint8
	RHS uint8
}

func (pair Pair) Key() string {
	return string(pair.LHS) + string(pair.RHS)
}

func Pairs(polymer string) []Pair {
	var pairs []Pair
	for i := 0; i < len(polymer)-1; i++ {
		pairs = append(pairs, Pair{LHS: polymer[i], RHS: polymer[i+1]})
	}
	return pairs
}

type Game struct {
	Polymer string
	Rules   map[string][]Pair
}

func readGame(filename string) Game {
	lines := rkutil.MustLines(filename)
	g := Game{Rules: make(map[string][]Pair)}
	g.Polymer = strings.TrimSpace(lines[0])
	for _, line := range lines[2:] {
		s := strings.Split(line, "->")
		lhs := strings.TrimSpace(s[0])
		rhs := strings.TrimSpace(s[1])
		lhsPair := Pair{LHS: lhs[0], RHS: rhs[0]}
		rhsPair := Pair{LHS: rhs[0], RHS: lhs[1]}
		g.Rules[lhs] = append(g.Rules[lhs], lhsPair)
		g.Rules[lhs] = append(g.Rules[lhs], rhsPair)
	}
	return g
}

func counts(polymer string) map[string]int {
	result := make(map[string]int)
	for _, c := range polymer {
		result[string(c)]++
	}
	return result
}

func toString(pairs []Pair) string {
	var result strings.Builder
	for _, pair := range pairs {
		result.WriteString(string(pair.LHS))
	}
	result.WriteString(string(pairs[len(pairs)-1].RHS))
	return result.String()
}

func solveEasier(filename string, steps int) int {
	g := readGame(filename)
	curPairs := Pairs(g.Polymer)
	for step := 0; step < steps; step++ {
		var newPairs []Pair
		for _, pair := range curPairs {
			insertion, ok := g.Rules[pair.Key()]
			if !ok {
				panic("could not find pair: " + pair.Key())
			}
			newPairs = append(newPairs, insertion...)
		}
		curPairs = newPairs
		fmt.Printf("%d: (%d)\n", step+1, len(newPairs))
	}
	countMap := counts(toString(curPairs))
	var results []int
	for _, count := range countMap {
		results = append(results, count)
	}
	sort.Ints(results)
	return results[len(results)-1] - results[0]
}

func copyMap(src map[string]int) map[string]int {
	result := make(map[string]int)
	for k, v := range src {
		result[k] = v
	}
	return result
}

func solveHarder(filename string, steps int) int {
	g := readGame(filename)
	curPairs := Pairs(g.Polymer)
	pairCounts := make(map[string]int)
	for _, pair := range curPairs {
		pairCounts[pair.Key()]++
	}
	for step := 0; step < steps; step++ {
		resultCounts := copyMap(pairCounts)
		for pairStr, j := range pairCounts {
			insertion, ok := g.Rules[pairStr]
			if !ok {
				panic("could not find pair: " + pairStr)
			}
			resultCounts[insertion[0].Key()] += j
			resultCounts[insertion[1].Key()] += j
			resultCounts[pairStr] -= j
		}
		pairCounts = resultCounts
	}
	charCounts := make(map[string]int)
	for k, pairCount := range pairCounts {
		charCounts[string(k[0:1])] += pairCount
	}
	var results []int
	for _, count := range charCounts {
		results = append(results, count)
	}
	sort.Ints(results)
	return results[len(results)-1] - results[0] + 1 // +1 is on account of the last character
}

func SolveEasierSample() int {
	return solveEasier("sample_input.txt", 10)
}

func SolveEasier() int {
	return solveEasier("input.txt", 10)
}

func SolveHarderSample() int {
	return solveHarder("sample_input.txt", 40)
}

func SolveHarder() int {
	return solveHarder("input.txt", 40)
}
