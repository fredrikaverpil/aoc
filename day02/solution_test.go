package day02_test

import (
	"fmt"
	"testing"

	"github.com/fredrikaverpil/aoc/day02"
)

func Example() {
	part1, part2 := day02.Solve("../testdata/day02_example.txt")
	fmt.Println(part1)
	fmt.Println(part2)
	// Output:
	// 1227775554
	// 4174379265
}

func TestPart1(t *testing.T) {
	want := 23560874270
	got, _ := day02.Solve("../testdata/day02_input.txt")
	if got != want {
		t.Fatalf("got %d but want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 44143124633
	_, got := day02.Solve("../testdata/day02_input.txt")
	if got != want {
		t.Fatalf("got %d but want %d", got, want)
	}
}

func BenchmarkSolve(b *testing.B) {
	for b.Loop() {
		day02.Solve("../testdata/day02_input.txt")
	}
}
