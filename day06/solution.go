package day06

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/fredrikaverpil/aoc/utils"
)

type problem struct {
	numbers  []int
	operator string
}

func (p *problem) solve() int {
	switch p.operator {
	case "+":
		sum := 0
		for _, n := range p.numbers {
			sum += n
		}
		return sum
	case "*":
		product := 1
		for _, n := range p.numbers {
			product *= n
		}
		return product
	default:
		panic("unexpected operator")
	}
}

type grid struct {
	cells [][]byte
	rows  int
	cols  int
}

// parseProblemsHorizontal prepares the problems to solve for part 1.
func parseProblemsHorizontal(lines []string) []problem {
	re := regexp.MustCompile(`\s+`)

	var numberRows [][]int
	var operators []string

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		tokens := re.Split(trimmed, -1) // no limit on number of splits

		if i < len(lines)-1 {
			// numbers only; do not parse last line which contains operators
			var row []int
			for _, token := range tokens {
				num, err := strconv.Atoi(token)
				if err != nil {
					panic(err)
				}
				row = append(row, num)
			}
			numberRows = append(numberRows, row)
		} else {
			operators = tokens
		}
	}

	var problems []problem
	for col, op := range operators {
		p := problem{operator: string(op[0])}
		for _, row := range numberRows {
			if col < len(row) {
				p.numbers = append(p.numbers, row[col])
			}
		}
		problems = append(problems, p)
	}

	return problems
}

// newGrid creates a 2d grid of integers, used to solve part 2.
func newGrid(lines []string) *grid {
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	cells := make([][]byte, maxLen)
	for i, line := range lines {
		cells[i] = make([]byte, maxLen)
		for j := range maxLen {
			if j < len(line) {
				cells[i][j] = line[j]
			}
		}
	}

	return &grid{cells: cells, rows: len(lines), cols: maxLen}
}

func (g *grid) printGrid() {
	for row := range g.rows {
		fmt.Println(string(g.cells[row]))
	}
}

func (g *grid) isColumnEmpty(col int) bool {
	for row := range g.rows {
		if g.cells[row][col] != ' ' {
			return false
		}
	}
	return true
}

func (g *grid) get(row, col int) byte {
	return g.cells[row][col]
}

func (g *grid) getColumnDigits(untilRow, col int) []byte {
	var digits []byte
	for row := range untilRow {
		char := g.cells[row][col]
		if char >= '0' && char <= '9' {
			digits = append(digits, char)
		}
	}

	return digits
}

// parseProblemsVertical prepares the problems to solve for part 2.
func parseProblemsVertical(g *grid) []problem {
	var problems []problem
	operatorRow := g.rows - 1
	col := 0

	for col < g.cols {
		for col < g.cols && g.isColumnEmpty(col) {
			col++
		}
		if col >= g.cols {
			break
		}

		p := problem{}
		for col < g.cols && !g.isColumnEmpty(col) {
			char := g.get(operatorRow, col)
			if char == '+' || char == '*' {
				p.operator = string(char)
			}

			digits := g.getColumnDigits(operatorRow, col)
			if len(digits) > 0 {
				num, err := strconv.Atoi(string(digits))
				if err != nil {
					panic(err)
				}
				p.numbers = append(p.numbers, num)
			}
			col++
		}
		problems = append(problems, p)
	}

	return problems
}

func sum(problems []problem) int {
	total := 0
	for _, p := range problems {
		total += p.solve()
	}
	return total
}

func part1(lines []string) int {
	problems := parseProblemsHorizontal(lines)
	return sum(problems)
}

func part2(lines []string) int {
	g := newGrid(lines)
	problems := parseProblemsVertical(g)
	return sum(problems)
}

func Solve(path string) (int, int) {
	lines, err := utils.ReadLines(path)
	if err != nil {
		panic(err)
	}

	return part1(lines), part2(lines)
}
