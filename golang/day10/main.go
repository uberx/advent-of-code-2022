package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/uberx/advent-of-code-2022/util"
)

type Operation string

type Instruction struct {
	operation Operation
	operand   int
}

func main() {
	start := time.Now()
	input := util.ReadFile("day10.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := sumOfSixSignalStrengths(input)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (sumOfSixSignalStrengths): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := crtOutput(input)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (crtOutput):%s(%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func sumOfSixSignalStrengths(input string) (int, time.Duration, time.Duration) {
	instructions, parseDuration := parseInput(input)

	start := time.Now()
	registerValueByCycle := registerValueByCycle(instructions)

	wantedCycles := []int{20, 60, 100, 140, 180, 220}
	sumOfWantedSignalStrengths := 0
	for _, wantedCycle := range wantedCycles {
		sumOfWantedSignalStrengths += wantedCycle * registerValueByCycle[wantedCycle]
	}

	return sumOfWantedSignalStrengths, parseDuration, time.Since(start)
}

func registerValueByCycle(instructions []Instruction) map[int]int {
	registerValueByCycle := map[int]int{}
	x := 1
	cycleNum := 1
	for _, instruction := range instructions {
		if instruction.operation == "addx" {
			registerValueByCycle[cycleNum] = x
			registerValueByCycle[cycleNum+1] = x
			x += instruction.operand
			cycleNum += 2
		} else if instruction.operation == "noop" {
			registerValueByCycle[cycleNum] = x
			cycleNum += 1
		}
	}
	return registerValueByCycle
}

func crtOutput(input string) (string, time.Duration, time.Duration) {
	instructions, parseDuration := parseInput(input)

	start := time.Now()
	registerValueByCycle := registerValueByCycle(instructions)
	image := "\n"
	for cycle := 1; cycle <= 240; cycle++ {
		registerValue := registerValueByCycle[cycle]
		if (cycle-1)%40 >= registerValue-1 && (cycle-1)%40 <= registerValue+1 {
			image += "#"
		} else {
			image += "."
		}
		if cycle%40 == 0 {
			image += "\n"
		}
	}
	return image + "\n", parseDuration, time.Since(start)
}

func parseInput(input string) ([]Instruction, time.Duration) {
	start := time.Now()
	instructions := []Instruction{}
	for _, instructionLine := range strings.Split(input, "\n") {
		instructionTokens := strings.Split(instructionLine, " ")
		operand := -1
		if len(instructionTokens) == 2 {
			operand = util.ToInt(instructionTokens[1])
		}
		instructions = append(instructions, Instruction{Operation(instructionTokens[0]), operand})
	}
	return instructions, time.Since(start)
}
