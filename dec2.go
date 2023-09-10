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

func play2(o Shape, result string, player *Player, rock Shape, paper Shape, scissors Shape) {
	if result == "draw" {
		player.draw(o)
	} else if result == "win" {
		switch o.alias {
		case "rock":
			player.win(paper)
		case "paper":
			player.win(scissors)
		case "scissors":
			player.win(rock)
		}
	} else {
		// loose
		switch o.alias {
		case "rock":
			player.loose(scissors)
		case "paper":
			player.loose(rock)
		case "scissors":
			player.loose(paper)
		}
	}
}

func dec2part1() {
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

	contents, _ := os.ReadFile("dec2.txt")
	r := strings.NewReader(string(contents))
	scanner := bufio.NewScanner(r)

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

func dec2part2() {
	player := Player{0}

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

	contents, _ := os.ReadFile("dec2.txt")
	r := strings.NewReader(string(contents))
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		characters := strings.Split(line, " ")
		opponentChar := characters[0]
		resultChar := characters[1]

		opponentHand := handMap[opponentChar]
		result := resultMap[resultChar]

		play2(opponentHand, result, &player, rock, paper, scissors)

	}

	log.Printf("Player's points: %d", player.Points)
}

func main() {
	dec2part1()
	dec2part2()
}
