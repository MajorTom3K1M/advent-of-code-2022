package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	calories := string(input)
	top := topThreeCalories(calories)

	fmt.Printf("Answer : %d", sum(top))
}

func sum(nums []int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

func topThreeCalories(calories string) []int {
	splittedInventories := strings.Split(calories, "\r\n")
	top := make([]int, 3)

	inventoryCalories := 0
	for _, calorie := range splittedInventories {
		if calorie != "" {
			i, _ := strconv.Atoi(calorie)
			inventoryCalories += i
		} else {
			if top[0] < inventoryCalories {
				top[0] = inventoryCalories
			}
			sort.Ints(top)
			inventoryCalories = 0
		}
	}

	if top[0] < inventoryCalories {
		top[0] = inventoryCalories
	}

	return top
}
