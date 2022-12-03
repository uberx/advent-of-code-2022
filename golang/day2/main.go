package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/uberx/advent-of-code-2022/util"
)

type RockPaperScissorsRound struct {
	apponentChoice string
	myChoice       string
}

const (
	Rock     = "Rock"
	Paper    = "Paper"
	Scissors = "Scissors"
)

var winConditions = map[string]string{
	Rock:     Paper,
	Paper:    Scissors,
	Scissors: Rock,
}

var loseConditions = map[string]string{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

func main() {
	start := time.Now()
	input := util.ReadFile("day2.txt")
	fileReadDuration := time.Since(start)

	answer1, parseDuration1, partDuration1 := totalMyRPSScore1(input)
	totalPartDuration1 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration1.Nanoseconds() + partDuration1.Nanoseconds())
	fmt.Printf("Part 1 (totalMyRPSScore1): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer1, totalPartDuration1, fileReadDuration, parseDuration1, partDuration1)

	answer2, parseDuration2, partDuration2 := totalMyRPSScore2(input)
	totalPartDuration2 := time.Duration(fileReadDuration.Nanoseconds() + parseDuration2.Nanoseconds() + partDuration2.Nanoseconds())
	fmt.Printf("Part 2 (totalMyRPSScore2): %d (%s - fileReadDuration: %s, parseDuration: %s, partDuration: %s)\n", answer2, totalPartDuration2, fileReadDuration, parseDuration2, partDuration2)
}

func totalMyRPSScore1(input string) (int, time.Duration, time.Duration) {
	start := time.Now()
	rpsRounds, parseDuration := parseInput1(input)

	return totalMyRPSScore(rpsRounds), parseDuration, time.Since(start)
}

func totalMyRPSScore2(input string) (int, time.Duration, time.Duration) {
	start := time.Now()
	rpsRounds, parseDuration := parseInput2(input)

	return totalMyRPSScore(rpsRounds), parseDuration, time.Since(start)
}

func totalMyRPSScore(rpsRounds []RockPaperScissorsRound) int {
	totalMyRPSScore := 0
	for _, rpsRound := range rpsRounds {
		totalMyRPSScore += computeMyScore(rpsRound)
	}
	return totalMyRPSScore
}

func computeMyScore(rpsRound RockPaperScissorsRound) int {
	opponentChoice := rpsChoice(rpsRound.apponentChoice)
	myChoice := rpsChoice(rpsRound.myChoice)

	var baseScore int
	switch myChoice {
	case Rock:
		baseScore = 1
	case Paper:
		baseScore = 2
	case Scissors:
		baseScore = 3
	}

	var roundScore int
	if opponentChoice == myChoice {
		roundScore = 3
	} else if myChoice == winConditions[opponentChoice] {
		roundScore = 6
	}

	return baseScore + roundScore
}

func rpsChoice(choice string) string {
	switch choice {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	}
	return choice
}

func parseInput1(input string) ([]RockPaperScissorsRound, time.Duration) {
	start := time.Now()
	rpsRounds := []RockPaperScissorsRound{}
	for _, rpsRound := range strings.Split(input, "\n") {
		roundChoices := strings.Split(rpsRound, " ")
		rpsRounds = append(rpsRounds, RockPaperScissorsRound{apponentChoice: roundChoices[0], myChoice: roundChoices[1]})
	}
	return rpsRounds, time.Since(start)
}

func parseInput2(input string) ([]RockPaperScissorsRound, time.Duration) {
	start := time.Now()
	rpsRounds := []RockPaperScissorsRound{}
	for _, rpsRound := range strings.Split(input, "\n") {
		roundChoices := strings.Split(rpsRound, " ")

		opponentChoice := rpsChoice(roundChoices[0])
		myChoice := roundChoices[1]
		if myChoice == "X" {
			myChoice = loseConditions[opponentChoice]
		} else if myChoice == "Y" {
			myChoice = opponentChoice
		} else if myChoice == "Z" {
			myChoice = winConditions[opponentChoice]
		}
		rpsRounds = append(rpsRounds, RockPaperScissorsRound{opponentChoice, myChoice})
	}
	return rpsRounds, time.Since(start)
}
