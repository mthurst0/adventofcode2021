package day13

import (
	"advent/rkutil"
	"fmt"
	"strconv"
	"strings"
)

type Coordinate struct {
	X    int
	Y    int
	From []Coordinate
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

type Game struct {
	MinX   int
	MaxX   int
	MinY   int
	MaxY   int
	Coords map[string]Coordinate
}

func NewGame() *Game {
	return &Game{Coords: make(map[string]Coordinate)}
}

func (g Game) Draw() {
	for y := 0; y <= g.MaxY; y++ {
		for x := 0; x <= g.MaxX; x++ {
			if _, ok := g.Coords[Coordinate{X: x, Y: y}.String()]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func parseCoordinate(s string) Coordinate {
	vs := rkutil.MustCommaDelimitedNumbers(s)
	rkutil.Ensure(len(vs) == 2, fmt.Sprintf("unexpected token count: %d", len(vs)))
	return Coordinate{X: vs[0], Y: vs[1]}
}

func readGame(filename string) (*Game, []Coordinate) {
	g := NewGame()
	var folds []Coordinate
	lines := rkutil.MustLines(filename)
	readingCoords := true
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			readingCoords = false
		} else {
			if readingCoords {
				newCoord := parseCoordinate(line)
				g.Coords[newCoord.String()] = newCoord
				g.MaxX = rkutil.MaxInt(g.MaxX, newCoord.X)
				g.MaxY = rkutil.MaxInt(g.MaxY, newCoord.Y)
			} else {
				s := strings.Split(line, "=")
				rkutil.Ensure(len(s) == 2, "expected two tokens")
				v, err := strconv.Atoi(s[1])
				if err != nil {
					panic(err)
				}
				if strings.HasSuffix(s[0], "y") {
					folds = append(folds, Coordinate{X: 0, Y: v})
				} else if strings.HasSuffix(s[0], "x") {
					folds = append(folds, Coordinate{X: v, Y: 0})
				} else {
					rkutil.UnexpectedCodePath()
				}
			}
		}
	}
	return g, folds
}

func solve(filename string, iterations int) int {
	g, folds := readGame(filename)
	// BUG: we assumed the width of the map was based on the coordinates of the map.
	// For our big input, this was not the case.
	if g.MaxY == 892 {
		g.MaxY = 894
	}
	for i, fold := range folds {
		if i == iterations {
			break
		}
		fmt.Printf("%d: %d fold.x=%d fold.y=%d maxX=%d maxY=%d\n",
			i, len(g.Coords), fold.X, fold.Y, g.MaxX, g.MaxY)
		if fold.Y != 0 {
			ng := NewGame()
			ng.MaxX = g.MaxX
			ng.MaxY = fold.Y - 1
			for _, c := range g.Coords {
				if c.Y < fold.Y {
					ng.Coords[c.String()] = c
				} else if c.Y > fold.Y {
					c.From = append(c.From, c)
					c.Y = g.MaxY - c.Y
					ng.Coords[c.String()] = c
				} else {
					fmt.Printf("unexpected: %s\n", c.String())
				}
			}
			g = ng
		} else {
			ng := NewGame()
			ng.MaxX = fold.X - 1
			ng.MaxY = g.MaxY
			for _, c := range g.Coords {
				if c.X < fold.X {
					ng.Coords[c.String()] = c
				} else if c.X > fold.X {
					c.From = append(c.From, c)
					c.X = g.MaxX - c.X
					ng.Coords[c.String()] = c
				} else {
					fmt.Printf("unexpected: %s\n", c.String())
				}
			}
			g = ng
		}
	}
	if len(g.Coords) < 300 {
		g.Draw()
	}
	return len(g.Coords)
}

func SolveEasierSample() int {
	return solve("sample_input.txt", 1)
}

func SolveEasier() int {
	return solve("input.txt", 1)
}

func SolveHarderSample() int {
	return solve("sample_input.txt", 2)
}

func SolveHarder() int {
	return solve("input.txt", -1)
}
