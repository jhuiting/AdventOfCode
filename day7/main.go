package day7

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
	solution := 0
	for _, line := range strings.Split(input, "\n") {
		var re = regexp.MustCompile("\\d+")
		if numbers := helpers.ToInts(re.FindAllString(line, -1)); matchPartOne(numbers, 1, 0) {
			solution += numbers[0]
		}
	}

	return solution
}

func matchPartOne(numbers []int, i int, total int) bool {
	expectedTotal := numbers[0]
	if endOfSlice := len(numbers)-1 == i; endOfSlice {
		return expectedTotal == total+numbers[i] || expectedTotal == total*numbers[i]
	}
	return matchPartOne(numbers, i+1, total+numbers[i]) || matchPartOne(numbers, i+1, total*numbers[i])
}

func part2(input string) int {
	solution := 0
	for _, line := range strings.Split(input, "\n") {
		var re = regexp.MustCompile("\\d+")
		if numbers := helpers.ToInts(re.FindAllString(line, -1)); matchPartTwo(numbers, 1, 0) {
			solution += numbers[0]
		}
	}

	return solution
}

func matchPartTwo(numbers []int, i int, total int) bool {
	expectedTotal := numbers[0]
	concatValue, _ := strconv.Atoi(strconv.Itoa(total) + strconv.Itoa(numbers[i]))

	if endOfSlice := len(numbers)-1 == i; endOfSlice {
		return expectedTotal == total+numbers[i] || expectedTotal == total*numbers[i] || expectedTotal == concatValue
	}
	return matchPartTwo(numbers, i+1, total+numbers[i]) || matchPartTwo(numbers, i+1, total*numbers[i]) || matchPartTwo(numbers, i+1, concatValue)
}
