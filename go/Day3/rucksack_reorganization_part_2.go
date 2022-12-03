package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	rucksacks := string(input)
	sum := prioritizeItems(rucksacks)
	fmt.Printf("Answer : %d", sum)
}

func prioritizeItems(rucksacks string) int {
	splittedRucksack := strings.Split(rucksacks, "\r\n")
	sum := 0
	for i := 0; i < len(splittedRucksack); i += 3 {
		for _, item := range splittedRucksack[i] {
			if strings.Contains(splittedRucksack[i+1], string(item)) && strings.Contains(splittedRucksack[i+2], string(item)) {
				sum += priorityNumber(item)
				break
			}
		}
	}
	return sum
}

func priorityNumber(letter rune) int {
	if byte(letter) > byte('Z') {
		return int(byte(letter) - byte('a') + 1)
	}
	return int(byte(letter) - byte('A') + 27)
}
