package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	P1_Rock    string = "A"
	P1_Paper          = "B"
	P1_Scissor        = "C"
)

const (
	P2_Rock    string = "X"
	P2_Paper          = "Y"
	P2_Scissor        = "Z"
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
		if choose[0] == P1_Rock {
			if choose[1] == P2_Rock {
				score += 1 + 3
			} else if choose[1] == P2_Paper {
				score += 2 + 6
			} else if choose[1] == P2_Scissor {
				score += 3 + 0
			}
		} else if choose[0] == P1_Paper {
			if choose[1] == P2_Rock {
				score += 1 + 0
			} else if choose[1] == P2_Paper {
				score += 2 + 3
			} else if choose[1] == P2_Scissor {
				score += 3 + 6
			}
		} else if choose[0] == P1_Scissor {
			if choose[1] == P2_Rock {
				score += 1 + 6
			} else if choose[1] == P2_Paper {
				score += 2 + 0
			} else if choose[1] == P2_Scissor {
				score += 3 + 3
			}
		}
	}
	return score
}
