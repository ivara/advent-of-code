package main

import (
	"bytes"
	"math/big"
	"sort"
	"strconv"
)

type FreshRange struct {
	start big.Int
	stop  big.Int
}

func (r *FreshRange) Contains(i big.Int) bool {
	result := i.Cmp(&r.start) >= 0 && i.Cmp(&r.stop) < 1
	// fmt.Printf("Is %v in range %v-%v: %v\n", &i, &r.start, &r.stop, result)
	return result
}

func part1(input []byte) int {
	data := bytes.Split(input, []byte("\n\n"))
	rangesRaw := bytes.Split(data[0], []byte{'\n'})
	ingredientsRaw := bytes.Split(data[1], []byte{'\n'})

	// The two critical lists to work with
	freshRanges := make([]FreshRange, len(rangesRaw))
	ingredients := make([]big.Int, len(ingredientsRaw))

	// Create list of ranges
	for i, r := range rangesRaw {
		d := bytes.Split(r, []byte{'-'})
		start := new(big.Int)
		start.SetString(string(d[0]), 10)
		stop := new(big.Int)
		stop.SetString(string(d[1]), 10)
		freshRanges[i] = FreshRange{start: *start, stop: *stop}
	}

	// Create list of ingredients
	for i := range ingredientsRaw {
		ingredient := new(big.Int)
		ingredient.SetString(string(ingredientsRaw[i]), 10)
		ingredients[i] = *ingredient
	}

	freshIngredients := 0
	// Loop over the ingredients to check which ones are fresh
	for i := range ingredients {
		for j := range freshRanges {
			if freshRanges[j].Contains(ingredients[i]) {
				freshIngredients += 1
				break
			}
		}
	}

	return freshIngredients
}

type Interval struct {
	Start int64
	End   int64
}

func part2(input []byte) int {

	crash := bytes.Split(input, []byte("\n\n"))
	boom := crash[0]
	bang := bytes.Split(boom, []byte{'\n'})

	ranges := make([]Interval, len(bang))
	for i := range bang { // lines
		values := bytes.Split(bang[i], []byte{'-'})
		start, _ := strconv.ParseInt(string(values[0]), 10, 64)
		end, _ := strconv.ParseInt(string(values[1]), 10, 64)
		ranges = append(ranges, Interval{Start: start, End: end})
	}

	// 1. Sort by start
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	// 2. Merge overlapping or adjacent ranges
	merged := make([]Interval, 0, len(ranges))
	current := ranges[0]

	for i := 1; i < len(ranges); i++ {
		r := ranges[i]

		if r.Start <= current.End+1 {
			// Overlap or directly adjacent → extend current
			if r.End > current.End {
				current.End = r.End
			}
		} else {
			// Disjoint → push current
			merged = append(merged, current)
			current = r
		}
	}
	merged = append(merged, current)
	// fmt.Printf("Merged: %v", merged)
	// 3. Count distinct integers
	var total int
	for _, m := range merged {
		if m.Start == m.End {
			continue
		}
		total += int(1 + m.End - m.Start)
	}

	// fmt.Println("Number of distinct integers:", total)
	return total
}
