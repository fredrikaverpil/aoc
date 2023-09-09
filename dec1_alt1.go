package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("dec1.txt")
	sc := bufio.NewScanner(file)

	maxWeight := 0
	currentWeights := 0

	for sc.Scan() {
		line := sc.Text()

		if line == "" {
			if currentWeights > maxWeight {
				maxWeight = currentWeights
			}
			currentWeights = 0
		} else {
			weight, _ := strconv.Atoi(line)
			currentWeights += weight
		}

	}

	log.Printf("Largest weight: %d", maxWeight)
}
