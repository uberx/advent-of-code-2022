package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/uberx/advent-of-code-2022/util"
)

type State struct {
	numbers    []*int
	numbersMap map[int]*int
}

func main() {
	start := time.Now()
	input := util.ReadFile("day20.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := sumOfNthNumbersAfterZero1(input)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (sumOfNthNumbersAfterZero1): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := sumOfNthNumbersAfterZero2(input)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (sumOfNthNumbersAfterZero2): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func sumOfNthNumbersAfterZero1(input string) (int, time.Duration, time.Duration) {
	state, parseDuration := parseInput(input)

	start := time.Now()
	mixNumbers(&state)

	return util.Sum(nthNumbersAfterZero(&state, 1000, 2000, 3000)), parseDuration, time.Since(start)
}

func mixNumbers(state *State) {
	for idx := 0; idx < len(state.numbersMap); idx++ {
		numberToMove := state.numbersMap[idx]
		currIdx := util.IndexOf(state.numbers, numberToMove, func(t1, t2 *int) bool { return t1 == t2 })
		newIdx := currIdx + *numberToMove
		if newIdx >= len(state.numbersMap) {
			newIdx = newIdx % (len(state.numbersMap) - 1)
		}
		if newIdx < 0 {
			newIdx = len(state.numbersMap) + (newIdx % (len(state.numbersMap) - 1)) - 1
		}
		if currIdx == newIdx {
			continue
		}
		newNumbers := []*int{}
		if currIdx < newIdx {
			newNumbers = append(newNumbers, state.numbers[0:currIdx]...)
			newNumbers = append(newNumbers, state.numbers[currIdx+1:newIdx+1]...)
			newNumbers = append(newNumbers, numberToMove)
			newNumbers = append(newNumbers, state.numbers[newIdx+1:]...)
		} else {
			newNumbers = append(newNumbers, state.numbers[0:newIdx]...)
			newNumbers = append(newNumbers, numberToMove)
			newNumbers = append(newNumbers, state.numbers[newIdx:currIdx]...)
			newNumbers = append(newNumbers, state.numbers[currIdx+1:]...)
		}
		state.numbers = newNumbers
	}
}

func nthNumbersAfterZero(state *State, ns ...int) []int {
	nthNumbers := []int{}
	zeroIdx := -1
	for idx, number := range state.numbers {
		if *number == 0 {
			zeroIdx = idx
		}
	}
	for _, n := range ns {
		nthNumbers = append(nthNumbers, *state.numbers[(zeroIdx+n)%len(state.numbersMap)])
	}
	return nthNumbers
}

func sumOfNthNumbersAfterZero2(input string) (int, time.Duration, time.Duration) {
	state, parseDuration := parseInput(input)

	start := time.Now()
	decryptedNumbers := []*int{}
	for idx, number := range state.numbers {
		decryptedNumber := 811589153 * *number
		decryptedNumbers = append(decryptedNumbers, &decryptedNumber)
		state.numbersMap[idx] = &decryptedNumber
	}
	state.numbers = decryptedNumbers
	for i := 0; i < 10; i++ {
		mixNumbers(&state)
	}
	return util.Sum(nthNumbersAfterZero(&state, 1000, 2000, 3000)), parseDuration, time.Since(start)
}

func parseInput(input string) (State, time.Duration) {
	start := time.Now()
	numbers := []*int{}
	numbersMap := map[int]*int{}
	for idx, line := range strings.Split(input, "\n") {
		number := util.ToInt(line)
		numbers = append(numbers, &number)
		numbersMap[idx] = &number
	}
	return State{numbers, numbersMap}, time.Since(start)
}
