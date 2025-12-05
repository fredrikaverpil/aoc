package day04_test

import (
	"fmt"
	"testing"

	"github.com/fredrikaverpil/aoc/day04"
)

func Example() {
	part1, part2 := day04.Solve("../testdata/day04_example.txt")
	fmt.Println(part1)
	fmt.Println(part2)
	// Output:
	// 13
	// 43
}

func TestPart1(t *testing.T) {
	want := 1518
	got, _ := day04.Solve("../testdata/day04_input.txt")
	if got != want {
		t.Fatalf("got %d but want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 8665
	_, got := day04.Solve("../testdata/day04_input.txt")
	if got != want {
		t.Fatalf("got %d but want %d", got, want)
	}
}

func BenchmarkSolve(b *testing.B) {
	for b.Loop() {
		day04.Solve("../testdata/day04_input.txt")
	}
}
