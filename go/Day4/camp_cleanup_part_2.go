package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	assignments := string(input)
	count := overlapAssignments(assignments)
	fmt.Printf("Answer : %d", count)
}

func overlapAssignments(assignments string) int {
	splittedAssignments := strings.Split(assignments, "\r\n")
	count := 0
	for _, assignment := range splittedAssignments {
		shifts := strings.Split(assignment, ",")
		left := make([]int, 0)
		right := make([]int, 0)
		left = append(left, splittedPart(shifts[0])...)
		right = append(right, splittedPart(shifts[1])...)
		if isOverlap(left, right) || isOverlap(right, left) {
			count++
		}
	}
	return count
}

func isOverlap(current []int, other []int) bool {
	return (current[0] <= other[0] && current[1] >= other[1]) || (current[0] <= other[0] && current[1] >= other[0])
}

func splittedPart(part string) []int {
	arr := strings.Split(part, "-")
	left, _ := strconv.Atoi(arr[0])
	right, _ := strconv.Atoi(arr[1])
	return []int{left, right}
}
