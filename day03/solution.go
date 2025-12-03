package day03

import (
	"strconv"

	"github.com/fredrikaverpil/aoc/utils"
)

func getLargestJoltage(bank string, keep int) string {
	result := make([]byte, 0, keep)

	searchStart := 0
	for i := range keep {
		reserved := keep - i
		searchEnd := len(bank) - reserved

		var maxChar byte
		var maxIdx int

		for j := searchStart; j <= searchEnd; j++ {
			if bank[j] > maxChar {
				maxChar = bank[j]
				maxIdx = j
				if maxChar == '9' {
					break
				}
			}
		}

		result = append(result, maxChar)
		searchStart = maxIdx + 1
	}
	return string(result)
}

func SolvePart1(lines []string) int {
	keep := 2
	var joltageSum int
	for _, line := range lines {
		largestJoltage := getLargestJoltage(line, keep)
		joltage, err := strconv.Atoi(largestJoltage)
		if err != nil {
			panic(err)
		}
		joltageSum += joltage
	}

	return joltageSum
}

func SolvePart2(lines []string) int {
	keep := 12
	var joltageSum int
	for _, line := range lines {
		largestJoltage := getLargestJoltage(line, keep)
		joltage, err := strconv.Atoi(largestJoltage)
		if err != nil {
			panic(err)
		}
		joltageSum += joltage
	}

	return joltageSum
}

func Solve(path string) (int, int) {
	lines, err := utils.ReadLines(path)
	if err != nil {
		panic(err)
	}

	part1 := SolvePart1(lines)
	part2 := SolvePart2(lines)

	return part1, part2
}
