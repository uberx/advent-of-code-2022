package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/uberx/advent-of-code-2022/util"
)

type Point struct {
	x int
	y int
}

type Direction string

const (
	Up    Direction = "^"
	Down  Direction = "v"
	Left  Direction = "<"
	Right Direction = ">"
)

type State struct {
	heightmap     [][]rune
	startPosition Point
	endPosition   Point
}

type StackItem struct {
	currentLocation Point
	previousItem    *StackItem
	directionTaken  *Direction
	pathLength      int
}

func main() {
	start := time.Now()
	input := util.ReadFile("day12.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := fewestStepsFromStartToEnd(input)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (fewestStepsFromStartToEnd): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := fewestStepsFromEndToAnyLowestElevationPoint(input)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (fewestStepsFromEndToAnyLowestElevationPoint): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func fewestStepsFromStartToEnd(input string) (int, time.Duration, time.Duration) {
	state, parseDuration := parseInput(input)

	start := time.Now()
	finishedPath := finishedPath(
		state,
		func(elevationDiff rune) bool {
			return elevationDiff <= 1
		},
		func(state State, currItem StackItem) bool {
			return currItem.currentLocation == state.endPosition
		})
	return finishedPath.pathLength, parseDuration, time.Since(start)
}

func finishedPath(state State, elevationDiffFunc func(diff rune) bool, endConditionFunc func(State, StackItem) bool) StackItem {
	queue := util.Queue[StackItem]{}
	queue.Queue(StackItem{state.startPosition, nil, nil, 0})
	visited := map[Point]int{}
	for !queue.IsEmpty() {
		currItem, _ := queue.Dequeue()
		if endConditionFunc(state, currItem) {
			return currItem
		}
		neighbors := validNeighbors(state.heightmap, currItem.currentLocation, elevationDiffFunc)
		for direction, location := range neighbors {
			newPathLength := currItem.pathLength + 1
			if existingPathLength, ok := visited[location]; !ok || newPathLength < existingPathLength {
				visited[location] = newPathLength
				queue.Queue(StackItem{location, &currItem, &direction, newPathLength})
			}
		}
	}
	panic("no path found")
}

func validNeighbors(heightmap [][]rune, currentLocation Point, elevationDiffFunc func(diff rune) bool) map[Direction]Point {
	currentHeight := heightmap[currentLocation.x][currentLocation.y]

	neighbors := map[Direction]Point{}
	// Up
	if currentLocation.x-1 >= 0 {
		destinationHeight := heightmap[currentLocation.x-1][currentLocation.y]
		elevationDiff := destinationHeight - currentHeight
		if elevationDiffFunc(elevationDiff) {
			neighbors[Up] = Point{currentLocation.x - 1, currentLocation.y}
		}
	}
	// Down
	if currentLocation.x+1 <= len(heightmap)-1 {
		destinationHeight := heightmap[currentLocation.x+1][currentLocation.y]
		elevationDiff := destinationHeight - currentHeight
		if elevationDiffFunc(elevationDiff) {
			neighbors[Down] = Point{currentLocation.x + 1, currentLocation.y}
		}
	}
	// Left
	if currentLocation.y-1 >= 0 {
		destinationHeight := heightmap[currentLocation.x][currentLocation.y-1]
		elevationDiff := destinationHeight - currentHeight
		if elevationDiffFunc(elevationDiff) {
			neighbors[Left] = Point{currentLocation.x, currentLocation.y - 1}
		}
	}
	// Right
	if currentLocation.y+1 <= len(heightmap[0])-1 {
		destinationHeight := heightmap[currentLocation.x][currentLocation.y+1]
		elevationDiff := destinationHeight - currentHeight
		if elevationDiffFunc(elevationDiff) {
			neighbors[Right] = Point{currentLocation.x, currentLocation.y + 1}
		}
	}
	return neighbors
}

func fewestStepsFromEndToAnyLowestElevationPoint(input string) (int, time.Duration, time.Duration) {
	state, parseDuration := parseInput(input)

	start := time.Now()
	finishedPath := finishedPath(
		State{state.heightmap, state.endPosition, Point{-1, -1}},
		func(elevationDiff rune) bool {
			return elevationDiff >= -1
		},
		func(state State, currItem StackItem) bool {
			return state.heightmap[currItem.currentLocation.x][currItem.currentLocation.y] == 'a'
		})
	return finishedPath.pathLength, parseDuration, time.Since(start)
}

func parseInput(input string) (State, time.Duration) {
	start := time.Now()
	heightmap := [][]rune{}
	var startPosition Point
	var endPosition Point
	for row, line := range strings.Split(input, "\n") {
		heightRow := []rune{}
		for col, height := range line {
			if height == 'S' {
				startPosition = Point{row, col}
				height = 'a'
			} else if height == 'E' {
				endPosition = Point{row, col}
				height = 'z'
			}
			heightRow = append(heightRow, height)
		}
		heightmap = append(heightmap, heightRow)
	}
	return State{heightmap, startPosition, endPosition}, time.Since(start)
}
