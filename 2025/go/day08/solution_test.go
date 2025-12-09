package main

import (
	"log"
	"os"
	"testing"
)

// func TestCircuitContains(t *testing.T) {
// 	c := Circuit{1, 2, 3}

//		if !c.Contains(1) {
//			t.Errorf("Circuit contains number!")
//		}
//	}
func TestCircuitContains(t *testing.T) {
	tests := []struct {
		name     string
		circuit  Circuit
		contains int
		expected bool
	}{
		{"Circuit contains number 1", Circuit{1, 2, 3}, 1, true},
		{"Circuit does not contain number 1", Circuit{2, 3}, 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.circuit.Contains(tt.contains)
			if got != tt.expected {
				t.Errorf("%v failed\n", tt.name)
			}
		})
	}
}
func TestPart1(t *testing.T) {
	input := []byte(`162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`)
	got := part1(input, 10)
	want := 40

	if got != want {
		t.Errorf("Got = %v; want %v", got, want)
	}
}

func TestPart1RealData(t *testing.T) {
	input := readFileContents("input.txt")
	got := part1(input, 1001)
	want := 40

	if got != want {
		t.Errorf("Got = %v; want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := []byte(`162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`)
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
	input := []byte(`162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`)

	for b.Loop() {
		part1(input, 10)
	}
}

func BenchmarkPart2TestData(b *testing.B) {
	data := []byte(`162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`)

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
		part1(data, 1000)
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
