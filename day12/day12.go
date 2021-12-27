package day12

import (
	"advent/rkutil"
	"sort"
	"strings"
)

type Segment struct {
	Lhs string
	Rhs string
}

type Node struct {
	Links      map[string]*Node
	S          string
	Visited    bool
	VisitTwice bool
}

func NewNode(s string) *Node {
	return &Node{Links: make(map[string]*Node), S: s}
}

type Graph struct {
	Start   *Node
	NodeMap map[string]*Node
}

func (g *Graph) UpsertNode(s string) *Node {
	if lhs, ok := g.NodeMap[s]; ok {
		return lhs
	}
	n := NewNode(s)
	g.NodeMap[s] = n
	return n
}

func readInput(filename string) *Graph {
	lines := rkutil.MustLines(filename)
	g := Graph{NodeMap: make(map[string]*Node)}
	for _, line := range lines {
		s := strings.Split(line, "-")
		rkutil.Ensure(len(s) == 2, "expecting two values")
		lhs := g.UpsertNode(s[0])
		rhs := g.UpsertNode(s[1])
		lhs.Links[rhs.S] = rhs
		rhs.Links[lhs.S] = lhs
	}
	return &g
}

func toString(q []*Node) string {
	var sb strings.Builder
	for _, e := range q {
		sb.WriteString(strings.ReplaceAll(e.S, "2", "") + ",")
	}
	sb.WriteString("end")
	return sb.String()
}

func visitEasier(q []*Node, n *Node, rec *resultRecorder) {
	for _, m := range n.Links {
		if m.S == "end" {
			rec.Record(toString(q))
		} else {
			if !m.Visited {
				q = append(q, m)
				if rkutil.IsLower(m.S) {
					m.Visited = true
				}
				visitEasier(q, m, rec)
				q = q[0 : len(q)-1]
				m.Visited = false
			}
		}
	}
}

func solveEasier(filename string) int {
	g := readInput(filename)
	n := g.NodeMap["start"]
	n.Visited = true
	var q []*Node
	q = append(q, n)
	var rec resultRecorder
	visitEasier(q, n, &rec)
	return len(rec.results)
}

func SolveEasierSample() int {
	return solveEasier("sample_input.txt")
}

func SolveEasier() int {
	return solveEasier("input.txt")
}

type resultRecorder struct {
	results []string
}

func (rec *resultRecorder) Record(s string) {
	rec.results = append(rec.results, s)
}

func (rec *resultRecorder) Final() []string {
	final := make(map[string]bool)
	for _, result := range rec.results {
		final[result] = true
	}
	var sorted []string
	for k := range final {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	return sorted
}

func (rec *resultRecorder) ResultDump() string {
	return strings.Join(rec.Final(), "\n")
}

func solveHarder(filename string) int {
	g := readInput(filename)
	var lowerNodes []*Node
	for k, v := range g.NodeMap {
		if k != "start" && k != "end" && rkutil.IsLower(v.S) {
			lowerNodes = append(lowerNodes, v)
		}
	}
	var rec resultRecorder
	for _, lowerNode := range lowerNodes {
		newNode := NewNode(lowerNode.S + "2")
		for _, link := range lowerNode.Links {
			newNode.Links[link.S] = link
			link.Links[lowerNode.S+"2"] = newNode
		}
		for _, v := range g.NodeMap {
			v.Visited = false
			if lowerNode.S == v.S {
				v.VisitTwice = true
			}
		}
		n := g.NodeMap["start"]
		n.Visited = true
		var q []*Node
		q = append(q, n)
		visitEasier(q, n, &rec)
		for _, link := range newNode.Links {
			delete(link.Links, newNode.S)
		}
	}
	return len(rec.Final())
}

func SolveHarderSample() int {
	return solveHarder("sample_input.txt")
}

func SolveHarder() int {
	return solveHarder("input.txt")
}
