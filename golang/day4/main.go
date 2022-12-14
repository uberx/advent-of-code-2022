package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/uberx/advent-of-code-2022/util"
)

func main() {
	start := time.Now()
	input := util.ReadFile("day4.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := totalFullyOverlappingAssignmentPairs(input)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (totalFullyOverlappingAssignmentPairs): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := totalPartiallyOverlappingAssignmentPairs(input)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (totalPartiallyOverlappingAssignmentPairs): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func totalFullyOverlappingAssignmentPairs(input string) (int, time.Duration, time.Duration) {
	start := time.Now()
	assignmentPairs, parseDuration := parseInput(input)

	fullyOverlappingAssignmentPairs := 0
	for _, assignmentPair := range assignmentPairs {
		if pairIsFullyOverlapping(assignmentPair) {
			fullyOverlappingAssignmentPairs++
		}
	}
	return fullyOverlappingAssignmentPairs, parseDuration, time.Since(start)
}

func pairIsFullyOverlapping(assignmentPair util.Pair[util.Pair[int, int], util.Pair[int, int]]) bool {
	firstPair := assignmentPair.First
	secondPair := assignmentPair.Second
	return (firstPair.First >= secondPair.First && firstPair.Second <= secondPair.Second) ||
		(secondPair.First >= firstPair.First && secondPair.Second <= firstPair.Second)
}

func totalPartiallyOverlappingAssignmentPairs(input string) (int, time.Duration, time.Duration) {
	start := time.Now()
	assignmentPairs, parseDuration := parseInput(input)

	fullyOverlappingAssignmentPairs := 0
	for _, assignmentPair := range assignmentPairs {
		if !pairIsNotOverlapping(assignmentPair) {
			fullyOverlappingAssignmentPairs++
		}
	}
	return fullyOverlappingAssignmentPairs, parseDuration, time.Since(start)
}

func pairIsNotOverlapping(assignmentPair util.Pair[util.Pair[int, int], util.Pair[int, int]]) bool {
	firstPair := assignmentPair.First
	secondPair := assignmentPair.Second
	return (firstPair.First < secondPair.First && firstPair.Second < secondPair.First) ||
		(secondPair.First < firstPair.First && secondPair.Second < firstPair.First)
}

func parseInput(input string) ([]util.Pair[util.Pair[int, int], util.Pair[int, int]], time.Duration) {
	start := time.Now()
	assignmentPairs := []util.Pair[util.Pair[int, int], util.Pair[int, int]]{}
	for _, currPair := range strings.Split(input, "\n") {
		individualPairs := strings.Split(currPair, ",")
		firstPair := strings.Split(individualPairs[0], "-")
		secondPair := strings.Split(individualPairs[1], "-")
		assignmentPairs = append(assignmentPairs, util.Pair[util.Pair[int, int], util.Pair[int, int]]{First: util.Pair[int, int]{First: util.ToInt(firstPair[0]), Second: util.ToInt(firstPair[1])}, Second: util.Pair[int, int]{First: util.ToInt(secondPair[0]), Second: util.ToInt(secondPair[1])}})
	}
	return assignmentPairs, time.Since(start)
}
