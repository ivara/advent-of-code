package main

import (
	"log"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	input := []byte(`123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   + `)
	got := part1(input)
	want := 4277556

	if got != want {
		t.Errorf("Got = %v; want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := []byte(`123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   + `)
	got := part2(input)
	want := 3263827

	if got != want {
		t.Errorf("Got = %v; want %v", got, want)
	}
}

// ---------------------------------------------------
//
//	BENCHMARKS
//
// ---------------------------------------------------
func BenchmarkPart1TestData(b *testing.B) {
	input := []byte(`3-5
10-14
16-20
12-18

1
5
8
11
17
32`)

	for b.Loop() {
		part1(input)
	}
}

func BenchmarkPart1RealData(b *testing.B) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	for b.Loop() {
		part1(data)
	}
}

func BenchmarkPart2RealData(b *testing.B) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	for b.Loop() {
		part2(data)
	}
}
