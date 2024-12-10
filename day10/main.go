package day10

import (
	_ "embed"
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

type Position struct {
	x int
	y int
}

func (position Position) Turn(direction Direction) Direction {
	switch direction {
	case Up:
		direction = Right
	case Right:
		direction = Down
	case Down:
		direction = Left
	case Left:
		direction = Up
	}
	return direction
}

func (position Position) Move(direction Direction) Position {
	switch direction {
	case Up:
		return Position{x: position.x, y: position.y - 1}
	case Down:
		return Position{x: position.x, y: position.y + 1}
	case Left:
		return Position{x: position.x - 1, y: position.y}
	case Right:
		return Position{x: position.x + 1, y: position.y}
	}

	panic("Invalid direction")
}

type Direction int

const (
	Up    Direction = 1
	Down  Direction = 2
	Left  Direction = 3
	Right Direction = 4
)

const (
	trailHead = 0
	trailEnd  = 9
)

func part1(input string) int {
	grid, trailHeads := makeGrid(input)

	trails := 0
	for trailStart := range trailHeads {
		trailsForTrailHead := map[Position]bool{}
		currentPositions := walkGrid([]Position{trailStart}, grid)

		for _, newPosition := range currentPositions {
			if !trailsForTrailHead[newPosition] && grid[newPosition] == trailEnd {
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
		currentPositions := walkGrid([]Position{trailStart}, grid)

		trails += len(currentPositions)

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
			grid[Position{x, y}] = val
			if val == trailHead {
				trailHeads[Position{x, y}] = true
			}

		}
	}
	return grid, trailHeads
}
