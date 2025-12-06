package day05_test

import (
	"fmt"
	"testing"

	"github.com/fredrikaverpil/aoc/day05"
)

func Example() {
	part1, part2 := day05.Solve("../testdata/day05_example.txt")
	fmt.Println(part1)
	fmt.Println(part2)
	// Output:
	// 3
	// 14
}

func TestPart1(t *testing.T) {
	want := 701
	got, _ := day05.Solve("../testdata/day05_input.txt")
	if got != want {
		t.Fatalf("got %d but want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 352340558684863
	_, got := day05.Solve("../testdata/day05_input.txt")
	if got != want {
		t.Fatalf("got %d but want %d", got, want)
	}
}

func BenchmarkSolve(b *testing.B) {
	for b.Loop() {
		day05.Solve("../testdata/day05_input.txt")
	}
}
