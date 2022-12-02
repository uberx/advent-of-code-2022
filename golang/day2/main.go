package main

import (
	"fmt"
	"strings"

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
	input := util.ReadFile("day2.txt")

	answer1 := totalMyRPSScore1(input)
	fmt.Println("Part 1:", answer1)

	answer2 := totalMyRPSScore2(input)
	fmt.Println("Part 2:", answer2)
}

func totalMyRPSScore1(input string) int {
	rpsRounds := parseInput1(input)

	return totalMyRPSScore(rpsRounds)
}

func totalMyRPSScore2(input string) int {
	rpsRounds := parseInput2(input)

	return totalMyRPSScore(rpsRounds)
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

func parseInput1(input string) (rpsRounds []RockPaperScissorsRound) {
	for _, rpsRound := range strings.Split(input, "\n") {
		roundChoices := strings.Split(rpsRound, " ")
		rpsRounds = append(rpsRounds, RockPaperScissorsRound{apponentChoice: roundChoices[0], myChoice: roundChoices[1]})
	}
	return rpsRounds
}

func parseInput2(input string) (rpsRounds []RockPaperScissorsRound) {
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
	return rpsRounds
}
