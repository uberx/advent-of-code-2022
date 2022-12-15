package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/uberx/advent-of-code-2022/mathy"
	"github.com/uberx/advent-of-code-2022/util"
)

type Position struct {
	sensor        util.Point
	closestBeacon util.Point
}

type State struct {
	positions    []Position
	sensors      map[util.Point]bool
	sensorsAtRow map[int]map[util.Point]bool
	beacons      map[util.Point]bool
	beaconsAtRow map[int]map[util.Point]bool
}

func main() {
	start := time.Now()
	input := util.ReadFile("day15.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := numNonViableDistressBeaconPositionsAtRow(input, 2000000)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (numNonViableDistressBeaconPositionsAtRow): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := distressBeaconTuningFrequency(input, util.Point{X: 0, Y: 4000000})
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (distressBeaconTuningFrequency): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func numNonViableDistressBeaconPositionsAtRow(input string, rowY int) (int, time.Duration, time.Duration) {
	state, parseDuration := parseInput(input)

	start := time.Now()
	nonViableDistressBeaconPointsForRowY := nonViableDistressBeaconPointsForRowY(state, rowY)
	numNonViableDistressBeaconPointsForRowY := numNonViableDistressBeaconPointsForRowY(state, nonViableDistressBeaconPointsForRowY, rowY)
	return numNonViableDistressBeaconPointsForRowY, parseDuration, time.Since(start)
}

func numNonViableDistressBeaconPointsForRowY(state State, nonViableDistressBeaconPointsForRowY []util.Point, rowY int) int {
	numNonViableDistressBeaconPointsForRowY := 0
	for _, mergedXRangeForRowY := range nonViableDistressBeaconPointsForRowY {
		numNonViableDistressBeaconPointsForRowY += mergedXRangeForRowY.Y - mergedXRangeForRowY.X + 1
	}

	return numNonViableDistressBeaconPointsForRowY - len(state.sensorsAtRow[rowY]) - len(state.beaconsAtRow[rowY])
}

func nonViableDistressBeaconPointsForRowY(state State, rowY int) []util.Point {
	nonViableDistressBeaconPoints := []util.Point{}
	for _, position := range state.positions {
		manhattanDistance := util.ManhattanDistance(position.sensor, position.closestBeacon)
		xDiff := -1
		if rowY >= position.sensor.Y-manhattanDistance && rowY <= position.sensor.Y {
			xDiff = rowY - position.sensor.Y + manhattanDistance
		} else if rowY >= position.sensor.Y+1 && rowY <= position.sensor.Y+manhattanDistance {
			xDiff = manhattanDistance - (rowY - position.sensor.Y)
		}
		if xDiff != -1 {
			nonViableDistressBeaconPoints = append(nonViableDistressBeaconPoints, util.Point{X: position.sensor.X - xDiff, Y: position.sensor.X + xDiff})
		}
	}
	return mergedXRanges(nonViableDistressBeaconPoints)
}

func mergedXRanges(xRanges []util.Point) []util.Point {
	mergedXRanges := xRanges
	fullyMerged := false
	for !fullyMerged {
		if len(mergedXRanges) == 1 {
			return mergedXRanges
		}
		newMergedXRanges := []util.Point{}
		for i := 0; i < len(mergedXRanges)-1; i++ {
			for fill := 0; fill < i; fill++ {
				newMergedXRanges = append(newMergedXRanges, mergedXRanges[fill])
			}
			jProgress := len(mergedXRanges)
			mergeOccurred := false
			for j := i + 1; j < len(mergedXRanges) && !mergeOccurred; j++ {
				mergable := (mergedXRanges[i].X >= mergedXRanges[j].X && mergedXRanges[i].X <= mergedXRanges[j].Y+1) ||
					(mergedXRanges[i].Y >= mergedXRanges[j].X-1 && mergedXRanges[i].Y <= mergedXRanges[j].Y) ||
					(mergedXRanges[i].X < mergedXRanges[j].X && mergedXRanges[i].Y > mergedXRanges[j].Y)
				if mergable {
					newMergedXRanges = append(newMergedXRanges, util.Point{X: mathy.Min(mergedXRanges[i].X, mergedXRanges[j].X), Y: mathy.Max(mergedXRanges[i].Y, mergedXRanges[j].Y)})
					mergeOccurred = true
					jProgress = j
					break
				} else {
					newMergedXRanges = append(newMergedXRanges, mergedXRanges[j])
				}
			}
			if mergeOccurred {
				for j := jProgress + 1; j < len(mergedXRanges); j++ {
					newMergedXRanges = append(newMergedXRanges, mergedXRanges[j])
				}
				mergedXRanges = newMergedXRanges
				break
			} else {
				if i == len(mergedXRanges)-2 {
					fullyMerged = true
				} else {
					newMergedXRanges = append(newMergedXRanges, mergedXRanges[i])
				}
			}
		}
	}
	return mergedXRanges
}

func distressBeaconTuningFrequency(input string, searchBoundary util.Point) (int, time.Duration, time.Duration) {
	state, parseDuration := parseInput(input)

	start := time.Now()
	for y := searchBoundary.Y; y >= searchBoundary.X; y-- {
		nonViableDistressBeaconPointsForRowY := nonViableDistressBeaconPointsForRowY(state, y)
		if len(nonViableDistressBeaconPointsForRowY) == 2 {
			var x int
			if nonViableDistressBeaconPointsForRowY[0].X < nonViableDistressBeaconPointsForRowY[1].X {
				x = nonViableDistressBeaconPointsForRowY[1].X - 1
			} else {
				x = nonViableDistressBeaconPointsForRowY[0].X - 1
			}
			return 4000000*x + y, parseDuration, time.Since(start)
		}
	}
	panic("could not find distress beacon")
}

func parseInput(input string) (State, time.Duration) {
	start := time.Now()
	positions := []Position{}
	sensors := map[util.Point]bool{}
	sensorsAtRow := map[int]map[util.Point]bool{}
	beacons := map[util.Point]bool{}
	beaconsAtRow := map[int]map[util.Point]bool{}
	r := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	for _, line := range strings.Split(input, "\n") {
		matches := r.FindStringSubmatch(line)
		sensor := util.Point{X: util.ToInt(matches[1]), Y: util.ToInt(matches[2])}
		beacon := util.Point{X: util.ToInt(matches[3]), Y: util.ToInt(matches[4])}
		positions = append(positions, Position{sensor, beacon})
		sensors[sensor] = true
		if _, ok := sensorsAtRow[sensor.Y]; !ok {
			sensorsAtRow[sensor.Y] = map[util.Point]bool{}
		}
		sensorsAtRow[sensor.Y][sensor] = true
		beacons[beacon] = true
		if _, ok := beaconsAtRow[sensor.Y]; !ok {
			beaconsAtRow[beacon.Y] = map[util.Point]bool{}
		}
		beaconsAtRow[beacon.Y][beacon] = true
	}
	return State{positions, sensors, sensorsAtRow, beacons, beaconsAtRow}, time.Since(start)
}
