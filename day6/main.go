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
	upChar       = "^"
	obstacleChar = "#"
	pathChar     = "."
)

func part1(input string) int {
	grid, position := createGrid(input)

	positions, _ := walkGrid(position, Up, grid, false)
	return len(positions)
}

func part2(input string) int {
	grid, position := createGrid(input)

	loops := 0
	visitedPositions, _ := walkGrid(position, Up, grid, false)

	for y, row := range grid {
		for x, cell := range row {
			if _, ok := visitedPositions[Position{x, y}]; cell == pathChar && ok {
				// Just try when it's on the guard's path
				grid[y][x] = obstacleChar
				if _, isLoop := walkGrid(position, Up, grid, true); isLoop {
					loops += 1
				}
				grid[y][x] = pathChar
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
		if result := slices.Index(grid[lineNumber], upChar); result != -1 {
			position = Position{result, lineNumber}
		}
	}

	return grid, position
}

func walkGrid(position Position, direction Direction, grid [][]string, checkLoops bool) (map[Position]int, bool) {
	positionsOccurrences := map[Position]int{
		position: 1,
	}

	for {
		newPosition := position.Move(direction)
		if (newPosition.x < 0 || newPosition.x >= len(grid[0])) || (newPosition.y < 0 || newPosition.y >= len(grid)) {
			return positionsOccurrences, false
		}

		if grid[newPosition.y][newPosition.x] == obstacleChar {
			// Just check for possible loops on turns
			if positionsOccurrences[newPosition] > 5 {
				return positionsOccurrences, true
			} else if checkLoops {
				positionsOccurrences[newPosition] += 1
			}

			direction = position.Turn(direction)
			continue
		}

		positionsOccurrences[newPosition] += 1
		position = newPosition

	}
}
