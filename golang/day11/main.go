package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/uberx/advent-of-code-2022/util"
)

type OperationType string

const (
	multiply OperationType = "old * "
	add      OperationType = "old + "
	square   OperationType = "old * old"
)

type Operation struct {
	operationType OperationType
	operand       int
}

func (o Operation) perform(input int) int {
	if o.operationType == multiply {
		return input * o.operand
	} else if o.operationType == add {
		return input + o.operand
	} else if o.operationType == square {
		return input * input
	}
	panic("unexpected operation")
}

type Monkey struct {
	items             []int
	operation         Operation
	divisibilityTest  int
	testSuccessMonkey int
	testFailMonkey    int
	inspections       int
}

func main() {
	start := time.Now()
	input := util.ReadFile("day11.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := monkeyBusiness(input, 20, 3)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (monkeyBusiness1): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := monkeyBusiness(input, 10000, 1)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (monkeyBusiness2): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func monkeyBusiness(input string, rounds int, worryFactor int) (int, time.Duration, time.Duration) {
	monkeys, parseDuration := parseInput(input)

	start := time.Now()
	productOfMonkeyDivisibilityTests := 1
	for _, monkey := range monkeys {
		productOfMonkeyDivisibilityTests *= monkey.divisibilityTest
	}

	for round := 1; round <= rounds; round++ {
		for monkeyIdx := 0; monkeyIdx < len(monkeys); monkeyIdx++ {
			monkey := &monkeys[monkeyIdx]
			for itemIdx := 0; itemIdx < len(monkey.items); itemIdx++ {
				item := monkey.items[itemIdx]
				itemWorryLevel := monkey.operation.perform(item) / worryFactor
				itemWorryLevel %= productOfMonkeyDivisibilityTests

				var recipientMonkeyIdx int
				if itemWorryLevel%monkey.divisibilityTest == 0 {
					recipientMonkeyIdx = monkey.testSuccessMonkey
				} else {
					recipientMonkeyIdx = monkey.testFailMonkey
				}
				monkeys[recipientMonkeyIdx].items = append(monkeys[recipientMonkeyIdx].items, itemWorryLevel)
				monkey.inspections = monkey.inspections + 1
			}
			monkey.items = []int{}
		}
	}

	inspections := []int{}
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.inspections)
	}
	sort.Ints(inspections)

	return inspections[len(inspections)-1] * inspections[len(inspections)-2], parseDuration, time.Since(start)
}

func parseInput(input string) ([]Monkey, time.Duration) {
	start := time.Now()
	monkeys := []Monkey{}
	for _, monkeyBlock := range strings.Split(input, "\n\n") {
		monkeyLines := strings.Split(monkeyBlock, "\n")

		items := []int{}
		itemsLine := strings.Split(monkeyLines[1][len("  Starting items: "):], ", ")
		for _, item := range itemsLine {
			items = append(items, util.ToInt(item))
		}

		var operationType OperationType
		operand := -1
		operationLine := monkeyLines[2][len("  Operation: new = "):]
		if strings.HasPrefix(operationLine, string(square)) {
			operationType = square
		} else if strings.HasPrefix(operationLine, string(add)) {
			operationType = add
			operand = util.ToInt(operationLine[len(add):])
		} else if strings.HasPrefix(operationLine, string(multiply)) {
			operationType = multiply
			operand = util.ToInt(operationLine[len(multiply):])
		}

		divisibilityTest := util.ToInt(monkeyLines[3][len("  Test: divisible by "):])
		testSuccessMonkey := util.ToInt(monkeyLines[4][len("    If true: throw to monkey "):])
		testFailMonkey := util.ToInt(monkeyLines[5][len("    If false: throw to monkey "):])

		monkeys = append(monkeys, Monkey{items, Operation{operationType, operand}, divisibilityTest, testSuccessMonkey, testFailMonkey, 0})
	}
	return monkeys, time.Since(start)
}
