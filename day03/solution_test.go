package day03_test

import (
	"fmt"
	"testing"

	"github.com/fredrikaverpil/aoc/day03"
)

func Example() {
	part1, part2 := day03.Solve("../testdata/day03_example.txt")
	fmt.Println(part1)
	fmt.Println(part2)
	// Output:
	// 357
	// 3121910778619
}

func TestPart1(t *testing.T) {
	want := 17229
	got, _ := day03.Solve("../testdata/day03_input.txt")
	if got != want {
		t.Fatalf("got %d but want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 170520923035051
	_, got := day03.Solve("../testdata/day03_input.txt")
	if got != want {
		t.Fatalf("got %d but want %d", got, want)
	}
}

func BenchmarkSolve(b *testing.B) {
	for b.Loop() {
		day03.Solve("../testdata/day03_input.txt")
	}
}
