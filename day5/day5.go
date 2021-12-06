package day5

import (
	"advent/rkutil"
	"fmt"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

func (coord *Coordinate) Add(c Coordinate) {
	coord.X += c.X
	coord.Y += c.Y
}

func parseCoordinate(s string) Coordinate {
	vs := rkutil.MustCommaDelimitedNumbers(s)
	rkutil.Ensure(len(vs) == 2, fmt.Sprintf("unexpected token count: %d", len(vs)))
	return Coordinate{X: vs[0], Y: vs[1]}
}

type Vent struct {
	Start Coordinate
	End   Coordinate
}

type Vents []Vent

func MustVents() Vents {
	var vents Vents
	lines := rkutil.MustLines("day5/input.txt")
	for _, line := range lines {
		cs := strings.Split(line, " -> ")
		rkutil.Ensure(len(cs) == 2, fmt.Sprintf("unexpected token count: %d", len(cs)))
		start := parseCoordinate(cs[0])
		end := parseCoordinate(cs[1])
		vents = append(vents, Vent{Start: start, End: end})
	}
	return vents
}

func (vents Vents) MaxXCoordinate() int {
	m := 0
	for _, vent := range vents {
		if vent.Start.X > m {
			m = vent.Start.X
		}
		if vent.End.X > m {
			m = vent.Start.X
		}
	}
	return m
}

func (vents Vents) MaxYCoordinate() int {
	m := 0
	for _, vent := range vents {
		if vent.Start.Y > m {
			m = vent.Start.Y
		}
		if vent.End.Y > m {
			m = vent.Start.Y
		}
	}
	return m
}

func drawMap(m [][]int, maxX, maxY int) {
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if m[x][y] == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", m[x][y])
			}
		}
		fmt.Println()
	}
}

func SolveEasier() string {
	vents := MustVents()
	maxX := vents.MaxXCoordinate() + 1
	maxY := vents.MaxYCoordinate() + 1
	m := make([][]int, maxX)
	for i := 0; i < maxX; i++ {
		m[i] = make([]int, maxY)
	}
	for _, vent := range vents {
		if vent.Start.X == vent.End.X {
			if vent.Start.Y > vent.End.Y {
				for y := vent.End.Y; y <= vent.Start.Y; y++ {
					m[vent.Start.X][y]++
				}
			} else {
				for y := vent.Start.Y; y <= vent.End.Y; y++ {
					m[vent.Start.X][y]++
				}
			}
		} else if vent.Start.Y == vent.End.Y {
			if vent.Start.X > vent.End.X {
				for x := vent.End.X; x <= vent.Start.X; x++ {
					m[x][vent.Start.Y]++
				}
			} else {
				for x := vent.Start.X; x <= vent.End.X; x++ {
					m[x][vent.Start.Y]++
				}
			}
		}
	}
	result := 0
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			if m[x][y] >= 2 {
				result++
			}
		}
	}
	return strconv.Itoa(result)
}

func SolveHarder() string {
	vents := MustVents()
	maxX := vents.MaxXCoordinate() + 1
	maxY := vents.MaxYCoordinate() + 1
	m := make([][]int, maxX)
	for i := 0; i < maxX; i++ {
		m[i] = make([]int, maxY)
	}
	for _, vent := range vents {
		if vent.Start.X == vent.End.X {
			if vent.Start.Y > vent.End.Y {
				for y := vent.End.Y; y <= vent.Start.Y; y++ {
					m[vent.Start.X][y]++
				}
			} else {
				for y := vent.Start.Y; y <= vent.End.Y; y++ {
					m[vent.Start.X][y]++
				}
			}
		} else if vent.Start.Y == vent.End.Y {
			if vent.Start.X > vent.End.X {
				for x := vent.End.X; x <= vent.Start.X; x++ {
					m[x][vent.Start.Y]++
				}
			} else {
				for x := vent.Start.X; x <= vent.End.X; x++ {
					m[x][vent.Start.Y]++
				}
			}
		} else {
			incCoord := Coordinate{}
			if vent.Start.X > vent.End.X {
				if vent.Start.Y > vent.End.Y {
					incCoord = Coordinate{-1, -1}
				} else {
					incCoord = Coordinate{-1, 1}
				}
			} else {
				if vent.Start.Y > vent.End.Y {
					incCoord = Coordinate{1, -1}
				} else {
					incCoord = Coordinate{1, 1}
				}
			}
			coord := vent.Start
			for {
				m[coord.X][coord.Y]++
				coord.Add(incCoord)
				if coord.X == vent.End.X && coord.Y == vent.End.Y {
					m[coord.X][coord.Y]++
					break
				}
			}
		}
	}
	result := 0
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			if m[x][y] >= 2 {
				result++
			}
		}
	}
	return strconv.Itoa(result)
}

// 1,1 -> 3,3
