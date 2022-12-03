package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/uberx/advent-of-code-2022/util"
)

func main() {
	start := time.Now()
	input := util.ReadFile("day1.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := maxTotalCalories(input)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (maxTotalCalories): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := sumOfTopThreeTotalCalories(input)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (sumOfTopThreeTotalCalories): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func maxTotalCalories(input string) (int, time.Duration, time.Duration) {
	elfItemCaloriesList, parseDuration := parseInput(input)

	start := time.Now()
	maxTotalCalories := math.MinInt
	for _, elfItemCalories := range elfItemCaloriesList {
		totalCalories := totalCalories(elfItemCalories)
		if totalCalories > maxTotalCalories {
			maxTotalCalories = totalCalories
		}
	}

	return maxTotalCalories, parseDuration, time.Since(start)
}

func totalCalories(elfItemCalories []int) int {
	var totalCalories int
	for _, itemCalory := range elfItemCalories {
		totalCalories += itemCalory
	}
	return totalCalories
}

func sumOfTopThreeTotalCalories(input string) (int, time.Duration, time.Duration) {
	elfItemCaloriesList, parseDuration := parseInput(input)

	start := time.Now()
	elfTotalCalories := []int{}
	for _, elfItemCalories := range elfItemCaloriesList {
		elfTotalCalories = append(elfTotalCalories, totalCalories(elfItemCalories))
	}

	sort.Ints(elfTotalCalories)
	return elfTotalCalories[len(elfTotalCalories)-1] + elfTotalCalories[len(elfTotalCalories)-2] + elfTotalCalories[len(elfTotalCalories)-3], parseDuration, time.Since(start)
}

func parseInput(input string) ([][]int, time.Duration) {
	start := time.Now()
	elfItemCaloriesList := [][]int{}
	elfItemCalories := []int{}
	for _, currItem := range strings.Split(input, "\n") {
		if currItem == "" {
			elfItemCaloriesList = append(elfItemCaloriesList, elfItemCalories)
			elfItemCalories = []int{}
		} else {
			elfItemCalories = append(elfItemCalories, util.ToInt(currItem))
		}
	}
	return elfItemCaloriesList, time.Since(start)
}
