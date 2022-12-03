package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Rock    string = "A"
	Paper          = "B"
	Scissor        = "C"
)

const (
	Lose string = "X"
	Draw        = "Y"
	Win         = "Z"
)

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	strategy := string(input)
	score := calculateScore(strategy)
	fmt.Printf("Answer : %d", score)
}

func calculateScore(strategy string) int {
	round := strings.Split(strategy, "\r\n")
	score := 0
	for _, r := range round {
		choose := strings.Split(r, " ")
		oponent, player := choose[0], choose[1]
		score += decideMove(oponent, player)
	}
	return score
}

func decideMove(oponent string, player string) int {
	score := 0
	switch player {
	case Lose:
		if oponent == Rock {
			score += 3
		} else if oponent == Paper {
			score += 1
		} else if oponent == Scissor {
			score += 2
		}
	case Win:
		if oponent == Rock {
			score += 2 + 6
		} else if oponent == Paper {
			score += 3 + 6
		} else if oponent == Scissor {
			score += 1 + 6
		}
	case Draw:
		if oponent == Rock {
			score += 1 + 3
		} else if oponent == Paper {
			score += 2 + 3
		} else if oponent == Scissor {
			score += 3 + 3
		}
	}
	return score
}
