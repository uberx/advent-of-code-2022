package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/uberx/advent-of-code-2022/mathy"
	"github.com/uberx/advent-of-code-2022/util"
)

type Point struct {
	x int
	y int
}

type State struct {
	rockPaths [][]Point
	minX      int
	maxX      int
	maxY      int
}

func main() {
	start := time.Now()
	input := util.ReadFile("day14.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := sandParticlesAtRest1(input)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (sandParticlesAtRest1): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := sandParticlesAtRest2(input)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (sandParticlesAtRest2): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func sandParticlesAtRest1(input string) (int, time.Duration, time.Duration) {
	state, parseDuration := parseInput(input)

	start := time.Now()
	sandParticlesAtRest := sandParticlesAtRest(
		state,
		func(state State, sandLocation Point) bool {
			return sandLocation.x < state.minX || sandLocation.x > state.maxX || sandLocation.y > state.maxY
		},
		func(state State, p Point) bool {
			return true
		})
	return sandParticlesAtRest, parseDuration, time.Since(start)
}

func sandParticlesAtRest(state State, stopSandPour func(State, Point) bool, additionalBoundaryCheck func(State, Point) bool) int {
	particles := particles(state.rockPaths)
	sandParticlesAtRest := 0
	sandPouring := true
	for sandPouring {
		sandLocation := Point{500, 0}
		sandAtRest := false
		for !sandAtRest {
			potentialSandLocations := []Point{{sandLocation.x, sandLocation.y + 1}, {sandLocation.x - 1, sandLocation.y + 1}, {sandLocation.x + 1, sandLocation.y + 1}}
			sandParticleMoved := false
			for _, potentialSandLocation := range potentialSandLocations {
				if _, ok := particles[potentialSandLocation]; !ok && additionalBoundaryCheck(state, potentialSandLocation) {
					sandLocation = potentialSandLocation
					sandParticleMoved = true
					break
				}
			}
			_, locationNotOccupied := particles[sandLocation]
			if !sandParticleMoved && !locationNotOccupied {
				sandAtRest = true
				particles[sandLocation] = true
				sandParticlesAtRest++
			} else {
				if stopSandPour(state, sandLocation) {
					sandPouring = false
					break
				}
			}
		}
	}
	return sandParticlesAtRest
}

func particles(rockPaths [][]Point) map[Point]bool {
	particles := map[Point]bool{}
	for _, rockPath := range rockPaths {
		for i := 1; i < len(rockPath); i++ {
			curr := rockPath[i]
			prev := rockPath[i-1]
			if curr.x == prev.x {
				for y := mathy.Min(curr.y, prev.y); y <= mathy.Max(curr.y, prev.y); y++ {
					particles[Point{curr.x, y}] = true
				}
			} else {
				for x := mathy.Min(curr.x, prev.x); x <= mathy.Max(curr.x, prev.x); x++ {
					particles[Point{x, curr.y}] = true
				}
			}
		}
	}
	return particles
}

func sandParticlesAtRest2(input string) (int, time.Duration, time.Duration) {
	state, parseDuration := parseInput(input)

	start := time.Now()
	sandParticlesAtRest := sandParticlesAtRest(
		state,
		func(state State, sandLocation Point) bool {
			return sandLocation == Point{500, 0} || sandLocation.y > state.maxY+1
		},
		func(state State, potentialSandLocation Point) bool {
			return potentialSandLocation.y < state.maxY+2
		})
	return sandParticlesAtRest, parseDuration, time.Since(start)
}

func parseInput(input string) (State, time.Duration) {
	start := time.Now()
	rockPaths := [][]Point{}
	minX := math.MaxInt
	maxX := math.MinInt
	maxY := math.MinInt
	for _, line := range strings.Split(input, "\n") {
		points := strings.Split(line, " -> ")
		rockPath := []Point{}
		for _, point := range points {
			coords := strings.Split(point, ",")
			x, y := util.ToInt(coords[0]), util.ToInt(coords[1])
			rockPath = append(rockPath, Point{x, y})
			if x < minX {
				minX = x
			}
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
		rockPaths = append(rockPaths, rockPath)
	}
	return State{rockPaths, minX, maxX, maxY}, time.Since(start)
}
