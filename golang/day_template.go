package main

import (
	"fmt"
	"time"

	"github.com/uberx/advent-of-code-2022/util"
)

func main() {
	start := time.Now()
	input := util.ReadFile("day_template.go")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := part1(input)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (part1): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := part2(input)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (part2): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func part1(input string) (int, time.Duration, time.Duration) {
	_, parseDuration := parseInput(input)

	start := time.Now()
	return -1, parseDuration, time.Since(start)
}

func part2(input string) (int, time.Duration, time.Duration) {
	_, parseDuration := parseInput(input)

	start := time.Now()
	return -1, parseDuration, time.Since(start)
}

func parseInput(input string) (interface{}, time.Duration) {
	start := time.Now()
	return "", time.Since(start)
}
