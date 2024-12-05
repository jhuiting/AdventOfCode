package day

import (
	_ "embed"
	"fmt"
	"slices"
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

var rules [][]int
var cmp = func(left, right int) int {
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
		if line == "" {
			continue
		}

		if strings.Contains(line, "|") {
			var n, m int
			_, _ = fmt.Sscanf(line, "%d|%d", &n, &m)
			rules = append(rules, []int{n, m})
		} else {
			items, _ := StringsToInts(strings.Split(line, ","))
			if slices.IsSortedFunc(items, cmp) {
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
		if line == "" {
			continue
		}

		if strings.Contains(line, "|") {
			var n, m int
			_, _ = fmt.Sscanf(line, "%d|%d", &n, &m)
			rules = append(rules, []int{n, m})
		} else {
			items, _ := StringsToInts(strings.Split(line, ","))
			if !slices.IsSortedFunc(items, cmp) {
				slices.SortFunc(items, cmp)
				validNumbers += items[len(items)/2]
			}
		}
	}

	return validNumbers
}

func StringsToInts(strings []string) ([]int, error) {
	ints := make([]int, len(strings))
	for i, s := range strings {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ints[i] = num
	}

	return ints, nil
}
