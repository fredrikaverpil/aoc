package day05

import (
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/fredrikaverpil/aoc/utils"
)

type interval struct {
	start int
	end   int
}

// getFreshIntervals creates intervals, sorts them and merges any overlaps.
func getFreshIntervals(ranges []string) []interval {
	var intervals []interval

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		intervals = append(intervals, interval{start: start, end: end})
	}

	slices.SortFunc(intervals, func(a, b interval) int {
		return a.start - b.start
	})

	merged := []interval{intervals[0]}
	for _, current := range intervals[1:] {
		last := &merged[len(merged)-1] // get pointer, do not get a copy of the last interval (or mutations won't persist)
		if current.start <= last.end+1 {
			if current.end > last.end {
				last.end = current.end
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged
}

// isFresh checks if an ingredient ID is within any of the fresh intervals.
func isFresh(id int, intervals []interval) bool {
	// Use binary search to find the first interval that ends at or after the ID.
	// sort.Search requires the condition to be monotonic: [False, False, ..., True, True].
	// Here, intervals with end < id return False, and end >= id return True.
	idx := sort.Search(len(intervals), func(i int) bool { return intervals[i].end >= id })
	// sort.Search returns len(intervals) if the condition is never true.
	// This means 'id' is greater than the 'end' of all intervals, so it cannot be inside any of them.
	if idx == len(intervals) {
		return false
	}
	// The interval was found where end >= id. Check if start <= id to ensure the ID is inside.
	return intervals[idx].start <= id
}

func countFreshIngredients(intervals []interval, ingredientIDs []int) int {
	count := 0

	for _, ingredientID := range ingredientIDs {
		if isFresh(ingredientID, intervals) {
			count++
		}
	}
	return count
}

func countPossibleFreshIngredients(intervals []interval) int {
	count := 0

	for _, interval := range intervals {
		count += (interval.end - interval.start) + 1
	}
	return count
}

func Solve(path string) (int, int) {
	lines, err := utils.ReadLines(path)
	if err != nil {
		panic(err)
	}

	emptyLine := slices.Index(lines, "")
	if emptyLine == -1 {
		panic("no empty line found in input")
	}

	ranges := lines[:emptyLine]
	ingredients := lines[emptyLine+1:]

	freshIDs := getFreshIntervals(ranges)
	ingredientIDs := make([]int, len(ingredients))
	for i, ingredient := range ingredients {
		converted, err := strconv.Atoi(ingredient)
		if err != nil {
			panic(err)
		}
		ingredientIDs[i] = converted
	}

	part1 := countFreshIngredients(freshIDs, ingredientIDs)
	part2 := countPossibleFreshIngredients(freshIDs)

	return part1, part2
}
