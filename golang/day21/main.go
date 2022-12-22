package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/uberx/advent-of-code-2022/util"
)

type Monkey struct {
	number    int
	operand1  string
	operand2  string
	operation string
}

func main() {
	start := time.Now()
	input := util.ReadFile("day21.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := rootMonkeyNumber(input)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (rootMonkeyNumber): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := rootMonkeyEqualityTestNumber(input)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (rootMonkeyEqualityTestNumber): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func rootMonkeyNumber(input string) (int, time.Duration, time.Duration) {
	monkeys, parseDuration := parseInput(input)

	start := time.Now()
	rootMonkeyNumber := monkeyNumber(monkeys, "root")
	return rootMonkeyNumber, parseDuration, time.Since(start)
}

func monkeyNumber(monkeys map[string]*Monkey, name string) int {
	if monkeys[name].number != -1 {
		return monkeys[name].number
	}
	number := 0
	switch monkeys[name].operation {
	case "+":
		number = monkeyNumber(monkeys, monkeys[name].operand1) + monkeyNumber(monkeys, monkeys[name].operand2)
	case "-":
		number = monkeyNumber(monkeys, monkeys[name].operand1) - monkeyNumber(monkeys, monkeys[name].operand2)
	case "*":
		number = monkeyNumber(monkeys, monkeys[name].operand1) * monkeyNumber(monkeys, monkeys[name].operand2)
	case "/":
		number = monkeyNumber(monkeys, monkeys[name].operand1) / monkeyNumber(monkeys, monkeys[name].operand2)
	}
	monkeys[name].number = number
	return number
}

func rootMonkeyEqualityTestNumber(input string) (int, time.Duration, time.Duration) {
	monkeys, parseDuration := parseInput(input)

	start := time.Now()
	monkeyName1 := monkeys["root"].operand1
	monkeyName2 := monkeys["root"].operand2
	humnStart := 0
	for humn := 1; humn < 100000000000000; humn *= 10 {
		monkeys["humn"].number = humn
		monkeyNumber(monkeys, "root")
		if monkeys[monkeyName1].number-monkeys[monkeyName2].number < 0 {
			humnStart = humn / 10
			break
		}
		resetMonkeys(monkeys)
	}
	humnValue := humnStart
	humnIncrement := humnStart
	for monkeys[monkeyName1].number != monkeys[monkeyName2].number {
		resetMonkeys(monkeys)
		humnValue += humnIncrement
		monkeys["humn"].number = humnValue
		monkeyNumber(monkeys, "root")
		if monkeys[monkeyName1].number-monkeys[monkeyName2].number < 0 {
			humnValue -= humnIncrement
			humnIncrement /= 10
		}
	}
	return humnValue, parseDuration, time.Since(start)
}

func resetMonkeys(monkeys map[string]*Monkey) {
	for _, monkey := range monkeys {
		if monkey.operation != "" {
			monkey.number = -1
		}
	}
}

func parseInput(input string) (map[string]*Monkey, time.Duration) {
	start := time.Now()
	monkeys := map[string]*Monkey{}
	mainRegex := regexp.MustCompile(`([a-z]+): (.*)`)
	expressionRegex := regexp.MustCompile(`([a-z]+) ([\+\-\*/]) ([a-z]+)`)
	digitRegex := regexp.MustCompile(`\d+`)
	for _, line := range strings.Split(input, "\n") {
		mainMatches := mainRegex.FindStringSubmatch(line)
		mainName := mainMatches[1]
		number := -1
		operand1 := ""
		operand2 := ""
		operation := ""
		if digitRegex.MatchString(mainMatches[2]) {
			number = util.ToInt(mainMatches[2])
		} else {
			expressionMatches := expressionRegex.FindStringSubmatch(mainMatches[2])
			operand1 = expressionMatches[1]
			operand2 = expressionMatches[3]
			operation = expressionMatches[2]
		}
		monkeys[mainName] = &Monkey{number, operand1, operand2, operation}
	}
	return monkeys, time.Since(start)
}
