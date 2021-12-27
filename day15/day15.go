package day15

import (
	"advent/rkutil"
	"fmt"
	"math"
	"strconv"
)

type Game struct {
	Grid   [][]int
	Width  int
	Height int
}

func (g Game) ValidCoordinate(c Coordinate) bool {
	return c.X >= 0 && c.X < g.Width && c.Y >= 0 && c.Y < g.Height
}

func readGame(filename string) Game {
	lines := rkutil.MustLines(filename)
	g := Game{Grid: make([][]int, len(lines))}
	for y, line := range lines {
		for _, c := range line {
			v, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			g.Grid[y] = append(g.Grid[y], v)
		}
	}
	g.Width = len(g.Grid[0])
	g.Height = len(lines)
	return g
}

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

type Path struct {
	Coords  []Coordinate
	Cost    int
	Visited [][]bool
}

func NewPath(w, h int) Path {
	visited := make([][]bool, w*h)
	for i := 0; i < h; i++ {
		visited[i] = make([]bool, w)
	}
	return Path{Visited: visited}
}

func (p *Path) Add(g Game, c Coordinate) {
	p.Coords = append(p.Coords, c)
	p.Cost += g.Grid[c.X][c.Y]
	p.Visited[c.X][c.Y] = true
}

func (p *Path) Pop(g Game) {
	c := p.Coords[len(p.Coords)-1]
	p.Cost -= g.Grid[c.X][c.Y]
	p.Coords = p.Coords[0 : len(p.Coords)-1]
	p.Visited[c.X][c.Y] = false
}

func (p Path) IsVisited(check Coordinate) bool {
	return p.Visited[check.X][check.Y]
}

type Recorder struct {
	//Paths    []Path
	Cheapest int
}

func NewRecoder() *Recorder {
	return &Recorder{Cheapest: math.MaxInt}
}

func (rec *Recorder) RecordPath(path Path) {
	//rec.Paths = append(rec.Paths, path)
	if path.Cost < rec.Cheapest {
		rec.Cheapest = path.Cost
	}
}

var thrownOut int
var recorded int

func visit(g Game, cur Coordinate, path Path, rec *Recorder) {
	if path.IsVisited(cur) {
		//fmt.Printf("%s: already visited\n", cur.String())
		return
	}
	//fmt.Printf("%s: adding\n", cur.String())
	path.Add(g, cur)
	if path.Cost > rec.Cheapest {
		path.Pop(g)
		thrownOut++
		return
	}
	if cur.X == g.Width-1 && cur.Y == g.Height-1 {
		rec.RecordPath(path)
		recorded++
		return
	}
	right := Coordinate{X: cur.X + 1, Y: cur.Y}
	if g.ValidCoordinate(right) {
		//fmt.Printf("%s->%s: going right\n", cur.String(), left.String())
		visit(g, right, path, rec)
	}
	down := Coordinate{X: cur.X, Y: cur.Y + 1}
	if g.ValidCoordinate(down) {
		//fmt.Printf("%s->%s: going down\n", cur.String(), left.String())
		visit(g, down, path, rec)
	}
	left := Coordinate{X: cur.X - 1, Y: cur.Y}
	if g.ValidCoordinate(left) {
		//fmt.Printf("%s->%s: going left\n", cur.String(), left.String())
		visit(g, left, path, rec)
	}
	up := Coordinate{X: cur.X, Y: cur.Y - 1}
	if g.ValidCoordinate(up) {
		//fmt.Printf("%s->%s: going up\n", cur.String(), left.String())
		visit(g, up, path, rec)
	}
	path.Pop(g)
}

type Node struct {
	Dist    int
	Prev    Coordinate
	Visited bool
	Cost    int
}

type Nodes [][]Node

func (n Nodes) Dist(c Coordinate) int {
	return n[c.X][c.Y].Dist
}

func (n Nodes) Cost(c Coordinate) int {
	return n[c.X][c.Y].Cost
}

func (n Nodes) Prev(c Coordinate) Coordinate {
	return n[c.X][c.Y].Prev
}

func (n *Nodes) AssignDistance(c Coordinate, dist int) {
	(*n)[c.X][c.Y].Dist = dist
}

func (n *Nodes) AssignPrev(to Coordinate, assign Coordinate) {
	(*n)[to.X][to.Y].Prev = assign
}

func (n *Nodes) VisitCoordinate(c Coordinate) {
	(*n)[c.X][c.Y].Visited = true
}

func allNeighbours(g Game, c Coordinate) []Coordinate {
	var result []Coordinate
	if c.X > 0 {
		result = append(result, Coordinate{X: c.X - 1, Y: c.Y})
	}
	if c.X < g.Width-1 {
		result = append(result, Coordinate{X: c.X + 1, Y: c.Y})
	}
	if c.Y > 0 {
		result = append(result, Coordinate{X: c.X, Y: c.Y - 1})
	}
	if c.Y < g.Height-1 {
		result = append(result, Coordinate{X: c.X, Y: c.Y + 1})
	}
	return result
}

func neighboursInPlay(g Game, nodes Nodes, c Coordinate) []Coordinate {
	var result []Coordinate
	all := allNeighbours(g, c)
	for _, neighbour := range all {
		n := nodes[neighbour.X][neighbour.Y]
		if !n.Visited {
			result = append(result, neighbour)
		}
	}
	return result
}

// bruteForceMinDist is brute force scan
func bruteForceMinDist(g Game, nodes Nodes) Coordinate {
	minCoord := Coordinate{}
	minD := math.MaxInt
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			if !nodes[x][y].Visited {
				coord := Coordinate{X: x, Y: y}
				if nodes.Dist(coord) < minD {
					minCoord = coord
					minD = nodes.Dist(coord)
				}
			}
		}
	}
	return minCoord
}

func Dijkstra(g Game) int {
	total := g.Width * g.Height
	nodes := make(Nodes, g.Width*g.Height)
	for i := 0; i < g.Height; i++ {
		nodes[i] = make([]Node, g.Width)
		for j := 0; j < g.Width; j++ {
			nodes[i][j].Dist = math.MaxInt
			nodes[i][j].Cost = g.Grid[i][j]
		}
	}
	cur := Coordinate{}
	target := Coordinate{X: g.Width - 1, Y: g.Height - 1}
	nodes[0][0].Dist = 0
	for total >= 0 {
		cur = bruteForceMinDist(g, nodes)
		nodes.VisitCoordinate(cur)
		neighbours := neighboursInPlay(g, nodes, cur)
		if neighbours != nil {
			for _, neighbour := range neighbours {
				alt := nodes.Dist(cur) + nodes.Cost(neighbour)
				if alt < nodes.Dist(neighbour) {
					nodes.AssignDistance(neighbour, alt)
					nodes.AssignPrev(neighbour, cur)
				}
			}
		}
		if cur == target {
			break
		}
	}
	result := 0
	c := target
	for {
		if c.X == 0 && c.Y == 0 {
			break
		}
		result += nodes.Cost(c)
		c = nodes.Prev(c)
	}
	return result
}

func solveEasier(filename string) int {
	g := readGame(filename)
	return Dijkstra(g)
}

func SolveEasierSample() int {
	//return solveEasier("trivial_input.txt")
	return solveEasier("sample_input.txt")
}

func SolveEasier() int {
	return solveEasier("input.txt")
}

func solveHarder(filename string) int {
	g := readGame(filename)
	ng := Game{Grid: make([][]int, g.Width*5)}
	ng.Width = g.Width * 5
	ng.Height = g.Height * 5
	for i := 0; i < g.Width*5; i++ {
		ng.Grid[i] = make([]int, g.Height*5)
	}
	for y := 0; y < g.Height*5; y++ {
		for x := 0; x < g.Width*5; x++ {
			if x < g.Width && y < g.Height {
				ng.Grid[x][y] = g.Grid[x][y]
			} else if y < g.Height {
				v := ng.Grid[x-g.Width][y] + 1
				if v > 9 {
					v = 1
				}
				ng.Grid[x][y] = v
			} else {
				v := ng.Grid[x][y-g.Height] + 1
				if v > 9 {
					v = 1
				}
				ng.Grid[x][y] = v
			}
		}
	}
	return Dijkstra(ng)
}

func SolveHarderSample() int {
	return solveHarder("sample_input.txt")
}

func SolveHarder() int {
	return solveHarder("input.txt")
}
