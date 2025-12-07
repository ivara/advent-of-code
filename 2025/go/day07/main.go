package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	input := readFileContents("input.txt")

	startTime := time.Now()
	result := part1(input)
	part1time := float64(time.Since(startTime).Microseconds()) / 1000

	startTime = time.Now()
	result2 := part2(input)
	part2time := float64(time.Since(startTime).Microseconds()) / 1000

	fmt.Printf("Result part1: %v in %vms\n", result, part1time)
	fmt.Printf("Result part2: %v in %v ms\n ", result2, part2time)
}

func readFileContents(filePath string) []byte {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	return data
}
