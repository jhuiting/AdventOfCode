package day8

import (
	_ "embed"
	"image"
	"strings"
)

type Point = image.Point

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("The input.txt file is empty")
	}
}

const (
	emptyFieldChar = "."
)

func part1(input string) int {
	gridBounds, antennaPositions := getBoundsAndPositions(input)

	solution := map[image.Point]bool{}
	for _, antennas := range antennaPositions {
		for _, a1 := range antennas {
			for _, a2 := range antennas {
				if a2 == a1 {
					continue
				}
				if antinode := a2.Add(a2.Sub(a1)); gridBounds[antinode] {
					solution[antinode] = true
				}
			}
		}
	}

	return len(solution)
}
func part2(input string) int {
	gridBounds, antennaPositions := getBoundsAndPositions(input)

	solution := map[image.Point]bool{}
	for _, antennas := range antennaPositions {
		for _, a1 := range antennas {
			for _, a2 := range antennas {
				if a2 == a1 {
					continue
				}

				for distance := a2.Sub(a1); gridBounds[a2]; a2 = a2.Add(distance) {
					solution[a2] = true
				}
			}
		}
	}

	return len(solution)
}

func getBoundsAndPositions(input string) (map[Point]bool, map[string][]Point) {
	gridBounds, antennaPositions := map[Point]bool{}, map[string][]Point{}
	for y, line := range strings.Split(input, "\n") {
		columnItems := strings.Split(line, "")
		for x, char := range columnItems {
			gridBounds[Point{X: x, Y: y}] = true
			if char != emptyFieldChar {
				antennaPositions[char] = append(antennaPositions[char], Point{X: x, Y: y})
			}
		}
	}

	return gridBounds, antennaPositions
}
