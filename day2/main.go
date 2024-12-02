package day2

import (
	_ "embed"
	"slices"
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
	safeReports := 0

	for _, line := range strings.Split(input, "\n") {
		report, _ := StringsToInts(strings.Split(line, " "))

		if isSafeReport(report, false) {
			safeReports++
		}
	}

	return safeReports
}

func part2(input string) int {
	safeReports := 0
	for _, line := range strings.Split(input, "\n") {
		report, _ := StringsToInts(strings.Split(line, " "))

		if isSafeReport(report, true) {
			safeReports++
		}
	}

	return safeReports
}

func isSafeReport(report []int, tolerateError bool) bool {
	increasingMode := report[0] < report[1]
	for index, value := range report {
		if index == 0 {
			continue
		}

		var indexMinOne = report[index-1]
		if increasingMode && value > indexMinOne && value-indexMinOne <= 3 {
			continue
		} else if !increasingMode && value < indexMinOne && indexMinOne-value <= 3 {
			continue
		}

		if tolerateError {
			var r = make([]int, len(report))
			for i, _ := range report {
				copy(r, report)
				if isSafeReport(slices.Delete(r, i, i+1), false) {
					return true
				}
			}
		}

		return false
	}

	return true
}
