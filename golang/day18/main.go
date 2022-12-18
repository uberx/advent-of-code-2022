package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/uberx/advent-of-code-2022/util"
)

type State struct {
	cubes  []util.Point3D
	xRange util.Point
	yRange util.Point
	zRange util.Point
}

func main() {
	start := time.Now()
	input := util.ReadFile("day18.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := lavaDropletsSurfaceArea(input)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (lavaDropletsSurfaceArea): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := lavaDropletsExternalSurfaceArea(input)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (lavaDropletsExternalSurfaceArea): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func lavaDropletsSurfaceArea(input string) (int, time.Duration, time.Duration) {
	state, parseDuration := parseInput(input)

	start := time.Now()
	totalSurfaceArea, _ := totalSurfaceArea(state)
	return totalSurfaceArea, parseDuration, time.Since(start)
}

func totalSurfaceArea(state State) (int, map[int]map[int]map[int]bool) {
	cubeCache := map[int]map[int]map[int]bool{}
	totalSurfaceArea := 0
	for _, cube := range state.cubes {
		updateCubeCache(cube, cubeCache)
		totalSurfaceArea += surfaceArea(cube, cubeCache)
	}
	return totalSurfaceArea, cubeCache
}

func updateCubeCache(cube util.Point3D, cubeCache map[int]map[int]map[int]bool) {
	if _, ok := cubeCache[cube.X]; !ok {
		cubeCache[cube.X] = map[int]map[int]bool{}
	}
	if _, ok := cubeCache[cube.X][cube.Y]; !ok {
		cubeCache[cube.X][cube.Y] = map[int]bool{}
	}
	cubeCache[cube.X][cube.Y][cube.Z] = true
}

func surfaceArea(cube util.Point3D, cubeCache map[int]map[int]map[int]bool) int {
	surfaceArea := 6
	if cubeCache[cube.X-1][cube.Y][cube.Z] {
		surfaceArea -= 2
	}
	if cubeCache[cube.X+1][cube.Y][cube.Z] {
		surfaceArea -= 2
	}
	if cubeCache[cube.X][cube.Y-1][cube.Z] {
		surfaceArea -= 2
	}
	if cubeCache[cube.X][cube.Y+1][cube.Z] {
		surfaceArea -= 2
	}
	if cubeCache[cube.X][cube.Y][cube.Z-1] {
		surfaceArea -= 2
	}
	if cubeCache[cube.X][cube.Y][cube.Z+1] {
		surfaceArea -= 2
	}
	return surfaceArea
}

func lavaDropletsExternalSurfaceArea(input string) (int, time.Duration, time.Duration) {
	state, parseDuration := parseInput(input)

	start := time.Now()
	totalSurfaceArea, cubeCache := totalSurfaceArea(state)
	bfs := util.Queue[util.Point3D]{}
	bfs.Enqueue(util.Point3D{X: state.xRange.X - 1, Y: state.yRange.X - 1, Z: state.zRange.X - 1})
	visited := map[util.Point3D]bool{}
	iterations := 0
	for !bfs.IsEmpty() {
		currPoint, _ := bfs.Dequeue()
		if visited[currPoint] {
			continue
		}
		iterations++
		visited[currPoint] = true
		currNeighbors := neighbors(currPoint, cubeCache, state)
		for _, currNeighbor := range currNeighbors {
			if !visited[currNeighbor] {
				bfs.Enqueue(currNeighbor)
			}
		}
	}

	totalInternalSurfaces := 0
	for x := state.xRange.X; x <= state.xRange.Y; x++ {
		for y := state.yRange.X; y <= state.yRange.Y; y++ {
			for z := state.zRange.X; z <= state.zRange.Y; z++ {
				if cubeCache[x][y][z] {
					continue
				}
				cube := util.Point3D{X: x, Y: y, Z: z}
				if !visited[cube] {
					totalInternalSurfaces += internalSurfaces(cube, cubeCache)
				}
			}
		}
	}
	return totalSurfaceArea - totalInternalSurfaces, parseDuration, time.Since(start)
}

func neighbors(currPoint util.Point3D, cubeCache map[int]map[int]map[int]bool, state State) []util.Point3D {
	neighborOffsets := []util.Point3D{
		{X: -1, Y: 0, Z: 0},
		{X: 1, Y: 0, Z: 0},
		{X: 0, Y: -1, Z: 0},
		{X: 0, Y: 1, Z: 0},
		{X: 0, Y: 0, Z: -1},
		{X: 0, Y: 0, Z: 1},
	}
	neighbors := []util.Point3D{}
	for _, neighborOffset := range neighborOffsets {
		neighbor := util.Point3D{X: currPoint.X + neighborOffset.X, Y: currPoint.Y + neighborOffset.Y, Z: currPoint.Z + neighborOffset.Z}
		if neighbor.X >= state.xRange.X-1 && neighbor.X <= state.xRange.Y+1 &&
			neighbor.Y >= state.yRange.X-1 && neighbor.Y <= state.yRange.Y+1 &&
			neighbor.Z >= state.zRange.X-1 && neighbor.Z <= state.zRange.Y+1 &&
			!cubeCache[neighbor.X][neighbor.Y][neighbor.Z] {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func internalSurfaces(cube util.Point3D, cubeCache map[int]map[int]map[int]bool) int {
	internalSurfaces := 0
	if _, ok := cubeCache[cube.X-1][cube.Y][cube.Z]; ok {
		internalSurfaces++
	}
	if _, ok := cubeCache[cube.X+1][cube.Y][cube.Z]; ok {
		internalSurfaces++
	}
	if _, ok := cubeCache[cube.X][cube.Y-1][cube.Z]; ok {
		internalSurfaces++
	}
	if _, ok := cubeCache[cube.X][cube.Y+1][cube.Z]; ok {
		internalSurfaces++
	}
	if _, ok := cubeCache[cube.X][cube.Y][cube.Z-1]; ok {
		internalSurfaces++
	}
	if _, ok := cubeCache[cube.X][cube.Y][cube.Z+1]; ok {
		internalSurfaces++
	}
	return internalSurfaces
}

func parseInput(input string) (State, time.Duration) {
	start := time.Now()
	cubes := []util.Point3D{}
	xRange := util.Point{X: math.MaxInt, Y: math.MinInt}
	yRange := util.Point{X: math.MaxInt, Y: math.MinInt}
	zRange := util.Point{X: math.MaxInt, Y: math.MinInt}
	for _, line := range strings.Split(input, "\n") {
		coords := strings.Split(line, ",")
		cube := util.Point3D{X: util.ToInt(coords[0]), Y: util.ToInt(coords[1]), Z: util.ToInt(coords[2])}
		cubes = append(cubes, cube)
		if cube.X < xRange.X {
			xRange.X = cube.X
		}
		if cube.X > xRange.Y {
			xRange.Y = cube.X
		}
		if cube.Y < yRange.X {
			yRange.X = cube.Y
		}
		if cube.Y > yRange.Y {
			yRange.Y = cube.Y
		}
		if cube.Z < zRange.X {
			zRange.X = cube.Z
		}
		if cube.Z > zRange.Y {
			zRange.Y = cube.Z
		}
	}
	return State{cubes, xRange, yRange, zRange}, time.Since(start)
}
