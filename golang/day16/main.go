package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/uberx/advent-of-code-2022/util"
)

type Valve struct {
	name     string
	flowRate int
	tunnels  []string
}

type StackItem1 struct {
	currValve    *Valve
	timeSpent    int
	pressure     int
	valvesOpened map[string]int
}

type StackItem2 struct {
	myCurrValve       *Valve
	elephantCurrValve *Valve
	myTimeSpent       int
	elephantTimeSpent int
	pressure          int
	valvesOpened      map[string]int
}

type QueueItem struct {
	currValve *Valve
	timeSpent int
}

func main() {
	start := time.Now()
	input := util.ReadFile("day16.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := maxPressure1(input)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (maxPressure1): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := maxPressure2(input)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (maxPressure2): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func maxPressure1(input string) (int, time.Duration, time.Duration) {
	valves, parseDuration := parseInput(input)

	start := time.Now()
	openableValves := openableValves(valves)
	minTimeBetweenOpenableValves := minTimeBetweenOpenableValves(valves, openableValves)
	maxPressureItem := maxPressureItem1(valves, minTimeBetweenOpenableValves, 30)
	return maxPressureItem.pressure, parseDuration, time.Since(start)
}

func openableValves(valves map[string]*Valve) map[string]bool {
	openableValves := map[string]bool{}
	for _, valve := range valves {
		if valve.flowRate > 0 {
			openableValves[valve.name] = true
		}
	}
	return openableValves
}

func minTimeBetweenOpenableValves(valves map[string]*Valve, openableValves map[string]bool) map[string]map[string]int {
	minTimeBetweenOpenableValves := map[string]map[string]int{}
	openableFromValves := util.CopyMap(openableValves)
	openableFromValves["AA"] = true
	for fromValveName := range openableFromValves {
		for toValveName := range openableValves {
			if fromValveName != toValveName {
				visitedValves := map[string]int{}
				queue := util.Queue[QueueItem]{}
				queue.Enqueue(QueueItem{valves[fromValveName], 0})
				for !queue.IsEmpty() {
					currItem, _ := queue.Dequeue()
					if currItem.currValve.name == toValveName {
						if _, ok := minTimeBetweenOpenableValves[fromValveName]; ok {
							minTimeBetweenOpenableValves[fromValveName][toValveName] = currItem.timeSpent
						} else {
							minTimeBetweenOpenableValves[fromValveName] = map[string]int{toValveName: currItem.timeSpent}
						}
						break
					}
					visitedValves[currItem.currValve.name] = currItem.timeSpent

					// possible moves
					for _, tunnel := range currItem.currValve.tunnels {
						if existingTimeSpent, ok := visitedValves[tunnel]; !ok || currItem.timeSpent+1 < existingTimeSpent {
							queue.Enqueue(QueueItem{valves[tunnel], currItem.timeSpent + 1})
						}
					}
				}
			}
		}
	}
	return minTimeBetweenOpenableValves
}

func maxPressureItem1(valves map[string]*Valve, minTimeBetweenOpenableValves map[string]map[string]int, totalTime int) StackItem1 {
	var maxPressureItem StackItem1
	stack := util.Stack[StackItem1]{}
	stack.Push(StackItem1{valves["AA"], 1, 0, map[string]int{}})
	for !stack.IsEmpty() {
		currItem, _ := stack.Pop()
		if currItem.pressure > maxPressureItem.pressure {
			maxPressureItem = currItem
		}
		// possible (move + open)s
		for toValveName, timeCost := range minTimeBetweenOpenableValves[currItem.currValve.name] {
			if _, ok := currItem.valvesOpened[toValveName]; !ok {
				newTimeSpent := currItem.timeSpent + timeCost
				if newTimeSpent+1 > totalTime {
					continue
				}
				valvesOpenedCopy := util.CopyMap(currItem.valvesOpened)
				valvesOpenedCopy[toValveName] = newTimeSpent
				newPressure := (totalTime - newTimeSpent) * valves[toValveName].flowRate
				stack.Push(StackItem1{valves[toValveName], newTimeSpent + 1, currItem.pressure + newPressure, valvesOpenedCopy})
			}
		}
	}
	return maxPressureItem
}

func maxPressure2(input string) (int, time.Duration, time.Duration) {
	valves, parseDuration := parseInput(input)

	start := time.Now()
	openableValves := openableValves(valves)
	minTimeBetweenOpenableValves := minTimeBetweenOpenableValves(valves, openableValves)
	maxPressureItem := maxPressureItem2(valves, minTimeBetweenOpenableValves, 26)
	return maxPressureItem.pressure, parseDuration, time.Since(start)
}

func maxPressureItem2(valves map[string]*Valve, minTimeBetweenOpenableValves map[string]map[string]int, totalTime int) StackItem2 {
	var maxPressureItem StackItem2
	stack := util.Stack[StackItem2]{}
	stack.Push(StackItem2{valves["AA"], valves["AA"], 1, 1, 0, map[string]int{}})
	for !stack.IsEmpty() {
		currItem, _ := stack.Pop()
		if currItem.pressure > maxPressureItem.pressure {
			maxPressureItem = currItem
		}
		// possible (move + open)s
		minTimeBetweenMyValveAndOpenableValves := map[string]int{}
		for openableValveName, timeCost := range minTimeBetweenOpenableValves[currItem.myCurrValve.name] {
			if _, ok := currItem.valvesOpened[openableValveName]; !ok && currItem.myTimeSpent+timeCost+1 <= totalTime {
				minTimeBetweenMyValveAndOpenableValves[openableValveName] = timeCost
			}
		}
		minTimeBetweenElephantValveAndOpenableValves := map[string]int{}
		for openableValveName, timeCost := range minTimeBetweenOpenableValves[currItem.elephantCurrValve.name] {
			if _, ok := currItem.valvesOpened[openableValveName]; !ok && currItem.elephantTimeSpent+timeCost+1 <= totalTime {
				minTimeBetweenElephantValveAndOpenableValves[openableValveName] = timeCost
			}
		}

		for myToValveName, myTimeCost := range minTimeBetweenMyValveAndOpenableValves {
			for elephantToValveName, elephantTimeCost := range minTimeBetweenElephantValveAndOpenableValves {
				if myToValveName != elephantToValveName {
					newMyTimeSpent := currItem.myTimeSpent + myTimeCost
					newMyPressure := (totalTime - newMyTimeSpent) * valves[myToValveName].flowRate
					newElephantTimeSpent := currItem.elephantTimeSpent + elephantTimeCost
					newElephantPressure := (totalTime - newElephantTimeSpent) * valves[elephantToValveName].flowRate

					valvesOpenedCopy := util.CopyMap(currItem.valvesOpened)
					valvesOpenedCopy[myToValveName] = newMyTimeSpent
					valvesOpenedCopy[elephantToValveName] = newElephantTimeSpent
					stack.Push(StackItem2{valves[myToValveName], valves[elephantToValveName], newMyTimeSpent + 1, newElephantTimeSpent + 1, currItem.pressure + newMyPressure + newElephantPressure, valvesOpenedCopy})
				}
			}
		}
	}
	return maxPressureItem
}

func parseInput(input string) (map[string]*Valve, time.Duration) {
	start := time.Now()
	valves := map[string]*Valve{}
	r := regexp.MustCompile(`Valve ([A-Z]{2}) has flow rate=(\d+); tunnel[s]* lead[s]* to valve[s]* ([A-Z, ]+)`)
	for _, line := range strings.Split(input, "\n") {
		matches := r.FindStringSubmatch(line)
		name := matches[1]
		flowRate := util.ToInt(matches[2])
		tunnels := strings.Split(matches[3], ", ")
		valves[name] = &Valve{name, flowRate, tunnels}
	}
	return valves, time.Since(start)
}
