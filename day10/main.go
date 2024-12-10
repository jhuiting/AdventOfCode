package day10

import (
	_ "embed"
	"strconv"
	"strings"

	. "eye.security/v2/helpers"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("The input.txt file is empty")
	}
}

const (
	trailHead = 0
)

func part1(input string) int {
	grid, trailHeads := makeGrid(input)

	trails := 0
	for trailStart := range trailHeads {
		trailsForTrailHead := map[Position]bool{}

		for _, newPosition := range walkGrid([]Position{trailStart}, grid) {
			if !trailsForTrailHead[newPosition] {
				trailsForTrailHead[newPosition] = true
			}
		}

		trails += len(trailsForTrailHead)
	}

	return trails
}

func part2(input string) int {
	grid, trailHeads := makeGrid(input)

	trails := 0
	for trailStart := range trailHeads {
		trails += len(walkGrid([]Position{trailStart}, grid))
	}

	return trails
}

func walkGrid(currentPositions []Position, grid map[Position]int) []Position {
	lastPositions := map[int][]Position{0: currentPositions}

	var checkAndAppendPosition = func(direction Direction, position Position, index int) {
		if currentValue := grid[position]; grid[position.Move(direction)] == currentValue+1 {
			lastPositions[index] = append(lastPositions[index], position.Move(direction))
		}
	}

	for i := 0; i < 9; i++ {
		for _, currentPosition := range lastPositions[i] {
			checkAndAppendPosition(Up, currentPosition, i+1)
			checkAndAppendPosition(Down, currentPosition, i+1)
			checkAndAppendPosition(Left, currentPosition, i+1)
			checkAndAppendPosition(Right, currentPosition, i+1)
		}

	}
	return lastPositions[9]
}

func makeGrid(input string) (map[Position]int, map[Position]bool) {
	var grid = make(map[Position]int)
	trailHeads := make(map[Position]bool)

	for y, line := range strings.Split(input, "\n") {
		for x, number := range strings.Split(line, "") {
			val, _ := strconv.Atoi(number)
			grid[Position{X: x, Y: y}] = val
			if val == trailHead {
				trailHeads[Position{X: x, Y: y}] = true
			}

		}
	}
	return grid, trailHeads
}
