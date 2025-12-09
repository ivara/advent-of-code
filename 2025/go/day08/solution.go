package main

import (
	"bytes"
	"math"
	"slices"
	"strconv"
)

// Holds list of "index" from the input lines
// for which point is in this circuit
type Circuit []int
type Vec3 [3]float64
type Distance struct {
	i, j     int
	dist     float64
	iv3, jv3 Vec3
}

func (c *Circuit) Contains(p int) bool {
	return slices.Contains(*c, p)
}

func dist(a, b Vec3) float64 {
	dx, dy, dz := a[0]-b[0], a[1]-b[1], a[2]-b[2]

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func shortestPairDistances(points []Vec3, cap int) []Distance {
	n := len(points)
	if n < 2 {
		return nil
	}

	// Optimization: Create MinHeap with cap size
	size := Min(cap, n*(n-1)/2)
	dists := make([]Distance, 0, size)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			d := Distance{
				dist: dist(points[i], points[j]),
				i:    i,
				j:    j,
				iv3:  points[i],
				jv3:  points[j],
			}
			dists = append(dists, d)
		}
	}

	// Sort on lowest distance
	slices.SortFunc(dists, func(a, b Distance) int {
		if a.dist < b.dist {
			return -1
		} else if a.dist > b.dist {
			return 1
		}
		return 0
	})

	return dists[:cap]
}

func parsePoints(b []byte) []Vec3 {
	lines := bytes.Split(b, []byte{'\n'})
	points := make([]Vec3, len(lines))

	for i, l := range lines {
		numbers := bytes.Split(l, []byte{','})
		point := new(Vec3)
		for j := range numbers {
			v, _ := strconv.ParseFloat(string(numbers[j]), 64)
			point[j] = v
		}
		points[i] = *point
	}
	return points
}

func createCircuits(distances []Distance) []Circuit {
	circuits := []Circuit{}

	// TODO
	// BUG IDENTIFIED
	// if the current shortest distance is point A and B
	// and A is in circuit a and B is in circuit b
	// then circuit a and b shall be merged into one new circuit
outer:
	for _, d := range distances {
		for j, c := range circuits {
			if c.Contains(d.i) && c.Contains(d.j) {
				continue outer
			} else if c.Contains(d.i) {
				circuits[j] = append(circuits[j], d.j)
				continue outer
			} else if c.Contains(d.j) {
				circuits[j] = append(circuits[j], d.i)
				continue outer
			}
		}
		// if we reach this, no existing circuit was found!
		circuits = append(circuits, Circuit{d.i, d.j})
	}
	return circuits
}

// 2548 too low (real data)
func part1(input []byte, cap int) int {
	sum := 0

	// Get the points
	points := parsePoints(input)
	distances := shortestPairDistances(points, cap)
	circuits := createCircuits(distances)

	// sort circuits descending
	slices.SortFunc(circuits, func(a, b Circuit) int {
		if len(a) > len(b) {
			return -1
		} else if len(a) < len(b) {
			return 1
		}
		return 0
	})

	sum = len(circuits[0]) * len(circuits[1]) * len(circuits[2])

	// fmt.Printf("Distances: %v\n", distances)
	// Use MinHeap to hold top 1000 distances?
	// see ../../../2022/go/day03/solution.go for a MinHeap

	return sum
}

func part2(input []byte) int {
	sum := 0

	return sum
}
