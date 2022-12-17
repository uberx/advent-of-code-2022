package main

import (
	"fmt"
	"time"

	"github.com/uberx/advent-of-code-2022/util"
	"golang.org/x/exp/slices"
)

func main() {
	start := time.Now()
	input := util.ReadFile("day17.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := rockTowerLength(input, 2022)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (rockTowerLength1): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := rockTowerLength(input, 1000000000000)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (rockTowerLength2): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

type State struct {
	jetPatterns []rune
	rockTypes   []rune
}

func rockTowerLength(input string, numRocks int) (int, time.Duration, time.Duration) {
	state, parseDuration := parseInput(input)

	start := time.Now()
	chamber := map[util.Point]bool{}
	rockPatterns := map[int]map[int]bool{}
	jetPatternIdx := 0
	highestRockY := -1
	highestRockYDiffs := []int{}
	patternFound := false
	rockAndJetPatternLength := -1
	rockNum := 0
	for !patternFound {
		if rockAndJetPatternLength == -1 {
			if _, ok := rockPatterns[rockNum%len(state.rockTypes)]; !ok {
				rockPatterns[rockNum%len(state.rockTypes)] = map[int]bool{}
			}
			if _, ok := rockPatterns[rockNum%len(state.rockTypes)][jetPatternIdx%len(state.jetPatterns)]; ok {
				rockAndJetPatternLength = rockNum * 4
			} else {
				rockPatterns[rockNum%len(state.rockTypes)][jetPatternIdx%len(state.jetPatterns)] = true
			}
		}
		currRock := spawnRock(state.rockTypes, rockNum, highestRockY)
		rockAtRest := false
		for !rockAtRest {
			newRock, _ := moveRock(currRock, state.jetPatterns[jetPatternIdx%len(state.jetPatterns)], chamber)
			jetPatternIdx++
			newRock, rockAtRest = moveRock(newRock, 'v', chamber)
			currRock = newRock
		}
		newHighestRockY := updateChamber(chamber, currRock, highestRockY)
		highestRockYDiffs = append(highestRockYDiffs, newHighestRockY-highestRockY)
		highestRockY = newHighestRockY
		rockNum++
		if rockNum == rockAndJetPatternLength {
			rockAndJetPatternLength /= 2
			patternFound = true
		}
	}
	repeatingPattern := findRepeatingPattern(highestRockYDiffs)
	extraLength := rockAndJetPatternLength - len(repeatingPattern)
	cycles := (numRocks - extraLength) / len(repeatingPattern)
	lastCycleLength := (numRocks - extraLength) % len(repeatingPattern)
	return util.Sum(highestRockYDiffs[0:extraLength]) + (cycles * util.Sum(repeatingPattern)) + util.Sum(repeatingPattern[0:lastCycleLength]), parseDuration, time.Since(start)
}

func spawnRock(rockTypes []rune, rockNum int, highestRockY int) map[util.Point]bool {
	rockToSpawn := rockTypes[rockNum%len(rockTypes)]
	switch rockToSpawn {
	case '-':
		return map[util.Point]bool{
			{X: 2, Y: highestRockY + 4}: true,
			{X: 3, Y: highestRockY + 4}: true,
			{X: 4, Y: highestRockY + 4}: true,
			{X: 5, Y: highestRockY + 4}: true,
		}
	case '+':
		return map[util.Point]bool{
			{X: 2, Y: highestRockY + 5}: true,
			{X: 3, Y: highestRockY + 4}: true,
			{X: 3, Y: highestRockY + 5}: true,
			{X: 3, Y: highestRockY + 6}: true,
			{X: 4, Y: highestRockY + 5}: true,
		}
	case 'L':
		return map[util.Point]bool{
			{X: 2, Y: highestRockY + 4}: true,
			{X: 3, Y: highestRockY + 4}: true,
			{X: 4, Y: highestRockY + 4}: true,
			{X: 4, Y: highestRockY + 5}: true,
			{X: 4, Y: highestRockY + 6}: true,
		}
	case '|':
		return map[util.Point]bool{
			{X: 2, Y: highestRockY + 4}: true,
			{X: 2, Y: highestRockY + 5}: true,
			{X: 2, Y: highestRockY + 6}: true,
			{X: 2, Y: highestRockY + 7}: true,
		}
	case 'O':
		return map[util.Point]bool{
			{X: 2, Y: highestRockY + 4}: true,
			{X: 2, Y: highestRockY + 5}: true,
			{X: 3, Y: highestRockY + 4}: true,
			{X: 3, Y: highestRockY + 5}: true,
		}
	}
	return map[util.Point]bool{}
}

func moveRock(rock map[util.Point]bool, jetPattern rune, chamber map[util.Point]bool) (map[util.Point]bool, bool) {
	newRock := map[util.Point]bool{}
	for point := range rock {
		var newPoint util.Point
		if jetPattern == '<' {
			if point.X == 0 {
				return rock, true
			}
			newPoint = util.Point{X: point.X - 1, Y: point.Y}
		} else if jetPattern == '>' {
			if point.X == 6 {
				return rock, true
			}
			newPoint = util.Point{X: point.X + 1, Y: point.Y}
		} else if jetPattern == 'v' {
			if point.Y == 0 {
				return rock, true
			}
			newPoint = util.Point{X: point.X, Y: point.Y - 1}
		}
		if _, ok := chamber[newPoint]; ok {
			return rock, true
		}
		newRock[newPoint] = true
	}
	return newRock, false
}

func updateChamber(chamber map[util.Point]bool, rock map[util.Point]bool, highestRockY int) int {
	newHighestRockY := highestRockY
	for point := range rock {
		chamber[point] = true
		if point.Y > newHighestRockY {
			newHighestRockY = point.Y
		}
	}
	return newHighestRockY
}

func findRepeatingPattern(highestRockYDiffs []int) []int {
	for i := 0; i < len(highestRockYDiffs)/2; i++ {
		if slices.Compare(highestRockYDiffs[i:len(highestRockYDiffs)/2], highestRockYDiffs[len(highestRockYDiffs)/2:len(highestRockYDiffs)-i]) == 0 {
			return highestRockYDiffs[i : len(highestRockYDiffs)/2]
		}
	}
	panic("no repeating pattern found")
}

func parseInput(input string) (State, time.Duration) {
	start := time.Now()
	jetPatterns := []rune{}
	for _, jetPattern := range input {
		jetPatterns = append(jetPatterns, jetPattern)
	}
	return State{jetPatterns, []rune{'-', '+', 'L', '|', 'O'}}, time.Since(start)
}
