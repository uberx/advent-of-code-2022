package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/uberx/advent-of-code-2022/mathy"
	"github.com/uberx/advent-of-code-2022/util"
)

type State struct {
	board       [][]string
	path        []util.Pair[string, int]
	xBoundaries map[int]util.Point
	yBoundaries map[int]util.Point
}

func main() {
	start := time.Now()
	input := util.ReadFile("day22.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := part1(input)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (part1): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := part2(input)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (part2): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func part1(input string) (int, time.Duration, time.Duration) {
	state, parseDuration := parseInput(input)

	start := time.Now()
	rotations := map[string]map[string]string{
		">": {
			"L": "^",
			"R": "v",
		},
		"v": {
			"L": ">",
			"R": "<",
		},
		"<": {
			"L": "v",
			"R": "^",
		},
		"^": {
			"L": "<",
			"R": ">",
		},
	}
	currPosition := util.Point{X: 0, Y: state.yBoundaries[0].X}
	var newPosition util.Point
	currentDirection := "^"
	for _, move := range state.path {
		currentDirection = rotations[currentDirection][move.First]
		numTiles := move.Second
		for i := 0; i < numTiles; i++ {
			switch currentDirection {
			case ">":
				newPosition = util.Point{X: currPosition.X, Y: currPosition.Y + 1}
			case "v":
				newPosition = util.Point{X: currPosition.X + 1, Y: currPosition.Y}
			case "<":
				newPosition = util.Point{X: currPosition.X, Y: currPosition.Y - 1}
			case "^":
				newPosition = util.Point{X: currPosition.X - 1, Y: currPosition.Y}
			}
			// wrap around
			if currentDirection == "v" || currentDirection == "^" {
				if newPosition.X < state.xBoundaries[newPosition.Y].X {
					newPosition.X = state.xBoundaries[newPosition.Y].Y
				}
				if newPosition.X > state.xBoundaries[newPosition.Y].Y {
					newPosition.X = state.xBoundaries[newPosition.Y].X
				}
			} else if currentDirection == ">" || currentDirection == "<" {
				if newPosition.Y < state.yBoundaries[newPosition.X].X {
					newPosition.Y = state.yBoundaries[newPosition.X].Y
				}
				if newPosition.Y > state.yBoundaries[newPosition.X].Y {
					newPosition.Y = state.yBoundaries[newPosition.X].X
				}
			}
			// wall check
			if state.board[newPosition.X][newPosition.Y] == "#" {
				break
			} else {
				currPosition = newPosition
			}
		}
	}
	var facing int
	switch currentDirection {
	case ">":
		facing = 0
	case "v":
		facing = 1
	case "<":
		facing = 2
	case "^":
		facing = 3
	}
	return 1000*(currPosition.X+1) + 4*(currPosition.Y+1) + facing, parseDuration, time.Since(start)
}

func part2(input string) (int, time.Duration, time.Duration) {
	_, parseDuration := parseInput(input)

	start := time.Now()
	return -1, parseDuration, time.Since(start)
}

func parseInput(input string) (State, time.Duration) {
	start := time.Now()
	inputTokens := strings.Split(input, "\n\n")
	boardLines := inputTokens[0]
	board := [][]string{}
	xBoundaries := map[int]util.Point{}
	yBoundaries := map[int]util.Point{}
	var cols int
	for _, boardLine := range strings.Split(inputTokens[0], "\n") {
		cols = mathy.Max(cols, len(boardLine))
	}
	for rowNum, boardLine := range strings.Split(boardLines, "\n") {
		for i := len(boardLine); i < cols; i++ {
			boardLine += " "
		}
		row := []string{}
		yBoundary := util.Point{X: -1, Y: -1}
		for colNum, cell := range boardLine {
			row = append(row, string(cell))
			if _, ok := xBoundaries[colNum]; !ok {
				xBoundaries[colNum] = util.Point{X: -1, Y: -1}
			}
			if cell == '.' || cell == '#' {
				xBoundary := xBoundaries[colNum]
				if xBoundary.X == -1 {
					xBoundary.X = rowNum
				}
				xBoundary.Y = rowNum
				xBoundaries[colNum] = xBoundary

				if yBoundary.X == -1 {
					yBoundary.X = colNum
				}
				yBoundary.Y = colNum
			}
		}
		yBoundaries[rowNum] = yBoundary
		board = append(board, row)
	}
	pathLine := inputTokens[1]
	path := []util.Pair[string, int]{}
	numTiles := 0
	direction := "R"
	for _, pathRune := range pathLine {
		if pathRune == 'L' || pathRune == 'R' {
			path = append(path, util.Pair[string, int]{First: direction, Second: numTiles})
			numTiles = 0
			direction = string(pathRune)
		} else {
			numTiles = numTiles*10 + (int(pathRune) - '0')
		}
	}
	path = append(path, util.Pair[string, int]{First: direction, Second: numTiles})
	return State{board, path, xBoundaries, yBoundaries}, time.Since(start)
}
