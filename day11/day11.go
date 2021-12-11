package day11

import (
	"advent/rkutil"
	"fmt"
)

type Cell struct {
	Count    int
	Exploded bool
}

type Grid struct {
	Flashes int
	Map     [][]Cell
}

func incAndExplode(grid *Grid, x, y int) {
	if y >= len(grid.Map[0]) || y < 0 || x >= len(grid.Map) || x < 0 {
		return
	}
	if !grid.Map[x][y].Exploded {
		grid.Map[x][y].Count++
		if grid.Map[x][y].Count > 9 {
			explode(grid, x, y)
		}
	}
}

func explode(grid *Grid, x, y int) {
	if !grid.Map[x][y].Exploded {
		grid.Map[x][y].Count = 0
		grid.Map[x][y].Exploded = true
		grid.Flashes++
		incAndExplode(grid, x-1, y-1)
		incAndExplode(grid, x-1, y)
		incAndExplode(grid, x-1, y+1)
		incAndExplode(grid, x, y-1)
		incAndExplode(grid, x, y+1)
		incAndExplode(grid, x+1, y-1)
		incAndExplode(grid, x+1, y)
		incAndExplode(grid, x+1, y+1)
	}
}

func newGrid(from *Grid) *Grid {
	var g Grid
	g.Flashes = from.Flashes
	g.Map = make([][]Cell, len(from.Map))
	for j := 0; j < len(from.Map); j++ {
		g.Map[j] = make([]Cell, len(from.Map[0]))
	}
	for x := 0; x < len(from.Map); x++ {
		for y := 0; y < len(from.Map[0]); y++ {
			g.Map[x][y].Count = from.Map[x][y].Count
		}
	}
	return &g
}

func draw(g *Grid) {
	for x := 0; x < len(g.Map); x++ {
		for y := 0; y < len(g.Map[0]); y++ {
			fmt.Print(g.Map[x][y])
		}
		fmt.Println()
	}
}

func readGrid(filename string) *Grid {
	var grid Grid
	raw := rkutil.ReadGrid(filename)
	grid.Map = make([][]Cell, len(raw))
	for x := 0; x < len(raw); x++ {
		grid.Map[x] = make([]Cell, len(raw[0]))
		for y := 0; y < len(raw[0]); y++ {
			grid.Map[x][y].Count = raw[x][y]
		}
	}
	return &grid
}

func step(g *Grid) *Grid {
	next := newGrid(g)
	for x := 0; x < len(g.Map); x++ {
		for y := 0; y < len(g.Map[0]); y++ {
			if !next.Map[x][y].Exploded {
				next.Map[x][y].Count = next.Map[x][y].Count + 1
				if next.Map[x][y].Count > 9 {
					explode(next, x, y)
				}
			}
		}
	}
	return next
}

func solveEasier(filename string) int {
	g := readGrid(filename)
	for i := 0; i < 100; i++ {
		next := step(g)
		g = next
	}
	return g.Flashes
}

func SolveEasierSample() int {
	//return solveEasier("trivial_input.txt")
	return solveEasier("sample_input.txt")
}

func SolveEasier() int {
	return solveEasier("input.txt")
}

const safetyLimit = 1000

func solveHarder(filename string) int {
	g := readGrid(filename)
	for i := 0; i < safetyLimit; i++ {
		next := step(g)
		if next.Flashes-g.Flashes == len(g.Map)*len(g.Map[0]) {
			return i + 1 // 1-index steps
		}
		g = next
	}
	rkutil.UnexpectedCodePath()
	return 0
}

func SolveHarderSample() int {
	return solveHarder("sample_input.txt")
}

func SolveHarder() int {
	return solveHarder("input.txt")
}
