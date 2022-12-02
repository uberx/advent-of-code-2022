package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/uberx/advent-of-code-2022/util"
)

func main() {
	input := util.ReadFile("day1.txt")

	answer1 := maxTotalCalories(input)
	fmt.Println("Part 1:", answer1)

	answer2 := sumOfTopThreeTotalCalories(input)
	fmt.Println("Part 2:", answer2)
}

func maxTotalCalories(input string) int {
	elfItemCaloriesList := parseInput(input)

	maxTotalCalories := math.MinInt
	for _, elfItemCalories := range elfItemCaloriesList {
		totalCalories := totalCalories(elfItemCalories)
		if totalCalories > maxTotalCalories {
			maxTotalCalories = totalCalories
		}
	}

	return maxTotalCalories
}

func totalCalories(elfItemCalories []int) int {
	var totalCalories int
	for _, itemCalory := range elfItemCalories {
		totalCalories += itemCalory
	}
	return totalCalories
}

func sumOfTopThreeTotalCalories(input string) int {
	elfItemCaloriesList := parseInput(input)

	elfTotalCalories := []int{}
	for _, elfItemCalories := range elfItemCaloriesList {
		elfTotalCalories = append(elfTotalCalories, totalCalories(elfItemCalories))
	}

	sort.Ints(elfTotalCalories)
	return elfTotalCalories[len(elfTotalCalories)-1] + elfTotalCalories[len(elfTotalCalories)-2] + elfTotalCalories[len(elfTotalCalories)-3]
}

func parseInput(input string) (elfItemCaloriesList [][]int) {
	elfItemCalories := []int{}
	for _, currItem := range strings.Split(input, "\n") {
		if currItem == "" {
			elfItemCaloriesList = append(elfItemCaloriesList, elfItemCalories)
			elfItemCalories = []int{}
		} else {
			elfItemCalories = append(elfItemCalories, util.ToInt(currItem))
		}
	}
	return elfItemCaloriesList
}
