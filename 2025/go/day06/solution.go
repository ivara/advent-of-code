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

	width := len(lines[0])
	height := len(lines)

	var currentNumbers []int
	currentOperator := ""

	// Loop over all columns, and for each column
	// Walk the lines to build numbers
	for columnIdx := range width { // Columns
		if isBlankColumn(lines, columnIdx) {
			// this means "current block" is done
			switch currentOperator {
			case "*":
				sum += multiply(currentNumbers)
			case "+":
				sum += add(currentNumbers)
			}

			// clear things up for next block
			currentOperator = ""
			currentNumbers = nil
		} else {
			if currentOperator == "" {
				// Oh noes! no operator, that means it is in this
				currentOperator = string(lines[height-1][columnIdx])
			}
			number := getNumberFromColumn(lines, columnIdx)
			currentNumbers = append(currentNumbers, number)
		}
	}

	switch currentOperator {
	case "*":
		sum += multiply(currentNumbers)
	case "+":
		sum += add(currentNumbers)
	}

	return sum
}

func getNumberFromColumn(lines [][]byte, columnIdx int) int {
	var bytes []byte
	// last byte is blank or operand
	stop := len(lines) - 1
	for i := range stop {
		if lines[i][columnIdx] != ' ' {
			bytes = append(bytes, lines[i][columnIdx])
		}
	}
	number, _ := strconv.Atoi(string(bytes))
	return number
}

func isBlankColumn(lines [][]byte, columnIdx int) bool {
	for i := range lines {
		if lines[i][columnIdx] != ' ' {
			return false
		}
	}
	return true
}

func multiply(numbers []int) int {
	sum := 1
	for _, n := range numbers {
		sum *= n
	}
	return sum
}

func add(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}
