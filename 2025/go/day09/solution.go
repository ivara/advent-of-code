package main

import (
	"bytes"
)

func part1(input []byte) int {
	sum := 0

	lines := bytes.Split(input, []byte{'\n'})
	h := len(lines)
	for i := range h {
		a := bytes.Split(lines[i], []byte{','})
		A := Point{x: 1, y: 2}

		for j + i := range h {
			b := bytes.Split(lines[j], []byte{','})

		}
	}
	return sum
}

func part2(input []byte) int {
	sum := 0

	return sum
}

type Point struct {
	x int
	y int
}

func rectArea()
