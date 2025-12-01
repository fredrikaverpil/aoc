package day01_test

import (
	"fmt"
	"testing"

	"github.com/fredrikaverpil/aoc/day01"
)

func Example() {
	part1, part2 := day01.Solve("../testdata/day01_example.txt")
	fmt.Println(part1)
	fmt.Println(part2)
	// Output:
	// 3
	// 6
}

func TestPart1(t *testing.T) {
	want := 1177
	got, _ := day01.Solve("../testdata/day01_input.txt")
	if got != want {
		t.Fatalf("got %d but want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 6768
	_, got := day01.Solve("../testdata/day01_input.txt")
	if got != want {
		t.Fatalf("got %d but want %d", got, want)
	}
}

func BenchmarkSolve(b *testing.B) {
	for b.Loop() {
		day01.Solve("../testdata/day01_input.txt")
	}
}
