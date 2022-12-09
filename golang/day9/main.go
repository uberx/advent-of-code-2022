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

func (p *Point) move(direction rune) {
	switch direction {
	case 'U':
		p.y += 1
	case 'D':
		p.y += -1
	case 'L':
		p.x += -1
	case 'R':
		p.x += 1
	}
}

type Motion struct {
	direction rune
	steps     int
}

type State struct {
	knots   []*Point
	motions []Motion
}

func main() {
	start := time.Now()
	input := util.ReadFile("day9.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := numPositionsVisitedByTail(input, 2)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (numPositionsVisitedByTail1): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := numPositionsVisitedByTail(input, 10)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (numPositionsVisitedByTail2): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func numPositionsVisitedByTail(input string, numKnots int) (int, time.Duration, time.Duration) {
	state, parseDuration := parseInput(input, numKnots)

	start := time.Now()
	tailVisits := map[Point]bool{}
	tailVisits[*state.knots[len(state.knots)-1]] = true

	motions := state.motions
	for _, motion := range motions {
		doMotion(motion, state.knots, tailVisits)
	}
	return len(tailVisits), parseDuration, time.Since(start)
}

func isAdjacent(a Point, b Point) bool {
	return a.x-b.x >= -1 && a.x-b.x <= 1 && a.y-b.y >= -1 && a.y-b.y <= 1
}

func doMotion(motion Motion, knots []*Point, tailVisits map[Point]bool) {
	direction := motion.direction
	for step := 1; step <= motion.steps; step++ {
		currKnot := 0
		knots[currKnot].move(direction)
		for currKnot < len(knots)-1 && !isAdjacent(*knots[currKnot], *knots[currKnot+1]) {
			moveTowards(knots[currKnot+1], knots[currKnot])
			tailVisits[*knots[len(knots)-1]] = true
			currKnot++
		}
	}
}

func moveTowards(a *Point, b *Point) {
	top := b.y > a.y
	bottom := b.y < a.y
	left := b.x < a.x
	right := b.x > a.x

	if top {
		a.y += 1
	} else if bottom {
		a.y += -1
	}
	if left {
		a.x += -1
	} else if right {
		a.x += 1
	}
}

func parseInput(input string, numKnots int) (State, time.Duration) {
	start := time.Now()
	motions := []Motion{}
	for _, motionLine := range strings.Split(input, "\n") {
		motionParams := strings.Split(motionLine, " ")
		motions = append(motions, Motion{rune(motionParams[0][0]), util.ToInt(motionParams[1])})
	}
	knots := []*Point{}
	for i := 0; i < numKnots; i++ {
		knots = append(knots, &Point{})
	}
	return State{knots, motions}, time.Since(start)
}
