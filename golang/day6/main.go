package main

import (
	"fmt"
	"time"

	"github.com/uberx/advent-of-code-2022/util"
)

func main() {
	start := time.Now()
	input := util.ReadFile("day6.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := startOfPacketMarker(input)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (startOfPacketMarker): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := startOfMessageMarker(input)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (startOfMessageMarker): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func startOfPacketMarker(input string) (int, time.Duration, time.Duration) {
	return startOfMarker(input, 4)
}

func startOfMessageMarker(input string) (int, time.Duration, time.Duration) {
	return startOfMarker(input, 14)
}

func startOfMarker(input string, length int) (int, time.Duration, time.Duration) {
	characters, parseDuration := parseInput(input)

	start := time.Now()
	parsedCharacters := map[rune]int{}
	leftPointer := -1
	for i, character := range characters {
		if currIndex, ok := parsedCharacters[character]; ok {
			if currIndex > leftPointer {
				leftPointer = currIndex
			}
		}
		if i == leftPointer+length {
			return i + 1, parseDuration, time.Since(start)
		}
		parsedCharacters[character] = i
	}

	return -1, parseDuration, time.Since(start)
}

func parseInput(input string) ([]rune, time.Duration) {
	start := time.Now()
	characters := []rune{}
	for _, character := range input {
		characters = append(characters, character)
	}
	return characters, time.Since(start)
}
