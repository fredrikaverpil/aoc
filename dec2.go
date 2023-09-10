package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Hand struct {
	alias  string
	points int
}

func part1() {
	contents, _ := os.ReadFile("dec2.txt")
	r := strings.NewReader(string(contents))
	scanner := bufio.NewScanner(r)

	opponentPoints := 0
	myPoints := 0

	handMap := map[string]Hand{
		"A": {"rock", 1},
		"B": {"paper", 2},
		"C": {"scissors", 3},

		"X": {"rock", 1},
		"Y": {"paper", 2},
		"Z": {"scissors", 3},
	}

	pointsWin := 6
	pointsDraw := 3

	for scanner.Scan() {
		line := scanner.Text()
		characters := strings.Split(line, " ")
		opponentChar := characters[0]
		myChar := characters[1]

		opponentHand := handMap[opponentChar]
		myHand := handMap[myChar]

		if myHand.alias == "rock" {
			if opponentHand.alias == "rock" {
				// draw
				myPoints += myHand.points + pointsDraw
				opponentPoints += opponentHand.points + pointsDraw
			} else if opponentHand.alias == "scissors" {
				// win
				myPoints += myHand.points + pointsWin
				opponentPoints += opponentHand.points
			} else if opponentHand.alias == "paper" {
				// loose
				myPoints += myHand.points
				opponentPoints += opponentHand.points + pointsWin
			}
		} else if myHand.alias == "paper" {
			if opponentHand.alias == "paper" {
				// draw
				myPoints += myHand.points + pointsDraw
				opponentPoints += opponentHand.points + pointsDraw
			} else if opponentHand.alias == "rock" {
				// win
				myPoints += myHand.points + pointsWin
				opponentPoints += opponentHand.points
			} else if opponentHand.alias == "scissors" {
				// loose
				myPoints += myHand.points
				opponentPoints += opponentHand.points + pointsWin
			}
		} else if myHand.alias == "scissors" {
			if opponentHand.alias == "scissors" {
				// draw
				myPoints += myHand.points + pointsDraw
				opponentPoints += opponentHand.points + pointsDraw
			} else if opponentHand.alias == "paper" {
				// win
				myPoints += myHand.points + pointsWin
				opponentPoints += opponentHand.points
			} else if opponentHand.alias == "rock" {
				// loose
				myPoints += myHand.points
				opponentPoints += opponentHand.points + pointsWin
			}
		}

	}

	log.Printf("My points: %d", myPoints)
	log.Printf("Opponent points: %d", opponentPoints)
}

func part2() {
	contents, _ := os.ReadFile("dec2.txt")
	r := strings.NewReader(string(contents))
	scanner := bufio.NewScanner(r)

	opponentPoints := 0
	myPoints := 0

	rock := Hand{"rock", 1}
	paper := Hand{"paper", 2}
	scissors := Hand{"scissors", 3}

	handMap := map[string]Hand{
		"A": rock,
		"B": paper,
		"C": scissors,
	}

	resultMap := map[string]string{
		"X": "loose",
		"Y": "draw",
		"Z": "win",
	}

	pointsWin := 6
	pointsDraw := 3

	for scanner.Scan() {
		line := scanner.Text()
		characters := strings.Split(line, " ")
		opponentChar := characters[0]
		resultChar := characters[1]

		opponentHand := handMap[opponentChar]
		result := resultMap[resultChar]

		if result == "win" {
			if opponentHand.alias == "rock" {
				myPoints += paper.points + pointsWin
				opponentPoints += opponentHand.points
			} else if opponentHand.alias == "paper" {
				myPoints += scissors.points + pointsWin
				opponentPoints += opponentHand.points
			} else if opponentHand.alias == "scissors" {
				myPoints += rock.points + pointsWin
				opponentPoints += opponentHand.points
			}
		} else if result == "draw" {
			if opponentHand.alias == "rock" {
				myPoints += rock.points + pointsDraw
				opponentPoints += opponentHand.points + pointsDraw
			} else if opponentHand.alias == "paper" {
				myPoints += paper.points + pointsDraw
				opponentPoints += opponentHand.points + pointsDraw
			} else if opponentHand.alias == "scissors" {
				myPoints += scissors.points + pointsDraw
				opponentPoints += opponentHand.points + pointsDraw
			}
		} else if result == "loose" {
			if opponentHand.alias == "rock" {
				myPoints += scissors.points
				opponentPoints += opponentHand.points + pointsWin
			} else if opponentHand.alias == "paper" {
				myPoints += rock.points
				opponentPoints += opponentHand.points + pointsWin
			} else if opponentHand.alias == "scissors" {
				myPoints += paper.points
				opponentPoints += opponentHand.points + pointsWin
			}
		}

	}

	log.Printf("My points: %d", myPoints)
	log.Printf("Opponent points: %d", opponentPoints)
}

func main() {
	// part1()
	part2()
}
