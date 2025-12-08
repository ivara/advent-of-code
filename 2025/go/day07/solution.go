package main

import (
	"bytes"
)

// Om du träffar en splitter försvinner du
// om du träffar en splitter blir det nya beams på -1 och +1
// Spara en lista på currentBeams
func part1(input []byte) int {
	splits := 0
	lines := bytes.Split(input, []byte{'\n'})
	beamsMap := make(map[int]bool)

	// Find initial beam (S)
	for i := 0; i < len(lines[0]); i++ {
		if lines[0][i] == 'S' {
			beamsMap[i] = true
			break
		}
	}

	// Start working the field!
	for i := 1; i < len(lines); i++ {
		// See if there is a splitter on current paths
		// bIdx is the value of a column where the beam is currently travelling
		beamsCopy := getBeamsFromMap(beamsMap)
		for j := range beamsCopy {
			if lines[i][beamsCopy[j]] == '^' {
				splits += 1
				beamsMap[beamsCopy[j]] = false
				beamsMap[beamsCopy[j]-1] = true
				beamsMap[beamsCopy[j]+1] = true
				// append -1 and +1 to beams
				// but also remove current beam
				// Remove element at j, then add the two new beams
				// beams = append(beams[:j], append(beams[j+1:], tmp[j]-1, tmp[j]+1)...)
				// beams = removeAtIndex(beams, j)
				// beams = append(beams, tmp[j]-1, tmp[j]+1)

				// remove beam since it got splitted
				// beams = beams[:]
			}
		}
	}

	return splits
}

func getBeamsFromMap(m map[int]bool) []int {
	count := 0
	for _, value := range m {
		if value {
			count++
		}
	}

	var result = make([]int, 0, count)
	for key, value := range m {
		if value == true {
			result = append(result, key)
		}
	}

	return result
}

func getBeamsFromMapV2(m map[int]int) []int {
	count := 0
	for _, value := range m {
		if value > 0 {
			count++
		}
	}

	var result = make([]int, 0, count)
	for key, value := range m {
		if value > 0 {
			result = append(result, key)
		}
	}

	return result
}

func part2(input []byte) int {
	sum := 0
	lines := bytes.Split(input, []byte{'\n'})
	maxRows := len(lines)
	startingColumn := 0
	// Find initial beam (S)
	for i := 0; i < len(lines[0]); i++ {
		if lines[0][i] == 'S' {
			startingColumn = i
			break
		}
	}

	sum = score(lines, maxRows, 0, startingColumn)
	// Start working the field!
	// for i := 1; i < len(lines); i++ {
	// 	tmp := make([]int, 0, 1000)
	// 	for _, col := range currentBeams {
	// 		if lines[i][col] == '^' {
	// 			tmp = append(tmp, col-1, col+1)
	// 		} else {
	// 			tmp = append(tmp, col)
	// 		}
	// 	}
	// 	currentBeams = tmp
	// }

	// Get sum "current beams"
	// sum := 0
	// for _, v := range beamsMap {
	// 	if v > 0 {
	// 		sum += v
	// 	}
	// }
	return sum
}

// How to handle split that goes off grid? (left/right)
func score(grid [][]byte, mr, r, c int) int {
	// next row
	nr := r + 1

	// mr = max rows (height of grid)
	if nr == mr {
		return 1
	}
	if grid[nr][c] == '^' {
		return score(grid, mr, nr, c-1) + score(grid, mr, nr, c+1)
	} else {
		return score(grid, mr, nr, c)
	}
}

func part2gpt(input []byte) int {
	lines := bytes.Split(input, []byte{'\n'})

	// Find initial beam (S)
	startCol := 0
	for i := 0; i < len(lines[0]); i++ {
		if lines[0][i] == 'S' {
			startCol = i
			break
		}
	}

	// Count all paths using BFS/queue approach
	type Beam struct {
		row int
		col int
	}

	queue := []Beam{{row: 0, col: startCol}}
	pathCount := 0

	for len(queue) > 0 {
		beam := queue[0]
		queue = queue[1:]

		// Move beam down until it hits bottom or a splitter
		for beam.row < len(lines)-1 {
			beam.row++

			// Check if we hit a splitter
			if lines[beam.row][beam.col] == '^' {
				// Split into left and right
				leftBeam := Beam{row: beam.row, col: beam.col - 1}
				rightBeam := Beam{row: beam.row, col: beam.col + 1}
				queue = append(queue, leftBeam, rightBeam)
				break // This beam path ends here (it split)
			}
		}

		// If we reached the bottom without splitting, count this path
		if beam.row == len(lines)-1 {
			pathCount++
		}
	}

	return pathCount
}
