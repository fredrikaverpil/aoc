package day06_test

import (
	"fmt"
	"testing"

	"github.com/fredrikaverpil/aoc/day06"
)

func Example() {
	part1, part2 := day06.Solve("../testdata/day06_example.txt")
	fmt.Println(part1)
	fmt.Println(part2)
	// Output:
	// 4277556
	// 3263827
}

func TestPart1(t *testing.T) {
	want := 5322004718681
	got, _ := day06.Solve("../testdata/day06_input.txt")
	if got != want {
		t.Fatalf("got %d but want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 9876636978528
	_, got := day06.Solve("../testdata/day06_input.txt")
	if got != want {
		t.Fatalf("got %d but want %d", got, want)
	}
}

func BenchmarkSolve(b *testing.B) {
	for b.Loop() {
		day06.Solve("../testdata/day06_input.txt")
	}
}
