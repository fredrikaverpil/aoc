//nolint:forbidigo // CLI output requires fmt.Print
package main

import (
	"fmt"

	"github.com/fredrikaverpil/aoc/day01"
)

func main() {
	fmt.Println("Advent of Code 2025")
	fmt.Println("===================")
	fmt.Println()

	runDay01()
}

func runDay01() {
	part1, part2 := day01.Solve("testdata/day01_input.txt")
	fmt.Println("Day 1: Secret Entrance")
	fmt.Printf("  Part 1: %d\n", part1)
	fmt.Printf("  Part 2: %d\n", part2)
	fmt.Println()
}
