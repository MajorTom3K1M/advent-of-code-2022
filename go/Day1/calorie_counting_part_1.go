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
	calories := string(input)
	result := mostCalories(calories)
	fmt.Printf("Answer : %d", result)
}

func mostCalories(calories string) int {
	splittedInventories := strings.Split(calories, "\r\n")

	inventoryCalories := 0
	maxCalorie := 0
	for _, calorie := range splittedInventories {
		if calorie != "" {
			i, _ := strconv.Atoi(calorie)
			inventoryCalories += i
		} else {
			if maxCalorie < inventoryCalories {
				maxCalorie = inventoryCalories
			}
			inventoryCalories = 0
		}
	}

	if maxCalorie < inventoryCalories {
		maxCalorie = inventoryCalories
	}

	return maxCalorie
}
