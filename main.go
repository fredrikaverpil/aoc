//nolint:forbidigo // CLI output requires fmt.Print
package main

import (
	"fmt"

	"github.com/fredrikaverpil/aoc/day01"
	"github.com/fredrikaverpil/aoc/day02"
	"github.com/fredrikaverpil/aoc/day03"
	"github.com/fredrikaverpil/aoc/day04"
	"github.com/fredrikaverpil/aoc/day05"
	"github.com/fredrikaverpil/aoc/day06"
)

func main() {
	fmt.Println("Advent of Code 2025")
	fmt.Println("===================")
	fmt.Println()

	runDay01()
	runDay02()
	runDay03()
	runDay04()
	runDay05()
	runDay06()
}

func runDay01() {
	part1, part2 := day01.Solve("testdata/day01_input.txt")
	fmt.Println("Day 1: Secret Entrance")
	fmt.Printf("  Part 1: %d\n", part1)
	fmt.Printf("  Part 2: %d\n", part2)
	fmt.Println()
}

func runDay02() {
	part1, part2 := day02.Solve("testdata/day02_input.txt")
	fmt.Println("Day 2: Gift Shop")
	fmt.Printf("  Part 1: %d\n", part1)
	fmt.Printf("  Part 2: %d\n", part2)
	fmt.Println()
}

func runDay03() {
	part1, part2 := day03.Solve("testdata/day03_input.txt")
	fmt.Println("Day 3: Lobby")
	fmt.Printf("  Part 1: %d\n", part1)
	fmt.Printf("  Part 2: %d\n", part2)
	fmt.Println()
}

func runDay04() {
	part1, part2 := day04.Solve("testdata/day04_input.txt")
	fmt.Println("Day 4: Printing Department")
	fmt.Printf("  Part 1: %d\n", part1)
	fmt.Printf("  Part 2: %d\n", part2)
	fmt.Println()
}

func runDay05() {
	part1, part2 := day05.Solve("testdata/day05_input.txt")
	fmt.Println("Day 5: Cafeteria")
	fmt.Printf("  Part 1: %d\n", part1)
	fmt.Printf("  Part 2: %d\n", part2)
	fmt.Println()
}

func runDay06() {
	part1, part2 := day06.Solve("testdata/day06_input.txt")
	fmt.Println("Day 6: Trash Compactor")
	fmt.Printf("  Part 1: %d\n", part1)
	fmt.Printf("  Part 2: %d\n", part2)
	fmt.Println()
}
