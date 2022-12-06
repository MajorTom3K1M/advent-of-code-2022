package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	datastream := string(input)
	marker := tuningTrouble(datastream)
	fmt.Printf("First marker after character : %d", marker)
}

func tuningTrouble(datastream string) int {
	for i := 0; i <= len(datastream)-4; i++ {
		found := false
		for _, stream := range datastream[i : i+4] {
			foundIndex := strings.Index(datastream[i:i+4], string(stream))
			compare := datastream[i:i+foundIndex] + datastream[i+foundIndex+1:i+4]
			if strings.Contains(compare, string(stream)) {
				found = true
				break
			}
		}

		if found {
			continue
		}
		return i + 4
	}
	return -1
}
