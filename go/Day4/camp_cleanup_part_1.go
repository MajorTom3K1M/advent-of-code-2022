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
		if isContain(left, right) || isContain(right, left) {
			count++
		}
	}
	return count
}

func isContain(shift1 []int, shift2 []int) bool {
	if shift1[0] <= shift2[0] && shift1[1] >= shift2[1] {
		return true
	}
	return false
}

func splittedPart(part string) []int {
	arr := strings.Split(part, "-")
	left, _ := strconv.Atoi(arr[0])
	right, _ := strconv.Atoi(arr[1])
	return []int{left, right}
}
