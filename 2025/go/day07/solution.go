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
	// sum := 0
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

	// Initialize cache
	cache := make(map[[2]int]int)
	return scoreMemoized(lines, maxRows, 1, startingColumn, cache)
}

func scoreMemoized(grid [][]byte, mr, r, c int, cache map[[2]int]int) int {
	// Create cache key
	key := [2]int{r, c}

	// Check if result already computed
	if result, exists := cache[key]; exists {
		return result
	}

	// Compute result
	nr := r + 1
	var result int

	if nr == mr {
		result = 1
	} else if grid[nr][c] == '^' {
		result = scoreMemoized(grid, mr, nr, c-1, cache) + scoreMemoized(grid, mr, nr, c+1, cache)
	} else {
		result = scoreMemoized(grid, mr, nr, c, cache)
	}

	// Store in cache and return
	cache[key] = result
	return result
}

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

func part2beta(input []byte) int {
	lines := bytes.Split(input, []byte{'\n'})
	beamsMap := make(map[int]int)

	// Find initial beam (S)
	for i := 0; i < len(lines[0]); i++ {
		if lines[0][i] == 'S' {
			beamsMap[i] = 1
			break
		}
	}

	// Start working the field!
	for i := 1; i < len(lines); i++ {
		// See if there is a splitter on current paths
		// bIdx is the value of a column where the beam is currently travelling

		beamsCopy := getBeamsFromMapV2(beamsMap) // make sure to get rid of this!
		for j := range beamsCopy {
			if lines[i][beamsCopy[j]] == '^' {
				// If we hit a splitter in part 2
				// continue left, but also start a new instance going right
				// remove current beam
				beamsMap[beamsCopy[j]] -= 1
				beamsMap[beamsCopy[j]-1] += 1
				beamsMap[beamsCopy[j]+1] += 1
			} else {

			}
		}
	}

	// Get sum "current beams"
	sum := 0
	for _, v := range beamsMap {
		if v > 0 {
			sum += v
		}
	}
	return sum
}
