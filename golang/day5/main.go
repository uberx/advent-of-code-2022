package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/uberx/advent-of-code-2022/util"
)

type stack struct {
	crates []rune
}

type move struct {
	numCrates int
	fromStack int
	toStack   int
}

func main() {
	start := time.Now()
	input := util.ReadFile("day5.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := topOfEachStack(input, false)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (topOfEachStack1): %s (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := topOfEachStack(input, true)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (topOfEachStack2): %s (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func topOfEachStack(input string, part2 bool) (string, time.Duration, time.Duration) {
	stacks, moves, parseDuration := parseInput(input)

	start := time.Now()
	for _, move := range moves {
		numCrates := move.numCrates
		fromStack := move.fromStack
		toStack := move.toStack

		cratesToMove := stacks[fromStack].crates[len(stacks[fromStack].crates)-numCrates:]
		stacks[fromStack].crates = stacks[fromStack].crates[:len(stacks[fromStack].crates)-numCrates]
		if !part2 {
			util.Reverse(cratesToMove)
		}
		stacks[toStack].crates = append(stacks[toStack].crates, cratesToMove...)
	}

	topOfEachStack := ""
	for _, stack := range stacks {
		topOfEachStack += string(stack.crates[len(stack.crates)-1])
	}

	return topOfEachStack, parseDuration, time.Since(start)
}

func parseInput(input string) ([]stack, []move, time.Duration) {
	start := time.Now()

	inputSections := strings.Split(input, "\n\n")
	stacksSection := inputSections[0]
	movesSection := inputSections[1]

	stackSections := strings.Split(stacksSection, "\n")
	numElements := (len(stackSections[0]) + 1) / 4
	stacks := []stack{}
	for i := 0; i < numElements; i++ {
		stacks = append(stacks, stack{[]rune{}})
	}
	for _, stackSection := range stackSections {
		stackSection += " "
		for i := 0; i < numElements; i++ {
			crate := strings.TrimSpace(stackSection[4*i : 4*i+4])
			if len(crate) == 3 {
				stacks[i].crates = append([]rune{rune(crate[1])}, stacks[i].crates...)
			}
		}
	}

	moves := []move{}
	r := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	for _, moveSection := range strings.Split(movesSection, "\n") {
		matches := r.FindStringSubmatch(moveSection)
		moves = append(moves, move{util.ToInt(matches[1]), util.ToInt(matches[2]) - 1, util.ToInt(matches[3]) - 1})
	}

	return stacks, moves, time.Since(start)
}
