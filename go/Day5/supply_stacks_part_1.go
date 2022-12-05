package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	stacksProcedure := string(input)
	top := rearrangement(stacksProcedure)
	fmt.Printf("Answer : %s", top)
}

func rearrangement(stacksProcedure string) string {
	top := ""
	split := strings.Split(stacksProcedure, "\r\n\r\n")
	stack := containerToSlice(split[0])
	instruction := instructionToSlice(split[1])
	for i := 0; i < len(instruction); i += 3 {
		move, _ := strconv.Atoi(instruction[i])
		from, _ := strconv.Atoi(instruction[i+1])
		to, _ := strconv.Atoi(instruction[i+2])
		for _, s := range stack[from-1][0:move] {
			stack[to-1] = append([]string{s}, stack[to-1]...) // this line doesn't need to be in loop
		}
		stack[from-1] = stack[from-1][move:]
	}
	for _, s := range stack {
		top += s[0]
	}
	return toAlphabets(top)
}

func toAlphabets(str string) string {
	alphabets := regexp.MustCompile("[A-Z]").FindAllString(str, -1)
	return strings.Join(alphabets, "")
}

func instructionToSlice(instructionString string) []string {
	digit := regexp.MustCompile("\\d+").FindAllString(instructionString, -1)
	return digit
}

func containerToSlice(stackString string) [][]string {
	whitespace := regexp.MustCompile("^[^\\w]+$")
	stackArr := regexp.MustCompile("\\s{3,4}|(\\[[A-Z]\\])").FindAllString(stackString, -1)
	digitArr := regexp.MustCompile("\\d+").FindAllString(stackString, -1)
	size, _ := strconv.Atoi(digitArr[len(digitArr)-1])
	stack := make([][]string, size)
	count := 0
	for _, stackElem := range stackArr {
		if count >= size {
			count = 0
		}

		if stack[count] == nil {
			stack[count] = make([]string, 0)
		}

		if !whitespace.MatchString(stackElem) {
			stack[count] = append(stack[count], stackElem)
		}
		count++
	}

	return stack
}
