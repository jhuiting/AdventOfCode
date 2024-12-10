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
				trails += 1
			}
		}

	}

	return trails
}

func part2(input string) int {
	grid, trailHeads := makeGrid(input)

	trails := 0
	for trailStart := range trailHeads {
		currentPositions := walkGrid([]Position{trailStart}, grid)

		for _, newPosition := range currentPositions {
			if grid[newPosition] == trailEnd {
				trails += 1
			}
		}

	}

	return trails
}

func walkGrid(currentPositions []Position, grid map[Position]int) []Position {
	for i := 1; i <= 9; i++ {
		for _, currentPosition := range currentPositions {
			if grid[currentPosition.Move(Up)] == grid[currentPosition]+1 {
				currentPositions = append(currentPositions, currentPosition.Move(Up))
			}
			if grid[currentPosition.Move(Down)] == grid[currentPosition]+1 {
				currentPositions = append(currentPositions, currentPosition.Move(Down))
			}
			if grid[currentPosition.Move(Left)] == grid[currentPosition]+1 {
				currentPositions = append(currentPositions, currentPosition.Move(Left))

			}
			if grid[currentPosition.Move(Right)] == grid[currentPosition]+1 {
				currentPositions = append(currentPositions, currentPosition.Move(Right))
			}
		}
	}
	return currentPositions
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
