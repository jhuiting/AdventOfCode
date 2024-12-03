package day

import (
	_ "embed"
	"fmt"
	"regexp"
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
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)

	totalCount := 0
	for _, match := range re.FindAllString(input, -1) {
		var n, m int
		_, _ = fmt.Sscanf(match, "mul(%d,%d)", &n, &m)
		totalCount += n * m
	}

	return totalCount
}

func part2(input string) int {
	re := regexp.MustCompile(`(do\(\))|(don't\(\))|mul\(\d+,\d+\)`)

	totalCount := 0
	enabled := true
	for _, match := range re.FindAllString(input, -1) {
		switch {
		case strings.Contains(match, "don't"):
			enabled = false
		case strings.Contains(match, "do"):
			enabled = true
		case strings.Contains(match, "mul"):
			if enabled {
				var n, m int
				_, _ = fmt.Sscanf(match, "mul(%d,%d)", &n, &m)
				totalCount += n * m
			}
		}
	}

	return totalCount
}
