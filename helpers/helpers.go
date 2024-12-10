package helpers

import "strconv"

type Position struct {
	X, Y int
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
		return Position{X: position.X, Y: position.Y - 1}
	case Down:
		return Position{X: position.X, Y: position.Y + 1}
	case Left:
		return Position{X: position.X - 1, Y: position.Y}
	case Right:
		return Position{X: position.X + 1, Y: position.Y}
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

func ToInts(strings []string) []int {
	ints := make([]int, len(strings))
	for i, s := range strings {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints[i] = num
	}

	return ints
}
