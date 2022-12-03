package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/uberx/advent-of-code-2022/util"
)

func main() {
	start := time.Now()
	input := util.ReadFile("day3.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := sumOfPrioritiesOfCommonItemsInRucksack(input)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (sumOfPrioritiesOfCommonItemsInRucksack): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := sumOfPrioritiesOfCommonItemsInRucksackGroups(input)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (sumOfPrioritiesOfCommonItemsInRucksackGroups): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func sumOfPrioritiesOfCommonItemsInRucksack(input string) (int, time.Duration, time.Duration) {
	rucksacks, parseDuration := parseInput(input)

	start := time.Now()
	totalItemPriority := 0
	for _, rucksack := range rucksacks {
		commonItems := commonItems(compartmentAsMap(rucksack[0]), compartmentAsMap(rucksack[1]))
		totalItemPriority += itemPriority(getOnlyElement(commonItems))
	}
	return totalItemPriority, parseDuration, time.Since(start)
}

func commonItems(compartment1 map[rune]bool, compartment2 map[rune]bool) map[rune]bool {
	commonItems := map[rune]bool{}

	for item := range compartment2 {
		if _, ok := compartment1[item]; ok {
			commonItems[item] = true
		}
	}
	return commonItems
}

func compartmentAsMap(compartment string) map[rune]bool {
	compartmentMap := map[rune]bool{}
	for _, item := range compartment {
		compartmentMap[item] = true
	}
	return compartmentMap
}

func getOnlyElement(items map[rune]bool) rune {
	if len(items) > 1 {
		panic("items map has more than 1 element")
	}
	for item := range items {
		return item
	}
	return -1
}

func itemPriority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item) - int('a') + 1
	} else if item >= 'A' && item <= 'Z' {
		return int(item) - int('A') + 27
	}
	return 0
}

func sumOfPrioritiesOfCommonItemsInRucksackGroups(input string) (int, time.Duration, time.Duration) {
	rucksacks, parseDuration := parseInput(input)

	start := time.Now()
	totalItemPriority := 0
	rucksackGroup := []string{}
	for i, rucksack := range rucksacks {
		rucksackGroup = append(rucksackGroup, rucksack[0]+rucksack[1])
		if (i+1)%3 == 0 {
			commonItemsBetweenFirstTwoInGroup := commonItems(compartmentAsMap(rucksackGroup[0]), compartmentAsMap(rucksackGroup[1]))
			commonItemsInGroup := commonItems(commonItemsBetweenFirstTwoInGroup, compartmentAsMap(rucksackGroup[2]))
			badgeItem := getOnlyElement(commonItemsInGroup)
			totalItemPriority += itemPriority(badgeItem)
			rucksackGroup = []string{}
		}
	}
	return totalItemPriority, parseDuration, time.Since(start)
}

func parseInput(input string) ([][]string, time.Duration) {
	start := time.Now()
	rucksacks := [][]string{}
	for _, line := range strings.Split(input, "\n") {
		rucksack := []string{string(line[0 : len(line)/2]), string(line[len(line)/2:])}
		rucksacks = append(rucksacks, rucksack)
	}
	return rucksacks, time.Since(start)
}
