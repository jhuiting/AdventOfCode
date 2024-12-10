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

func part1(input string) int {
	var grid [][]int
	trailHeads := make(map[Position]bool)
	for y, line := range strings.Split(input, "\n") {
		for x, number := range strings.Split(line, " ") {
			val, _ := strconv.Atoi(number)
			grid = append(grid, []int{x, y, val})
			if val == 0 {
				trailHeads[Position{x, y}] = true
			}

		}
	}

	return 0
}

func part2(input string) int {

	return 0
}
