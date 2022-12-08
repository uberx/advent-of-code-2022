package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/uberx/advent-of-code-2022/util"
)

func main() {
	start := time.Now()
	input := util.ReadFile("day8.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := numVisibleTrees(input)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (numVisibleTrees): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := highestTreeScenicScore(input)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (highestTreeScenicScore): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func numVisibleTrees(input string) (int, time.Duration, time.Duration) {
	treeHeightGrid, parseDuration := parseInput(input)

	start := time.Now()
	rows := len(treeHeightGrid)
	cols := len(treeHeightGrid[0])
	edgeTrees := (cols * 2) + (2 * (rows - 2))
	visibleTrees := 0
	for row := 1; row < rows-1; row++ {
		for col := 1; col < cols-1; col++ {
			if isTreeVisible(treeHeightGrid, rows, cols, row, col) {
				visibleTrees++
			}
		}
	}

	return edgeTrees + visibleTrees, parseDuration, time.Since(start)
}

func isTreeVisible(treeHeightGrid [][]int, rows, cols, rowNum, colNum int) bool {
	// top
	invisibleFromTop := false
	for row := 0; row < rowNum && !invisibleFromTop; row++ {
		if treeHeightGrid[row][colNum] >= treeHeightGrid[rowNum][colNum] {
			invisibleFromTop = true
		}
	}

	// bottom
	invisibleFromBottom := false
	for row := rowNum + 1; row < rows && !invisibleFromBottom; row++ {
		if treeHeightGrid[row][colNum] >= treeHeightGrid[rowNum][colNum] {
			invisibleFromBottom = true
		}
	}

	// left
	invisibleFromLeft := false
	for col := 0; col < colNum && !invisibleFromLeft; col++ {
		if treeHeightGrid[rowNum][col] >= treeHeightGrid[rowNum][colNum] {
			invisibleFromLeft = true
		}
	}

	// right
	invisibleFromRight := false
	for col := colNum + 1; col < cols && !invisibleFromRight; col++ {
		if treeHeightGrid[rowNum][col] >= treeHeightGrid[rowNum][colNum] {
			invisibleFromRight = true
		}
	}

	return !(invisibleFromTop && invisibleFromBottom && invisibleFromLeft && invisibleFromRight)
}

func highestTreeScenicScore(input string) (int, time.Duration, time.Duration) {
	treeHeightGrid, parseDuration := parseInput(input)

	start := time.Now()
	rows := len(treeHeightGrid)
	cols := len(treeHeightGrid[0])
	maxScenicScore := math.MinInt
	for row := 0; row < rows-1; row++ {
		for col := 0; col < cols-1; col++ {
			scenicScore := computeScenicScore(treeHeightGrid, rows, cols, row, col)
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	return maxScenicScore, parseDuration, time.Since(start)
}

func computeScenicScore(treeHeightGrid [][]int, rows, cols, rowNum, colNum int) int {
	// top
	topViewingDistance := 0
	for row := rowNum - 1; row >= 0; row-- {
		topViewingDistance++
		if treeHeightGrid[row][colNum] >= treeHeightGrid[rowNum][colNum] {
			break
		}
	}

	// bottom
	bottomViewingDistance := 0
	for row := rowNum + 1; row < rows; row++ {
		bottomViewingDistance++
		if treeHeightGrid[row][colNum] >= treeHeightGrid[rowNum][colNum] {
			break
		}
	}

	// left
	leftViewingDistance := 0
	for col := colNum - 1; col >= 0; col-- {
		leftViewingDistance++
		if treeHeightGrid[rowNum][col] >= treeHeightGrid[rowNum][colNum] {
			break
		}
	}

	// right
	rightViewingDistance := 0
	for col := colNum + 1; col < cols; col++ {
		rightViewingDistance++
		if treeHeightGrid[rowNum][col] >= treeHeightGrid[rowNum][colNum] {
			break
		}
	}

	return topViewingDistance * bottomViewingDistance * leftViewingDistance * rightViewingDistance
}

func parseInput(input string) ([][]int, time.Duration) {
	start := time.Now()
	treeHeightGrid := [][]int{}
	for _, rowLine := range strings.Split(input, "\n") {
		rowHeights := []int{}
		for _, height := range rowLine {
			rowHeights = append(rowHeights, util.ToInt(string(height)))
		}
		treeHeightGrid = append(treeHeightGrid, rowHeights)
	}
	return treeHeightGrid, time.Since(start)
}
