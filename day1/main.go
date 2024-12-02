package day1

import (
	_ "embed"
	"fmt"
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
	var leftPairs, rightPairs = createPairs(input)
	slices.Sort(leftPairs)
	slices.Sort(rightPairs)

	totalDistance := 0
	for index, left := range leftPairs {
		right := rightPairs[index]
		if left < right {
			totalDistance += right - left
		} else {
			totalDistance += left - right
		}
	}

	return totalDistance
}

func part2(input string) int {
	var leftPairs, rightPairs = createPairs(input)

	totalSimilarity := 0
	for _, left := range leftPairs {
		totalSimilarity += Reduce(rightPairs, 0, func(acc int, i int, elem int) int {
			if left == elem {
				return acc + (1 * left)
			}

			return acc
		})
	}

	return totalSimilarity
}

func createPairs(input string) ([]int, []int) {
	var leftPairs, rightPairs = make([]int, 0), make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		var n1, n2 int
		_, _ = fmt.Sscanf(line, "%d   %d", &n1, &n2)
		leftPairs, rightPairs = append(leftPairs, n1), append(rightPairs, n2)
	}

	return leftPairs, rightPairs
}

type AccumulatorFunc[T any] func(acc T, i int, s T) T

func Reduce[T any](items []T, initialAccumulator T, f AccumulatorFunc[T]) T {
	if items == nil {
		return initialAccumulator
	}
	acc := initialAccumulator
	for i, s := range items {
		acc = f(acc, i, s)
	}
	return acc
}
