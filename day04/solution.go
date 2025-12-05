package day04

import (
	"github.com/fredrikaverpil/aoc/utils"
)

const maxNeighbours = 4

// func printGrid(grid [][]byte) {
// 	for _, row := range grid {
// 		fmt.Println(string(row))
// 	}
// }

func countNeighbors(grid [][]byte, x, y, rows, cols int) int {
	directions := [][2]int{
		{-1, -1}, // upper left
		{0, -1},  // up
		{1, -1},  // upper right
		{-1, 0},  // left
		// {0, 0}, // the roll, '@'
		{1, 0},  // right
		{-1, 1}, // lower left
		{0, 1},  // down
		{1, 1},  // lower right
	}

	count := 0
	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]

		insideRowBoundary := ny >= 0 && ny <= rows-1
		insideColBoundary := nx >= 0 && nx <= cols-1

		if insideColBoundary && insideRowBoundary {
			if grid[ny][nx] == '@' {
				count++
			}
		}
	}

	return count
}

func countAccessible(grid [][]byte, rows, cols int) int {
	count := 0

	for y := range rows {
		for x := range cols {
			if grid[y][x] == '@' {
				neighbors := countNeighbors(grid, x, y, rows, cols)
				if neighbors < maxNeighbours {
					count++
				}
			}
		}
	}

	// printGrid(grid)

	return count
}

func countRollsRemoved(grid [][]byte, rows, cols int) int {
	count := 0
	iterate := true

	for iterate {
		iterate = false

		for y := range rows {
			for x := range cols {
				if grid[y][x] == '@' {
					neighbors := countNeighbors(grid, x, y, rows, cols)
					if neighbors < maxNeighbours {
						count++
						iterate = true
						grid[y][x] = 'x'
					}
				}
			}
		}
	}

	// printGrid(grid)

	return count
}

func Solve(path string) (int, int) {
	lines, err := utils.ReadLines(path)
	if err != nil {
		panic(err)
	}

	rows := len(lines)
	cols := len(lines[0])
	grid := make([][]byte, rows)
	for col, line := range lines {
		grid[col] = []byte(line)
	}

	part1 := countAccessible(grid, rows, cols)
	part2 := countRollsRemoved(grid, rows, cols)

	return part1, part2
}
