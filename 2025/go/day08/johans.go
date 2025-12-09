package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	X, Y, Z float64
}

// --- Edge type ---
type Edge struct {
	A, B int
	D    float64
}

// --- DSU (Disjoint Set Union) ---
type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	p := make([]int, n)
	s := make([]int, n)
	for i := range p {
		p[i] = i
		s[i] = 1
	}
	return &DSU{parent: p, size: s}
}

func (d *DSU) Find(x int) int {
	for d.parent[x] != x {
		x = d.parent[x]
	}
	return x
}

func (d *DSU) Union(a, b int) bool {
	ra := d.Find(a)
	rb := d.Find(b)
	if ra == rb {
		return false
	}
	d.parent[rb] = ra
	d.size[ra] += d.size[rb]
	return true
}

// --- Distance function ---
func distance(a, b Point) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func part1johan(input []string) int {
	// --- Parse points ---
	points := make([]Point, len(input))
	for i, line := range input {
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(parts[0], 64)
		y, _ := strconv.ParseFloat(parts[1], 64)
		z, _ := strconv.ParseFloat(parts[2], 64)
		points[i] = Point{X: x, Y: y, Z: z}
	}

	// --- Build all edges ---
	var edges []Edge
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, Edge{
				A: i,
				B: j,
				D: distance(points[i], points[j]),
			})
		}
	}

	// --- Sort edges by distance ---
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].D < edges[j].D
	})

	// --- Connect the 1000 closest edges ---
	numConnections := 1000
	if numConnections > len(edges) {
		numConnections = len(edges)
	}
	dsu := NewDSU(len(points))
	for i := 0; i < numConnections; i++ {
		dsu.Union(edges[i].A, edges[i].B)
	}

	// --- Count sizes of all connected components ---
	groupSizes := make(map[int]int)
	for i := range points {
		root := dsu.Find(i)
		groupSizes[root]++
	}

	// --- Extract sizes and sort descending ---
	var sizes []int
	for _, s := range groupSizes {
		sizes = append(sizes, s)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	// --- Multiply the three largest ---
	answer := 1
	for i := 0; i < 3 && i < len(sizes); i++ {
		answer *= sizes[i]
	}

	return answer
}
