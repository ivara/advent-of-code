package main

import (
	"log"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	input := []byte(`7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`)
	got := part1(input)
	want := 50

	if got != want {
		t.Errorf("Got = %v; want %v", got, want)
	}
}

func TestPart1RealData(t *testing.T) {
	input := readFileContents("input.txt")
	got := part1(input)
	want := 40

	if got != want {
		t.Errorf("Got = %v; want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := []byte(`7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`)
	got := part2(input)
	want := -1

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
	input := []byte(`7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`)

	for b.Loop() {
		part1(input)
	}
}

func BenchmarkPart2TestData(b *testing.B) {
	data := []byte(`7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`)

	for b.Loop() {
		part2(data)
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
