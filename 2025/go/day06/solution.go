package main

import (
	"bytes"
	"strconv"
)

func part1(input []byte) int {
	sum := 0
	lines := bytes.Split(input, []byte{'\n'})

	// numbers := make([][]int, 5)
	numbers := make(map[int][]int)

	// Gather the numbers, loop all lines except the last
	for i := 0; i < len(lines)-1; i++ {
		fields := bytes.Fields(lines[i])

		for j, f := range fields {
			number, _ := strconv.Atoi(string(f))
			numbers[j] = append(numbers[j], number)
		}
	}

	// Do the calculus
	lastLine := lines[len(lines)-1]
	operations := bytes.Fields(lastLine)
	for i, o := range operations {
		toAdd := 0
		if string(o) == "+" {
			for _, y := range numbers[i] {
				toAdd += y
			}
		} else if string(o) == "*" {
			toAdd = 1
			for _, y := range numbers[i] {
				toAdd *= y
			}
		}
		sum += toAdd
	}
	return sum
}

func part2(input []byte) int {
	sum := 0
	lines := bytes.Split(input, []byte{'\n'})

	// numbers := make([][]int, 5)
	numbers := make(map[int][]int)

	// First "invert" the matrix
	// Then parse fields, last field will be operator!

	// Gather the numbers, loop all lines except the last
	for i := 0; i < len(lines)-1; i++ {
		fields := bytes.Fields(lines[i])

		for j, f := range fields {
			number, _ := strconv.Atoi(string(f))
			numbers[j] = append(numbers[j], number)
		}
	}

	// Figure out the REAL numbers (vertical column wise)
	for i := 0; i < len(numbers); i++ {
		// first figure out the longest number in this column, that determines everything }
	}

	lastLine := lines[len(lines)-1]
	operations := bytes.Fields(lastLine)
	for i, o := range operations {
		toAdd := 0
		if string(o) == "+" {
			for _, y := range numbers[i] {
				toAdd += y
			}
		} else if string(o) == "*" {
			toAdd = 1
			for _, y := range numbers[i] {
				toAdd *= y
			}
		}
		sum += toAdd
	}
	return sum
}
