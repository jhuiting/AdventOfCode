package day9

import (
	_ "embed"
	"image"
	"slices"
	"strconv"
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

type File struct {
	ID   int
	Size int
}

func part1(input string) int {
	var uncompressed []int
	for index, blockItem := range strings.Split(input, "") {
		val, _ := strconv.Atoi(blockItem)
		if (index % 2) == 0 {
			uncompressed = append(uncompressed, slices.Repeat([]int{index / 2}, val)...)
		} else {
			uncompressed = append(uncompressed, slices.Repeat([]int{-1}, val)...)
		}
	}

	leftPointer := 0
	rightPointer := len(uncompressed) - 1

	for {
		for leftPointer < len(uncompressed) && uncompressed[leftPointer] != -1 {
			leftPointer++
		}

		for rightPointer >= 0 && uncompressed[rightPointer] == -1 {
			rightPointer--
		}

		if leftPointer > rightPointer {
			break
		}

		uncompressed[leftPointer] = uncompressed[rightPointer]
		uncompressed[rightPointer] = -1
	}

	sum := 0
	for index, number := range uncompressed {
		if number == -1 {
			continue
		}

		sum += index * number
	}

	return sum
}

func part2(input string) int {

	return 0
}
