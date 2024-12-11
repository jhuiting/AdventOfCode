package day11

import (
	_ "embed"
	"eye.security/v2/helpers"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("The input.txt file is empty")
	}
}

func part1(input string) int {
	totalCount := 0
	for _, count := range getStonesForInput(input, 25) {
		totalCount += int(count)
	}

	return totalCount
}

func part2(input string) int {
	totalCount := 0
	for _, count := range getStonesForInput(input, 75) {
		totalCount += int(count)
	}

	return totalCount
}

func getStonesForInput(input string, turns int) map[int]int {
	var allNumbers = make(map[int]int)
	for _, number := range helpers.ToInts(regexp.MustCompile(`\d+`).FindAllString(input, -1)) {
		allNumbers[number] = 1
	}

	for i := 0; i < turns; i++ {
		allNumbers = blink(allNumbers)
	}

	return allNumbers
}

func blink(input map[int]int) map[int]int {
	var allNumbers = make(map[int]int)

	for n, count := range input {
		if n == 0 {
			allNumbers[1] += count
		} else if strNumber := strconv.Itoa(n); len(strNumber)%2 == 0 {
			firstValue, _ := strconv.Atoi(strNumber[:len(strNumber)/2])
			lastValue, _ := strconv.Atoi(strNumber[len(strNumber)/2:])

			allNumbers[firstValue] = allNumbers[firstValue] + count
			allNumbers[lastValue] = allNumbers[lastValue] + count
		} else {
			allNumbers[n*2024] = allNumbers[n*2024] + count
		}
	}

	return allNumbers
}
