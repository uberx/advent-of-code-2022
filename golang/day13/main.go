package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/uberx/advent-of-code-2022/util"
)

type ListItem struct {
	value *int
	items *[]ListItem
}

func (l ListItem) isNumber() bool {
	return l.value != nil
}

func (l ListItem) isList() bool {
	return l.value == nil
}

func main() {
	start := time.Now()
	input := util.ReadFile("day13.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := sumOfIndicesOfPacketPairsInRightOrder(input)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (sumOfIndicesOfPacketPairsInRightOrder): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := distressSignalDecoderKey(input)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (distressSignalDecoderKey): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func sumOfIndicesOfPacketPairsInRightOrder(input string) (int, time.Duration, time.Duration) {
	packetPairs, parseDuration := parseInput(input)

	start := time.Now()
	sumOfIndicesOfPacketPairsInRightOrder := 0
	for idx, packetPair := range packetPairs {
		leftPacket := packetPair.First
		rightPacket := packetPair.Second
		if order, _ := isPacketPairInRightOrder(leftPacket, rightPacket); order {
			sumOfIndicesOfPacketPairsInRightOrder += idx + 1
		}
	}
	return sumOfIndicesOfPacketPairsInRightOrder, parseDuration, time.Since(start)
}

func isPacketPairInRightOrder(leftPacket *ListItem, rightPacket *ListItem) (inOrder bool, skip bool) {
	if leftPacket.isNumber() && rightPacket.isNumber() {
		return *leftPacket.value < *rightPacket.value, *leftPacket.value == *rightPacket.value
	}
	if leftPacket.isList() && rightPacket.isList() {
		leftPacketItems := *leftPacket.items
		rightPacketItems := *rightPacket.items
		for i := 0; i < len(*leftPacket.items); i++ {
			if i > len(*rightPacket.items)-1 {
				return inOrder, skip
			}
			if inOrder, skip := isPacketPairInRightOrder(&leftPacketItems[i], &rightPacketItems[i]); !skip {
				return inOrder, false
			} else {
				continue
			}
		}
		if len(rightPacketItems) > len(leftPacketItems) {
			return true, skip
		}
		return inOrder, true
	} else if leftPacket.isList() != rightPacket.isList() {
		if leftPacket.isNumber() {
			newItems := []ListItem{{leftPacket.value, &[]ListItem{}}}
			return isPacketPairInRightOrder(&ListItem{nil, &newItems}, rightPacket)
		} else if rightPacket.isNumber() {
			newItems := []ListItem{{rightPacket.value, &[]ListItem{}}}
			return isPacketPairInRightOrder(leftPacket, &ListItem{nil, &newItems})
		}
	}
	return true, skip
}

func distressSignalDecoderKey(input string) (int, time.Duration, time.Duration) {
	packetPairs, parseDuration := parseInput(input)

	start := time.Now()
	packets := []*ListItem{}
	for _, packetPair := range packetPairs {
		packets = append(packets, packetPair.First, packetPair.Second)
	}
	dividerPacketPairs, _ := parseInput("[[2]]\n[[6]]")
	dividerPacketPair := dividerPacketPairs[0]
	packets = append(packets, dividerPacketPair.First, dividerPacketPair.Second)
	sort.Slice(packets, func(i, j int) bool {
		inOrder, _ := isPacketPairInRightOrder(packets[i], packets[j])
		return inOrder
	})
	var decoderPacketIdx1 int
	var decoderPacketIdx2 int
	for idx, packet := range packets {
		if packet == dividerPacketPair.First {
			decoderPacketIdx1 = idx + 1
		} else if packet == dividerPacketPair.Second {
			decoderPacketIdx2 = idx + 1
		}
	}

	return decoderPacketIdx1 * decoderPacketIdx2, parseDuration, time.Since(start)
}

func parseInput(input string) ([]util.Pair[*ListItem, *ListItem], time.Duration) {
	start := time.Now()
	packetPairs := []util.Pair[*ListItem, *ListItem]{}
	for _, packetPair := range strings.Split(input, "\n\n") {
		packetPairLines := strings.Split(packetPair, "\n")
		packetPairs = append(packetPairs, util.Pair[*ListItem, *ListItem]{parsePacket(packetPairLines[0]), parsePacket(packetPairLines[1])})
	}
	return packetPairs, time.Since(start)
}

func parsePacket(packet string) *ListItem {
	stack := util.Stack[*ListItem]{}
	number := ""
	var currOutermostItem *ListItem
	for _, char := range packet {
		if char == '[' {
			stack.Push(&ListItem{nil, &[]ListItem{}})
		} else if char >= '0' && char <= '9' {
			number += string(char)
		} else if char == ',' || char == ']' {
			if len(number) > 0 {
				curr, _ := stack.Peek()
				num := util.ToInt(number)
				*curr.items = append(*curr.items, ListItem{&num, &[]ListItem{}})
				number = ""
			}
			if char == ']' {
				currOutermostItem, _ = stack.Pop()
				if !stack.IsEmpty() {
					item, _ := stack.Peek()
					*item.items = append(*item.items, *currOutermostItem)
				}
			}
		}
	}

	if currOutermostItem == nil {
		num := util.ToInt(number)
		return &ListItem{&num, &[]ListItem{}}
	}
	return currOutermostItem
}
