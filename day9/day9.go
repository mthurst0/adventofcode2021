package day9

import (
	"advent/rkutil"
	"sort"
)

func val(p [][]int, x, y int) int {
	if x < 0 {
		return 10
	}
	if y < 0 {
		return 10
	}
	if x >= len(p) {
		return 10
	}
	if y >= len(p[0]) {
		return 10
	}
	return p[x][y]
}

func solveEasier(filename string) int {
	p := rkutil.ReadGrid(filename)
	var risks []int
	for x := 0; x < len(p); x++ {
		for y := 0; y < len(p[x]); y++ {
			v := p[x][y]
			if v < val(p, x-1, y) &&
				v < val(p, x+1, y) &&
				v < val(p, x, y-1) &&
				v < val(p, x, y+1) {
				risks = append(risks, v)
			}
		}
	}
	result := 0
	for _, r := range risks {
		result += r + 1
	}
	return result
}

func SolveEasierSample() int {
	return solveEasier("sample_input.txt")
}

func SolveEasier() int {
	return solveEasier("input.txt")
}

type Coordinate struct {
	X int
	Y int
}

func NewInvalidCoordinate() Coordinate {
	return Coordinate{X: -1, Y: -1}
}

func (c Coordinate) IsValid() bool {
	return c.X != -1 && c.Y != -1
}

type Basin struct {
	Unexplored []Coordinate
	Explored   []Coordinate
}

func (basin *Basin) AddExplored(c Coordinate) {
	if !basin.InExplored(c) {
		basin.Explored = append(basin.Explored, c)
	}
}

func (basin *Basin) AddUnexplored(c Coordinate) {
	if !basin.InUnexplored(c) && !basin.InExplored(c) {
		basin.Unexplored = append(basin.Unexplored, c)
	}
}

func (basin Basin) InExplored(c Coordinate) bool {
	for _, coord := range basin.Explored {
		if c.X == coord.X && c.Y == coord.Y {
			return true
		}
	}
	return false
}

func (basin Basin) InUnexplored(c Coordinate) bool {
	for _, coord := range basin.Unexplored {
		if c.X == coord.X && c.Y == coord.Y {
			return true
		}
	}
	return false
}

func (basin Basin) HasUnexplored() bool {
	return len(basin.Unexplored) > 0
}

type Basins []Basin

func (basins Basins) InAnyBasin(c Coordinate) bool {
	for _, basin := range basins {
		if basin.InExplored(c) {
			return true
		}
	}
	return false
}

func findStartOfBasin(p [][]int, startCoord Coordinate, basins Basins) Coordinate {
	for x := startCoord.X; x < len(p); x++ {
		for y := startCoord.Y; y < len(p[0]); y++ {
			if p[x][y] != 9 {
				coord := Coordinate{X: x, Y: y}
				if !basins.InAnyBasin(coord) {
					return coord
				}
			}
		}
	}
	return NewInvalidCoordinate()
}

func exploreBasin(p [][]int, startCoord Coordinate) Basin {
	var basin Basin
	exploreCoordinate(p, &basin, startCoord)
	for basin.HasUnexplored() {
		unexplored := basin.Unexplored[0]
		basin.Unexplored = basin.Unexplored[1:]
		exploreCoordinate(p, &basin, unexplored)
	}
	return basin
}

func exploreCoordinate(p [][]int, basin *Basin, startCoord Coordinate) {
	basin.AddExplored(startCoord)
	for x := startCoord.X - 1; x >= 0; x-- {
		if x > len(p) {
			break
		}
		if p[x][startCoord.Y] == 9 {
			break
		}
		basin.AddUnexplored(Coordinate{X: x, Y: startCoord.Y})
	}
	for x := startCoord.X + 1; x < len(p); x++ {
		if x >= len(p) {
			break
		}
		if p[x][startCoord.Y] == 9 {
			break
		}
		basin.AddUnexplored(Coordinate{X: x, Y: startCoord.Y})
	}
	for y := startCoord.Y - 1; y >= 0; y-- {
		if p[startCoord.X][y] == 9 {
			break
		}
		basin.AddUnexplored(Coordinate{X: startCoord.X, Y: y})
	}
	for y := startCoord.Y + 1; y < len(p[0]); y++ {
		if y >= len(p[0]) {
			break
		}
		if p[startCoord.X][y] == 9 {
			break
		}
		basin.AddUnexplored(Coordinate{X: startCoord.X, Y: y})
	}
}

// TODO: brute force performance
func solveHarder(filename string) int {
	p := rkutil.ReadGrid(filename)
	var basins Basins
	var curCoordinate Coordinate
	for {
		startCoord := findStartOfBasin(p, curCoordinate, basins)
		if !startCoord.IsValid() {
			break
		}
		newBasin := exploreBasin(p, startCoord)
		basins = append(basins, newBasin)
	}
	var sizes []int
	for _, basin := range basins {
		sizes = append(sizes, len(basin.Explored))
	}
	sort.Ints(sizes)
	rkutil.Ensure(len(sizes) > 3, "expected at least 3 basins")
	return sizes[len(sizes)-3] * sizes[len(sizes)-2] * sizes[len(sizes)-1]
}

func SolveHarderSample() int {
	return solveHarder("sample_input.txt")
}

func SolveHarder() int {
	return solveHarder("input.txt")
}
