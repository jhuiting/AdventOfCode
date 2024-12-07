package day5

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"

	helpers "eye.security/v2/helpers"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("The input.txt file is empty")
	}
}

var rules [][]int
var cmp = func(left int, right int) int {
	for _, rule := range rules {
		if rule[0] == right && rule[1] == left {
			return 1
		}
		if rule[0] == left && rule[1] == right {
			return -1
		}
	}
	return 0
}

func part1(input string) int {
	validNumbers := 0
	rules = [][]int{}

	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, "|") {
			var n, m int
			_, _ = fmt.Sscanf(line, "%d|%d", &n, &m)
			rules = append(rules, []int{n, m})
		} else if line != "" {
			if items := helpers.ToInts(strings.Split(line, ",")); slices.IsSortedFunc(items, cmp) {
				validNumbers += items[len(items)/2]
			}
		}
	}

	return validNumbers
}

func part2(input string) int {
	validNumbers := 0
	rules = [][]int{}

	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, "|") {
			var n, m int
			_, _ = fmt.Sscanf(line, "%d|%d", &n, &m)
			rules = append(rules, []int{n, m})
		} else if line != "" {
			if items := helpers.ToInts(strings.Split(line, ",")); !slices.IsSortedFunc(items, cmp) {
				slices.SortFunc(items, cmp)
				validNumbers += items[len(items)/2]
			}
		}
	}

	return validNumbers
}
