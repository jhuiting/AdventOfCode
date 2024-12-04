package day

import (
	_ "embed"
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
	xmasMatches := 0
	var columns = make(map[int][]string)

	for lineNumber, line := range strings.Split(input, "\n") {
		xmasMatches += strings.Count(line, "XMAS")
		xmasMatches += strings.Count(line, "SAMX")

		for index, character := range strings.Split(line, "") {
			columns[index] = append(columns[index], character)

			// Backward
			if lineNumber >= 3 && index >= 3 && (character == "X" &&
				columns[index-1][lineNumber-1] == "M" &&
				columns[index-2][lineNumber-2] == "A" &&
				columns[index-3][lineNumber-3] == "S" ||
				character == "S" &&
					columns[index-1][lineNumber-1] == "A" &&
					columns[index-2][lineNumber-2] == "M" &&
					columns[index-3][lineNumber-3] == "X") {
				xmasMatches += 1
			}

			// Forward
			if lineNumber >= 3 && index <= len(line)-4 && (character == "X" &&
				columns[index+1][lineNumber-1] == "M" &&
				columns[index+2][lineNumber-2] == "A" &&
				columns[index+3][lineNumber-3] == "S" || character == "S" &&
				columns[index+1][lineNumber-1] == "A" &&
				columns[index+2][lineNumber-2] == "M" &&
				columns[index+3][lineNumber-3] == "X") {
				xmasMatches += 1
			}
		}
	}

	for _, column := range columns {
		xmasMatches += strings.Count(strings.Join(column, ""), "XMAS")
		xmasMatches += strings.Count(strings.Join(column, ""), "SAMX")
	}

	return xmasMatches
}

func part2(input string) int {
	xmasMatches := 0
	var columns = make(map[int][]string)

	for lineNumber, line := range strings.Split(input, "\n") {
		for index, character := range strings.Split(line, "") {
			columns[index] = append(columns[index], character)

			if (lineNumber >= 2 && index >= 2 && character == "M" &&
				columns[index-1][lineNumber-1] == "A" &&
				columns[index-2][lineNumber-2] == "S") ||
				(lineNumber >= 2 && index >= 2 && character == "S" &&
					columns[index-1][lineNumber-1] == "A" &&
					columns[index-2][lineNumber-2] == "M") {

				if line[index-2] == 'S' &&
					columns[index-1][lineNumber-1] == "A" &&
					columns[index][lineNumber-2] == "M" {
					xmasMatches += 1
				} else if line[index-2] == 'M' &&
					columns[index-1][lineNumber-1] == "A" &&
					columns[index][lineNumber-2] == "S" {
					xmasMatches += 1
				}
			}
		}
	}

	return xmasMatches
}
