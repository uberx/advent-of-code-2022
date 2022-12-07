package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/uberx/advent-of-code-2022/util"
)

type TreeNode struct {
	name      string
	dir       bool
	parentDir *TreeNode
	children  map[string]*TreeNode
	size      int
	depth     int
}

func NewDir(name string, parentDir *TreeNode, depth int) *TreeNode {
	return &TreeNode{name, true, parentDir, map[string]*TreeNode{}, -1, depth}
}

func NewFile(name string, parentDir *TreeNode, size int, depth int) *TreeNode {
	return &TreeNode{name, false, parentDir, nil, size, depth}
}

func main() {
	start := time.Now()
	input := util.ReadFile("day7.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := totalSizeOfDirsWithAtmostGivenSize(input, 100000)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (totalSizeOfDirsWithAtmostGivenSize): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := smallestSizeOfDirToDelete(input, 70000000, 30000000)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (smallestSizeOfDirToDelete): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func totalSizeOfDirsWithAtmostGivenSize(input string, maxSize int) (int, time.Duration, time.Duration) {
	rootNode, parseDuration := parseInput(input)

	start := time.Now()
	dirSizes := map[*TreeNode]int{}
	computeSize(rootNode, dirSizes)

	totalSizeOfDirsWithAtmostGivenSize := 0
	for _, dirSize := range dirSizes {
		if dirSize <= maxSize {
			totalSizeOfDirsWithAtmostGivenSize += dirSize
		}
	}

	return totalSizeOfDirsWithAtmostGivenSize, parseDuration, time.Since(start)
}

func smallestSizeOfDirToDelete(input string, totalSpace int, requiredSpace int) (int, time.Duration, time.Duration) {
	rootNode, parseDuration := parseInput(input)

	start := time.Now()
	dirSizes := map[*TreeNode]int{}
	rootSize := computeSize(rootNode, dirSizes)

	freeSpace := totalSpace - rootSize
	spaceToCleanUp := requiredSpace - freeSpace

	smallestSizeOfDirToDelete := math.MaxInt
	for _, dirSize := range dirSizes {
		if dirSize >= spaceToCleanUp && dirSize < smallestSizeOfDirToDelete {
			smallestSizeOfDirToDelete = dirSize
		}
	}
	return smallestSizeOfDirToDelete, parseDuration, time.Since(start)
}

func computeSize(treeNode *TreeNode, dirSizes map[*TreeNode]int) int {
	// file
	if !treeNode.dir {
		return treeNode.size
	}

	// dir
	if treeNode.size == -1 {
		size := 0
		for _, child := range treeNode.children {
			childSize := computeSize(child, dirSizes)
			if child.dir {
				child.size = childSize
			}
			size += childSize
		}
		dirSizes[treeNode] = size
		return size
	}
	return treeNode.size
}

func parseInput(input string) (*TreeNode, time.Duration) {
	start := time.Now()

	rootNode := NewDir("/", nil, 0)
	listMode := false
	depth := 0
	var currentDir *TreeNode
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "$ cd /") {
			listMode = false
			depth = 0
			currentDir = rootNode
		} else if strings.HasPrefix(line, "$ ls") {
			listMode = true
		} else if strings.HasPrefix(line, "$ cd ..") {
			listMode = false
			depth--
			currentDir = currentDir.parentDir
		} else if strings.HasPrefix(line, "$ cd ") {
			listMode = false
			depth++
			currentDir = currentDir.children[line[5:]]
		} else if listMode {
			if strings.HasPrefix(line, "dir ") {
				currentDir.children[line[4:]] = NewDir(line[4:], currentDir, depth)
			} else {
				fileLine := strings.Split(line, " ")
				currentDir.children[fileLine[1]] = NewFile(fileLine[1], currentDir, util.ToInt(fileLine[0]), depth)
			}
		}
	}

	return rootNode, time.Since(start)
}
