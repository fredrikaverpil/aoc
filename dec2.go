package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Player struct {
	Points int
}

type Shape struct {
	alias  string
	points int
}

func (p *Player) draw(s Shape) {
	drawPoints := 3
	p.Points += drawPoints + s.points
}

func (p *Player) win(s Shape) {
	winPoints := 6
	p.Points += winPoints + s.points
}

func (p *Player) loose(s Shape) {
	p.Points += s.points
}

func play(o Shape, p Shape, player *Player) {
	if o.alias == p.alias {
		player.draw(p)
	} else if (o.alias == "rock" && p.alias == "paper") ||
		(o.alias == "paper" && p.alias == "scissors") ||
		(o.alias == "scissors" && p.alias == "rock") {
		player.win(p)
	} else {
		player.loose(p)
	}
}

func part1() {
	contents, _ := os.ReadFile("dec2.txt")
	r := strings.NewReader(string(contents))
	scanner := bufio.NewScanner(r)

	player := Player{0}

	rock := Shape{"rock", 1}
	paper := Shape{"paper", 2}
	scissors := Shape{"scissors", 3}

	shapeMap := map[string]Shape{
		"A": rock,
		"B": paper,
		"C": scissors,

		"X": rock,
		"Y": paper,
		"Z": scissors,
	}

	for scanner.Scan() {
		line := scanner.Text()
		characters := strings.Split(line, " ")
		opponentChar := characters[0]
		myChar := characters[1]

		opponentShape := shapeMap[opponentChar]
		playerShape := shapeMap[myChar]

		play(opponentShape, playerShape, &player)
	}

	log.Printf("Player's points: %d", player.Points)
}

func part2() {
	contents, _ := os.ReadFile("dec2.txt")
	r := strings.NewReader(string(contents))
	scanner := bufio.NewScanner(r)

	opponentPoints := 0
	myPoints := 0

	rock := Shape{"rock", 1}
	paper := Shape{"paper", 2}
	scissors := Shape{"scissors", 3}

	handMap := map[string]Shape{
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
	part1()
	// part2()
}
