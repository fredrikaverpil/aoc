package day01

import (
	"strconv"

	"github.com/fredrikaverpil/aoc/utils"
)

type dial struct {
	maxValues          int
	currentValue       int
	stoppedAtZeroCount int
	zeroCount          int
}

func newDial(maxValues, currentValue int) *dial {
	return &dial{maxValues: maxValues, currentValue: currentValue}
}

func (d *dial) turnLeft(steps int) {
	for range steps {
		d.click(-1)
	}
	if d.currentValue == 0 {
		d.stoppedAtZeroCount++
	}
}

func (d *dial) turnRight(steps int) {
	for range steps {
		d.click(1)
	}
	if d.currentValue == 0 {
		d.stoppedAtZeroCount++
	}
}

func (d *dial) click(step int) {
	if d.currentValue+step == 0 || d.currentValue+step == d.maxValues || d.currentValue+step == d.maxValues*-1 {
		d.currentValue = 0
		d.zeroCount++
	} else {
		d.currentValue += step
	}
}

func Solve(path string) (int, int) {
	maxValues := 100
	startValue := 50
	dial := newDial(maxValues, startValue)

	lines, err := utils.ReadLines(path)
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		direction := line[0:1]
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		switch direction {
		case "L":
			dial.turnLeft(steps)
		case "R":
			dial.turnRight(steps)
		default:
			panic("invalid direction: " + direction)
		}
	}

	return dial.stoppedAtZeroCount, dial.zeroCount
}
