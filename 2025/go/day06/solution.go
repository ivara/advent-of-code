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

// w = width (column)
// h = height (lines)
func isBlankColumn(lines []byte, column, h int) bool {
	for i := 0; i < h; i++ {
		if lines[i][column] != []byte{' '} {
			return false
		}
	}
	return true
}

func multiply(numbers []int) {
	sum := 0
	for _, n := range numbers {
		sum *= n
	}
	return sum
}

func add(numbers []int) {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func part2(input []byte) int {
	sum := 0
	lines := bytes.Split(input, []byte{'\n'})

	width := len(lines[0])
	height := len(lines)

	currentNumbers := []int
	currentOperator := "*"

	// Loop over all columns, and for each column
	// Walk the lines to build numbers
	for c := 0; c < width; c++ { // Columns
		// the number we are currently building for this colum
		var currentNumber []byte
		for h := 0; h < height; h++ { // Lines
			if isBlankLine(lines, c, h) {
				// this means "current block" is done
				switch currentOperator {
				case "*":
					sum += multiply(currentNumbers)
					break
				case "+":
					sum += add(currentNumbers)
					break
				}

				// clear things up for next block
				currentOperator := ""
				currentNumbers = nil
			} else {
				// build a number!
				if operator == "" {
					// Oh noes! no operator, that means it is in this
					operator = lines[height-1][c]
				}
				number := getNumberFromColumn(lines, c)
				currentNumbers = append(currentNumbers, number)
			}
		}
	}

	return sum
}

func getNumberFromColumn(lines []byte, column c) int {
	return 0
}
