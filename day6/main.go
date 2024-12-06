package day6

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

type Position struct {
	x int
	y int
}

func (position Position) ChangePosition(direction Direction) Position {
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
	UpChar       = "^"
	ObstacleChar = "#"
	PathChar     = "."
)

func part1(input string) int {
	grid, position := createGrid(input)

	positions, _ := walkGrid(position, Up, grid)
	return len(positions)
}

func part2(input string) int {
	grid, position := createGrid(input)

	loops := 0
	positionsOnPath, _ := walkGrid(position, Up, grid)

	for y, row := range grid {
		for x, cell := range row {
			if cell == PathChar && slices.Index(positionsOnPath, Position{x, y}) != -1 {
				// Just try when it's on the guard's path
				grid[y][x] = ObstacleChar
				if _, isLoop := walkGrid(position, Up, grid); isLoop {
					loops += 1
				}
				grid[y][x] = PathChar
			}

		}
	}

	return loops
}

func createGrid(input string) ([][]string, Position) {
	position := Position{0, 0}
	grid := make([][]string, 0)
	for lineNumber, line := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(line, ""))
		if result := slices.Index(grid[lineNumber], UpChar); result != -1 {
			position = Position{result, lineNumber}
		}
	}

	return grid, position
}

func walkGrid(position Position, direction Direction, grid [][]string) ([]Position, bool) {
	visitedPositions := []Position{position}
	positionOccurrences := make(map[Position]int)

	for {
		newPosition := position.ChangePosition(direction)
		positionOccurrences[newPosition] += 1

		if (newPosition.x < 0 || newPosition.x >= len(grid[0])) || (newPosition.y < 0 || newPosition.y >= len(grid)) {
			return visitedPositions, false
		} else if positionOccurrences[newPosition] > 5 {
			return visitedPositions, true
		}

		if direction == Up && grid[newPosition.y][newPosition.x] == ObstacleChar {
			direction = Right
			continue
		} else if direction == Right && grid[newPosition.y][newPosition.x] == ObstacleChar {
			direction = Down
			continue
		} else if direction == Down && grid[newPosition.y][newPosition.x] == ObstacleChar {
			direction = Left
			continue
		} else if direction == Left && grid[newPosition.y][newPosition.x] == ObstacleChar {
			direction = Up
			continue
		}

		position = newPosition
		if slices.Index(visitedPositions, newPosition) == -1 {
			visitedPositions = append(visitedPositions, position)
		}
	}
}
