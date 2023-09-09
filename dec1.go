package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFile(filename string) []byte {
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return contents
}

func fileContentsToMap(contents []byte) map[int][]int {
	lines := strings.Split(string(contents), "\n")
	contentMap := make(map[int][]int)

	key := 0
	for _, line := range lines {
		if line == "" {
			key++
		} else {
			lineAsInt, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			contentMap[key] = append(contentMap[key], lineAsInt)
		}
	}

	return contentMap
}

func contentMapSums(contentMap map[int][]int) []int {
	sums := []int{}

	for _, v := range contentMap {
		sum_weights := 0
		for _, weight := range v {
			sum_weights += weight
		}
		sums = append(sums, sum_weights)
	}

	slices.Sort(sums)
	slices.Reverse(sums)

	return sums
}

func part1(puzzleInput []byte) {
	contentMap := fileContentsToMap(puzzleInput)
	sums := contentMapSums(contentMap)
	log.Printf("Largest weight: %d", sums[0])
}

func part2(puzzleInput []byte) {
	contentMap := fileContentsToMap(puzzleInput)
	sums := contentMapSums(contentMap)
	topThree := sums[0] + sums[1] + sums[2]

	log.Printf("Top three largest weights combined: %d", topThree)
}

func part1Alt1() {
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

	log.Printf("[Aternative solution] Largest weight: %d", maxWeight)
}

func main() {
	puzzleInput := readFile("dec1.txt")
	part1(puzzleInput)
	part2(puzzleInput)
	part1Alt1()
}
