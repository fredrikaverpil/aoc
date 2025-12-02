package day02

import (
	"slices"
	"strconv"
	"strings"

	"github.com/fredrikaverpil/aoc/utils"
)

func checkForRepetition(i int) int {
	if len(strconv.Itoa(i))%2 != 0 {
		return 0
	}

	s := strconv.Itoa(i)
	middle := len(s) / 2
	left := s[:middle]
	right := s[middle:]

	if left != right {
		return 0
	}

	return i
}

func checkForRepetition2(i int) int {
	numberAsString := strconv.Itoa(i)
	var repetitions []string

	for i := range len(numberAsString) {
		partial := numberAsString[0:i]
		candidate := partial
		for i := len(candidate); i < len(numberAsString); i++ {
			candidate += partial
			if candidate == numberAsString {
				if !slices.Contains(repetitions, candidate) {
					repetitions = append(repetitions, candidate)
				}
			}
		}
	}

	ret := 0
	for _, repetition := range repetitions {
		num, _ := strconv.Atoi(repetition)
		ret += num
	}

	return ret
}

func Solve(path string) (int, int) {
	lines, err := utils.ReadLines(path)
	if err != nil {
		panic(err)
	}

	part1 := 0
	part2 := 0
	for r := range strings.SplitSeq(lines[0], ",") {
		parts := strings.Split(r, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for i := start; i <= end; i++ {
			part1 += checkForRepetition(i)
		}

		for i := start; i <= end; i++ {
			part2 += checkForRepetition2(i)
		}
	}

	return part1, part2
}
