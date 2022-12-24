package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/uberx/advent-of-code-2022/mathy"
	"github.com/uberx/advent-of-code-2022/util"
)

func main() {
	start := time.Now()
	input := util.ReadFile("day23.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := emptyGroundTilesAfter10RoundsOfElfConway(input)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (emptyGroundTilesAfter10RoundsOfElfConway): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := firstRoundWhereNoElfMoves(input)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (firstRoundWhereNoElfMoves): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func emptyGroundTilesAfter10RoundsOfElfConway(input string) (int, time.Duration, time.Duration) {
	elves, parseDuration := parseInput(input)

	start := time.Now()
	movedElves, _ := simulateElfMovement(elves, 10)
	topLeft := util.Point{X: math.MaxInt, Y: math.MaxInt}
	bottomRight := util.Point{X: math.MinInt, Y: math.MinInt}
	for elf := range movedElves {
		topLeft.X = mathy.Min(topLeft.X, elf.X)
		topLeft.Y = mathy.Min(topLeft.Y, elf.Y)
		bottomRight.X = mathy.Max(bottomRight.X, elf.X)
		bottomRight.Y = mathy.Max(bottomRight.Y, elf.Y)
	}
	return (bottomRight.X-topLeft.X+1)*(bottomRight.Y-topLeft.Y+1) - len(elves), parseDuration, time.Since(start)
}

func simulateElfMovement(elves map[util.Point]bool, rounds int) (map[util.Point]bool, int) {
	validDirectionFuncs := []func(util.Point, map[util.Point]bool) (util.Point, bool){
		func(elf util.Point, elves map[util.Point]bool) (util.Point, bool) { // north
			if !elves[util.Point{X: elf.X - 1, Y: elf.Y - 1}] && !elves[util.Point{X: elf.X - 1, Y: elf.Y}] && !elves[util.Point{X: elf.X - 1, Y: elf.Y + 1}] {
				return util.Point{X: elf.X - 1, Y: elf.Y}, true
			}
			return elf, false
		},
		func(elf util.Point, elves map[util.Point]bool) (util.Point, bool) { // south
			if !elves[util.Point{X: elf.X + 1, Y: elf.Y - 1}] && !elves[util.Point{X: elf.X + 1, Y: elf.Y}] && !elves[util.Point{X: elf.X + 1, Y: elf.Y + 1}] {
				return util.Point{X: elf.X + 1, Y: elf.Y}, true
			}
			return elf, false
		},
		func(elf util.Point, elves map[util.Point]bool) (util.Point, bool) { // west
			if !elves[util.Point{X: elf.X - 1, Y: elf.Y - 1}] && !elves[util.Point{X: elf.X, Y: elf.Y - 1}] && !elves[util.Point{X: elf.X + 1, Y: elf.Y - 1}] {
				return util.Point{X: elf.X, Y: elf.Y - 1}, true
			}
			return elf, false
		},
		func(elf util.Point, elves map[util.Point]bool) (util.Point, bool) { // east
			if !elves[util.Point{X: elf.X - 1, Y: elf.Y + 1}] && !elves[util.Point{X: elf.X, Y: elf.Y + 1}] && !elves[util.Point{X: elf.X + 1, Y: elf.Y + 1}] {
				return util.Point{X: elf.X, Y: elf.Y + 1}, true
			}
			return elf, false
		},
	}
	for round := 1; round <= rounds; round++ {
		elfMoveProposals := map[util.Point]util.Point{}
		moveProposals := map[util.Point]int{}
		noElfHasAnyNeighbors := true
		// move proposals
		for elf := range elves {
			elfMoveProposal := elf
			hasNeighbor := false
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if i == 0 && j == 0 {
						continue
					}
					if elves[util.Point{X: elf.X - i, Y: elf.Y - j}] {
						hasNeighbor = true
						noElfHasAnyNeighbors = false
						break
					}
				}
			}
			if hasNeighbor {
				for _, validDirectionFunc := range validDirectionFuncs {
					if point, ok := validDirectionFunc(elf, elves); ok {
						elfMoveProposal = point
						break
					}
				}
			}
			elfMoveProposals[elf] = elfMoveProposal
			if count, ok := moveProposals[elfMoveProposal]; ok {
				moveProposals[elfMoveProposal] = count + 1
			} else {
				moveProposals[elfMoveProposal] = 1
			}
		}
		// moves
		newElves := map[util.Point]bool{}
		for elf, elfMoveProposal := range elfMoveProposals {
			if moveProposals[elfMoveProposal] > 1 {
				newElves[elf] = true
			} else {
				newElves[elfMoveProposal] = true
			}
		}
		elves = newElves
		// rotate valid directions list
		firstFunc := validDirectionFuncs[0]
		validDirectionFuncs = append(validDirectionFuncs[1:], firstFunc)
		if noElfHasAnyNeighbors {
			return elves, round
		}
	}
	return elves, -1
}

func firstRoundWhereNoElfMoves(input string) (int, time.Duration, time.Duration) {
	elves, parseDuration := parseInput(input)

	start := time.Now()
	_, round := simulateElfMovement(elves, 1000)
	return round, parseDuration, time.Since(start)
}

func parseInput(input string) (map[util.Point]bool, time.Duration) {
	start := time.Now()
	elves := map[util.Point]bool{}
	for row, line := range strings.Split(input, "\n") {
		for col, cell := range line {
			if cell == '#' {
				elves[util.Point{X: row, Y: col}] = true
			}
		}
	}
	return elves, time.Since(start)
}
